package model

// BambooHR Employee Related structs
type BambooHREmployees struct {
	Employees []BambooHREmployee `json:"employees"`
}

type BambooHREmployee struct {
	ID                 string `json:"id,omitempty"`
	EmployeeNumber     string `json:"employeeNumber,omitempty"`
	FirstName          string `json:"firstName,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	PreferredName      string `json:"preferredName,omitempty"`
	DisplayName        string `json:"displayName,omitempty"`
	Email              string `json:"workEmail,omitempty"`
	MobilePhone        string `json:"mobilePhone,omitempty"`
	WorkPhone          string `json:"workPhone,omitempty"`
	Department         string `json:"department,omitempty"`
	JobTitle           string `json:"jobTitle,omitempty"`
	Location           string `json:"location,omitempty"`
	Division           string `json:"division,omitempty"`
	HireDate           string `json:"hireDate,omitempty"`
	TerminationDate    string `json:"terminationDate,omitempty"`
	Status             string `json:"status,omitempty"`
	Supervisor         string `json:"supervisor,omitempty"`
	SupervisorEmail    string `json:"supervisorEmail,omitempty"`
	PayRate            string `json:"payRate,omitempty"`
	PayType            string `json:"payType,omitempty"`
	EmploymentStatus   string `json:"employmentStatus,omitempty"`
	Address1           string `json:"address1,omitempty"`
	Address2           string `json:"address2,omitempty"`
	City               string `json:"city,omitempty"`
	State              string `json:"state,omitempty"`
	Zip                string `json:"zipcode,omitempty"`
	Country            string `json:"country,omitempty"`
	DateOfBirth        string `json:"dateOfBirth,omitempty"`
	Gender             string `json:"gender,omitempty"`
	MaritalStatus      string `json:"maritalStatus,omitempty"`
	SSN                string `json:"ssn,omitempty"`
	CreatedByUserId    string `json:"createdByUserId,omitempty"`
	LastChanged        string `json:"lastChanged,omitempty"`
}
