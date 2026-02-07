package model

// -------- Employee Models --------

type BambooEmployee struct {
	ID             string `json:"id,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	Email          string `json:"workEmail,omitempty"`
	PersonalEmail  string `json:"homeEmail,omitempty"`
	Phone          string `json:"workPhone,omitempty"`
	MobilePhone    string `json:"mobilePhone,omitempty"`
	JobTitle       string `json:"jobTitle,omitempty"`
	Department     string `json:"department,omitempty"`
	Division       string `json:"division,omitempty"`
	Location       string `json:"location,omitempty"`
	Status         string `json:"status,omitempty"`
	HireDate       string `json:"hireDate,omitempty"`
	TerminationDate string `json:"terminationDate,omitempty"`
	SupervisorID   string `json:"supervisorId,omitempty"`
	Supervisor     string `json:"supervisor,omitempty"`
	PayRate        string `json:"payRate,omitempty"`
	PayType        string `json:"payType,omitempty"`
	PayPeriod      string `json:"payPeriod,omitempty"`
	EmploymentType string `json:"employmentHistoryStatus,omitempty"`
	PhotoURL       string `json:"photoUrl,omitempty"`
}

type BambooEmployeeDirectory struct {
	Employees []BambooEmployee `json:"employees"`
}

type BambooEmployeeResponse struct {
	ID     string            `json:"id"`
	Fields map[string]string `json:"fields"`
}

// -------- Time Off Models --------

type BambooTimeOffRequest struct {
	ID        string             `json:"id,omitempty"`
	EmployeeID string            `json:"employeeId,omitempty"`
	Status    BambooApprovalStatus `json:"status,omitempty"`
	Start     string             `json:"start,omitempty"`
	End       string             `json:"end,omitempty"`
	Type      BambooTimeOffType  `json:"type,omitempty"`
	Amount    BambooTimeOffAmount `json:"amount,omitempty"`
	Notes     BambooTimeOffNotes `json:"notes,omitempty"`
	Created   string             `json:"created,omitempty"`
}

type BambooApprovalStatus struct {
	LastChanged    string `json:"lastChanged,omitempty"`
	LastChangedByID string `json:"lastChangedByUserId,omitempty"`
	Status         string `json:"status,omitempty"`
}

type BambooTimeOffType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BambooTimeOffAmount struct {
	Unit   string `json:"unit,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type BambooTimeOffNotes struct {
	Employee string `json:"employee,omitempty"`
	Manager  string `json:"manager,omitempty"`
}

type BambooTimeOffRequests struct {
	Requests []BambooTimeOffRequest `json:"requests"`
}

// -------- Applicant Tracking (Hiring) Models --------

type BambooJobOpening struct {
	ID              string `json:"id,omitempty"`
	Title           string `json:"jobTitle,omitempty"`
	Department      string `json:"departmentId,omitempty"`
	Location        string `json:"locationId,omitempty"`
	Status          string `json:"jobStatus,omitempty"`
	HiringLead      string `json:"hiringLead,omitempty"`
	EmploymentType  string `json:"employmentType,omitempty"`
	MinSalary       string `json:"minimumExperience,omitempty"`
	Description     string `json:"description,omitempty"`
	DateCreated     string `json:"dateCreated,omitempty"`
}

type BambooApplicant struct {
	ID            string `json:"id,omitempty"`
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phoneNumber,omitempty"`
	JobOpeningID  string `json:"jobOpeningId,omitempty"`
	Status        string `json:"status,omitempty"`
	Rating        int    `json:"rating,omitempty"`
	Source        string `json:"source,omitempty"`
	DateApplied   string `json:"dateApplied,omitempty"`
	ResumeURL     string `json:"resumeFileUrl,omitempty"`
	CoverLetterURL string `json:"coverLetterFileUrl,omitempty"`
}

type BambooApplicants struct {
	Applicants []BambooApplicant `json:"applicants"`
}

// -------- Performance Review Models --------

type BambooGoal struct {
	ID           string `json:"id,omitempty"`
	EmployeeID   string `json:"employeeId,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	PercentComplete int `json:"percentComplete,omitempty"`
	Status       string `json:"status,omitempty"`
	DueDate      string `json:"dueDate,omitempty"`
	MilestonesCompleted int `json:"completedCount,omitempty"`
	MilestonesTotal     int `json:"totalCount,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
}

type BambooGoals struct {
	Goals []BambooGoal `json:"goals"`
}

// -------- Training Models --------

type BambooTrainingRecord struct {
	ID          string `json:"id,omitempty"`
	EmployeeID  string `json:"employeeId,omitempty"`
	Type        string `json:"type,omitempty"`
	Category    string `json:"category,omitempty"`
	Completed   string `json:"completed,omitempty"`
	Cost        string `json:"cost,omitempty"`
	Instructor  string `json:"instructor,omitempty"`
	Hours       string `json:"hours,omitempty"`
	Credits     string `json:"credits,omitempty"`
	Notes       string `json:"notes,omitempty"`
}

type BambooTrainingRecords struct {
	Records []BambooTrainingRecord `json:"records"`
}

// -------- Onboarding Models --------

type BambooOnboardingTask struct {
	ID          string `json:"id,omitempty"`
	EmployeeID  string `json:"employeeId,omitempty"`
	TaskName    string `json:"taskName,omitempty"`
	Category    string `json:"category,omitempty"`
	DueDate     string `json:"dueDate,omitempty"`
	Status      string `json:"status,omitempty"`
	AssignedTo  string `json:"assignedTo,omitempty"`
	CompletedAt string `json:"completedAt,omitempty"`
}

// -------- HR Agent Request/Response Models --------

type HRAgentRequest struct {
	Action     string                 `json:"action"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Context    string                 `json:"context,omitempty"`
}

type HRAgentResponse struct {
	Action  string      `json:"action"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
