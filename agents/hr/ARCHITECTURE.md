# HR Sub-Agent Architecture

## Overview

The HR Sub-Agent is an AI-powered HR operations system for Herbal Goodness that connects
to BambooHR and automates the full employee lifecycle. It runs in two modes:

1. **Direct API** — Go service integrated into inventoryflo-api (AWS Lambda)
2. **n8n Orchestration** — Visual workflow automation with AI-powered intent classification

```
┌──────────────────────────────────────────────────────────┐
│                     Entry Points                          │
│  Slack Bot │ API Gateway │ n8n Webhooks │ Scheduled Jobs  │
└──────┬─────────────┬──────────────┬──────────────┬───────┘
       │             │              │              │
       ▼             ▼              ▼              ▼
┌──────────────────────────────────────────────────────────┐
│              AI Intent Classifier (LLM)                   │
│  "Show me all open positions" → list_job_openings         │
│  "Onboard Jane Smith"         → add_employee              │
│  "Fire someone"               → escalate (human required) │
└──────────────────────┬───────────────────────────────────┘
                       │
                       ▼
┌──────────────────────────────────────────────────────────┐
│              Action Router (Switch)                        │
│                                                           │
│  list_employees ──────┐                                   │
│  get_employee ────────┤                                   │
│  add_employee ────────┤                                   │
│  update_employee ─────┤                                   │
│  get_time_off ────────┼──► BambooHR API                   │
│  list_job_openings ───┤                                   │
│  get_applicants ──────┤                                   │
│  change_applicant_status─┤                                │
│  get_goals ───────────┤                                   │
│  get_training ────────┘                                   │
│  escalate ────────────────► Slack #hr-escalations         │
└──────────────────────┬───────────────────────────────────┘
                       │
                       ▼
┌──────────────────────────────────────────────────────────┐
│           AI Response Formatter (LLM)                     │
│  Raw API data → Human-readable summary with tables        │
└──────────────────────┬───────────────────────────────────┘
                       │
                       ▼
┌──────────────────────────────────────────────────────────┐
│              Response Delivery                            │
│  Slack message │ API JSON │ Email │ Webhook response      │
└──────────────────────────────────────────────────────────┘
```

## File Structure

```
inventoryflo-api/
├── pkg/
│   ├── model/
│   │   └── bamboohr.go              # BambooHR data models
│   ├── service/
│   │   └── bamboohr/
│   │       └── bamboohr.go          # BambooHR API client + HR action dispatcher
│   ├── router/
│   │   └── router.go                # POST /hr route added
│   └── util/config/
│       └── config.go                # bamboohrDomain + bamboohrApiKey added
├── agents/hr/
│   ├── ARCHITECTURE.md              # This file
│   ├── system_prompt.md             # System prompt for the AI agent
│   └── functions.json               # Function definitions (tool schema)
├── n8n/workflows/
│   ├── hr_agent_main.json           # Main orchestrator workflow
│   ├── hr_onboarding_automation.json # New hire onboarding automation
│   ├── hr_performance_review_cycle.json # Quarterly review generation
│   └── hr_offboarding_automation.json   # Employee separation workflow
└── n8n/
    └── N8N_SETUP.md                 # Step-by-step n8n setup guide
```

## BambooHR API Integration

### Authentication
- BambooHR uses HTTP Basic Auth: API key as username, "x" as password
- API keys are encrypted with AES-256-GCM via AWS KMS (same pattern as Shopify)
- Base URL: `https://api.bamboohr.com/api/gateway.php/{companyDomain}/v1/`

### Supported Operations

| Action | HTTP Method | BambooHR Endpoint | Side Effects | Needs Approval |
|--------|-------------|-------------------|--------------|----------------|
| list_employees | GET | /employees/directory | No | No |
| get_employee | GET | /employees/{id}/ | No | No |
| add_employee | POST | /employees/ | Yes | Yes |
| update_employee | POST | /employees/{id}/ | Yes | Yes |
| get_time_off | GET | /time_off/requests/ | No | No |
| list_job_openings | GET | /applicant_tracking/job_summaries | No | No |
| get_applicants | GET | /applicant_tracking/jobs/{id}/applicants | No | No |
| change_applicant_status | POST | /applicant_tracking/applicants/{id}/status | Yes | Yes |
| get_goals | GET | /performance/employees/{id}/goals | No | No |
| get_training | GET | /training/record/{id} | No | No |

