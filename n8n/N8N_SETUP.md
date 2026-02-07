# n8n Setup Guide for HR Sub-Agent

## Prerequisites

1. **n8n instance** — Self-hosted or n8n Cloud
2. **BambooHR account** with API access enabled
3. **Slack workspace** with a bot token (for notifications)
4. **OpenAI API key** (for AI classification and formatting nodes)

## Step 1: Install n8n

```bash
# Option A: Docker (recommended for production)
docker run -it --rm \
  --name n8n \
  -p 5678:5678 \
  -v ~/.n8n:/home/node/.n8n \
  -e N8N_ENCRYPTION_KEY=your-encryption-key \
  n8nio/n8n

# Option B: npm (for development)
npm install n8n -g
n8n start
```

## Step 2: Configure Credentials

In n8n, go to **Settings > Credentials** and create:

### BambooHR API (HTTP Basic Auth)
- **Name**: `BambooHR API Key`
- **User**: Your BambooHR API key (from BambooHR > Account > API Keys)
- **Password**: `x` (literal string "x" — BambooHR convention)

### Slack Bot
- **Name**: `Slack HR Bot`
- **Access Token**: Your Slack Bot OAuth token (`xoxb-...`)
- Required Slack scopes: `chat:write`, `channels:read`

### OpenAI
- **Name**: `OpenAI`
- **API Key**: Your OpenAI API key

## Step 3: Set Environment Variables

In n8n settings or your `.env` file:

```
BAMBOOHR_DOMAIN=herbalgoodness
```

This is your BambooHR subdomain (from `https://herbalgoodness.bamboohr.com`).

## Step 4: Import Workflows

1. Open n8n at `http://localhost:5678`
2. Go to **Workflows > Import from File**
3. Import each workflow JSON file from `n8n/workflows/`:
   - `hr_agent_main.json` — Main orchestrator
   - `hr_onboarding_automation.json` — New hire automation
   - `hr_performance_review_cycle.json` — Quarterly reviews
   - `hr_offboarding_automation.json` — Separation automation

4. For each workflow:
   - Open it in the editor
   - Click each node that has credential references
   - Select the matching credential you created in Step 2
   - Save and **Activate** the workflow

## Step 5: Create Slack Channels

Create these Slack channels and invite the HR bot:
- `#hr-escalations` — For items requiring human HR review
- `#hr-onboarding` — Onboarding checklists and tracking
- `#hr-offboarding` — Offboarding checklists and tracking
- `#hr-performance` — Performance review cycle notifications

## Step 6: Test the Main Orchestrator

Send a POST request to test the webhook:

```bash
curl -X POST http://localhost:5678/webhook/hr-agent \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Show me all employees in the Engineering department"
  }'
```

Expected flow:
1. Webhook receives the request
2. AI classifies intent as `list_employees`
3. Router sends request to BambooHR directory API
4. AI formats employee list into readable response
5. JSON response returned

## Step 7: Test Onboarding

```bash
curl -X POST http://localhost:5678/webhook/hr-agent/onboard \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Jane",
    "lastName": "Smith",
    "email": "jane.smith@herbalgoodness.com",
    "jobTitle": "Product Manager",
    "department": "Product",
    "hireDate": "2026-03-01",
    "managerId": "123"
  }'
```

## Step 8: Connect to Slack (Optional Chat Interface)

To allow HR requests via Slack:
1. In n8n, create a new workflow with a **Slack Trigger** node (listens for messages)
2. Filter for messages in a `#hr-requests` channel or DMs to the bot
3. Pass the message text to the Main Orchestrator webhook
4. Return the response as a Slack message reply

## Workflow Diagram

```
Slack / API / Cron
       │
       ▼
  ┌─────────┐     ┌───────────────┐     ┌─────────────┐
  │ Webhook  │────►│ AI Classifier │────►│ Action      │
  │ Trigger  │     │ (GPT-4)       │     │ Router      │
  └─────────┘     └───────────────┘     └──────┬──────┘
                                               │
                    ┌──────────────────────────┤
                    │          │               │
                    ▼          ▼               ▼
              ┌──────────┐ ┌────────┐  ┌───────────┐
              │ BambooHR │ │ Slack  │  │ Email     │
              │ API Call │ │ Notify │  │ Send      │
              └────┬─────┘ └────────┘  └───────────┘
                   │
                   ▼
            ┌──────────────┐     ┌─────────────┐
            │ AI Formatter │────►│ Respond to  │
            │ (GPT-4)      │     │ Webhook     │
            └──────────────┘     └─────────────┘
```

## Troubleshooting

- **401 from BambooHR**: Check API key is correct and has proper permissions
- **AI classification wrong**: Adjust the system prompt temperature (lower = more consistent)
- **Slack messages not posting**: Verify bot is invited to the target channels
- **n8n webhook not reachable**: Ensure n8n is exposed (use ngrok for local dev)
