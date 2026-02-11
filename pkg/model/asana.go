package model

// AsanaTasks represents the response from Asana API for tasks
type AsanaTasks struct {
	Data []AsanaTask `json:"data"`
}

// AsanaTask represents a single task from Asana
type AsanaTask struct {
	Gid          string         `json:"gid,omitempty"`
	Name         string         `json:"name,omitempty"`
	ResourceType string         `json:"resource_type,omitempty"`
	Assignee     *AsanaUser     `json:"assignee,omitempty"`
	Completed    bool           `json:"completed,omitempty"`
	DueOn        string         `json:"due_on,omitempty"`
	DueAt        string         `json:"due_at,omitempty"`
	CreatedAt    string         `json:"created_at,omitempty"`
	ModifiedAt   string         `json:"modified_at,omitempty"`
	Notes        string         `json:"notes,omitempty"`
	Projects     []AsanaProject `json:"projects,omitempty"`
	Workspace    *AsanaWorkspace `json:"workspace,omitempty"`
}

// AsanaUser represents an Asana user (assignee)
type AsanaUser struct {
	Gid          string `json:"gid,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

// AsanaProject represents a project reference in a task
type AsanaProject struct {
	Gid          string `json:"gid,omitempty"`
	Name         string `json:"name,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

// AsanaWorkspace represents a workspace reference
type AsanaWorkspace struct {
	Gid          string `json:"gid,omitempty"`
	Name         string `json:"name,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

// AsanaWorkspaces represents the response for workspaces list
type AsanaWorkspaces struct {
	Data []AsanaWorkspace `json:"data"`
}