## n8n Workflows

### 1. Main Orchestrator (`hr_agent_main.json`)
The primary workflow that handles all ad-hoc HR requests:
- Webhook receives natural language or structured requests
- AI classifies intent and extracts parameters
- Switch routes to the correct BambooHR API call
- AI formats the response into human-readable text
- Returns via webhook response

### 2. Onboarding Automation (`hr_onboarding_automation.json`)
Triggered when a new hire is confirmed:
- Creates employee in BambooHR
- Sends Slack welcome message to department channel
- Sends welcome email with first-week schedule
- Posts onboarding checklist to #hr-onboarding

### 3. Performance Review Cycle (`hr_performance_review_cycle.json`)
Runs quarterly on a cron schedule:
- Pulls all employees from directory
- Fetches goals/progress for each employee
- AI generates a draft performance summary
- Emails summary to employee
- Logs to #hr-performance Slack channel

### 4. Offboarding Automation (`hr_offboarding_automation.json`)
Triggered after HR manager approves separation:
- Fetches employee details from BambooHR
- Sets employee status to Inactive with termination date
- Posts offboarding checklist to Slack (IT, HR, Manager tasks)
- Sends transition email to departing employee

## AI Functions the Agent Can Manage

### Fully Automated (No Human Needed)
- Employee directory lookups and searches
- Individual employee profile retrieval
- Time-off balance and request reporting
- Job opening status reports
- Applicant pipeline visibility
- Goal progress tracking
- Training record retrieval
- Generating performance review drafts
- Sending automated onboarding communications
- Creating onboarding/offboarding checklists

### Semi-Automated (AI Prepares, Human Approves)
- Creating new employee records
- Moving applicants through hiring stages
- Updating employee profile fields
- Generating offer letter drafts
- Scheduling interviews based on availability
- Benefits enrollment reminders

### Human-Only (Agent Escalates)
- Termination / involuntary separation execution
- Salary and compensation changes
- Legal and compliance decisions (ADA, FMLA, EEOC)
- Harassment or discrimination case handling
- Final performance review ratings
- Offer approvals

## Best Prompts for the HR Agent

### For Intent Classification (in n8n AI node)
```
You are an HR operations agent for Herbal Goodness. You are connected to BambooHR.
Parse the user's request and determine which HR action to take.

Respond with a JSON object containing:
- "action": one of [list_employees, get_employee, add_employee, update_employee,
  get_time_off, list_job_openings, get_applicants, change_applicant_status,
  get_goals, get_training, escalate]
- "parameters": object with required params for the action

If the request involves termination, legal issues, salary changes, or harassment,
always set action to "escalate".
```

### For Response Formatting (in n8n AI node)
```
You are an HR assistant. Take the BambooHR API response data and the original user
request, then compose a clear, professional, human-readable summary. Use tables for
lists. Be concise. Never expose sensitive PII like SSN or bank details.
```

### For Performance Review Generation
```
You are an HR performance analyst. Given an employee's goals data, generate a brief
quarterly performance summary. Include: 1) Overall progress percentage, 2) Goals on
track vs behind, 3) Recommended focus areas for next quarter. Keep it to 3-4 sentences.
Be constructive and professional.
```

## Security Considerations

1. **API Key Storage**: BambooHR API keys are encrypted at rest using AES-256-GCM
   with AWS KMS master keys, following the same pattern as existing Shopify secrets.

2. **PII Protection**: The AI response formatter is instructed to never expose SSN,
   bank details, salary, or medical information in summaries.

3. **Escalation Guardrails**: Termination, legal, and compensation actions always
   route to human HR managers — the agent cannot execute these autonomously.

4. **Audit Trail**: All n8n workflow executions are logged. Slack notifications
   create a searchable record of all HR actions taken.

5. **Least Privilege**: The BambooHR API key should be scoped to only the endpoints
   the agent needs. Avoid using an admin-level key.
