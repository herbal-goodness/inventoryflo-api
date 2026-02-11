package asana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"
)

// GetOverdueTasks fetches all overdue tasks assigned to a specific email
func GetOverdueTasks(assigneeEmail string) (map[string]interface{}, error) {
	// First, get the user's workspaces
	workspaces, err := getWorkspaces()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch workspaces: %v", err)
	}

	if len(workspaces.Data) == 0 {
		return map[string]interface{}{
			"tasks": []model.AsanaTask{},
			"count": 0,
		}, nil
	}

	// Get tasks from all workspaces
	var allTasks []model.AsanaTask
	today := time.Now().Format("2006-01-02")

	for _, workspace := range workspaces.Data {
		tasks, err := getTasksForWorkspace(workspace.Gid, assigneeEmail)
		if err != nil {
			// Log error but continue with other workspaces
			fmt.Printf("Warning: failed to fetch tasks from workspace %s: %v\n", workspace.Name, err)
			continue
		}

		// Filter for overdue tasks
		for _, task := range tasks.Data {
			if !task.Completed && task.DueOn != "" {
				// Check if task is overdue
				if task.DueOn < today {
					allTasks = append(allTasks, task)
				}
			}
		}
	}

	return map[string]interface{}{
		"tasks": allTasks,
		"count": len(allTasks),
		"email": assigneeEmail,
	}, nil
}

// getWorkspaces fetches all workspaces the authenticated user has access to
func getWorkspaces() (*model.AsanaWorkspaces, error) {
	url := fmt.Sprintf("%s/workspaces", getBaseURL())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	setAuthHeader(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("asana API returned status %d: %s", resp.StatusCode, string(body))
	}

	var workspaces model.AsanaWorkspaces
	err = json.NewDecoder(resp.Body).Decode(&workspaces)
	if err != nil {
		return nil, err
	}

	return &workspaces, nil
}

// getTasksForWorkspace fetches tasks for a specific workspace and assignee
func getTasksForWorkspace(workspaceGid string, assigneeEmail string) (*model.AsanaTasks, error) {
	// Build URL with query parameters
	baseURL := fmt.Sprintf("%s/tasks", getBaseURL())
	params := url.Values{}
	params.Add("workspace", workspaceGid)
	params.Add("assignee.email", assigneeEmail)
	params.Add("completed_since", "now") // Only get incomplete tasks
	params.Add("opt_fields", "gid,name,assignee,assignee.name,assignee.email,completed,due_on,due_at,created_at,modified_at,notes,projects,projects.name,workspace,workspace.name")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	setAuthHeader(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("asana API returned status %d: %s", resp.StatusCode, string(body))
	}

	var tasks model.AsanaTasks
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}

// getBaseURL returns the Asana API base URL
func getBaseURL() string {
	baseURL := config.Get("asanaUrl")
	if baseURL == "" {
		return "https://app.asana.com/api/1.0"
	}
	return baseURL
}

// setAuthHeader sets the Bearer token authentication header
func setAuthHeader(req *http.Request) {
	token := config.Get("asanaToken")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json")
}
