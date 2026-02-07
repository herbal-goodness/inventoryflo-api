package bamboohr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"
)

const apiVersion = "v1"

// ---------- Employee Operations ----------

// GetEmployeeDirectory returns all employees in the company directory
func GetEmployeeDirectory() (*model.BambooEmployeeDirectory, error) {
	url := buildURL("employees/directory")

	body, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee directory: %w", err)
	}

	var directory model.BambooEmployeeDirectory
	if err := json.Unmarshal(body, &directory); err != nil {
		return nil, fmt.Errorf("failed to decode employee directory: %w", err)
	}
	return &directory, nil
}

// GetEmployee returns a single employee by ID with the specified fields
func GetEmployee(employeeID string) (*model.BambooEmployee, error) {
	fields := "firstName,lastName,displayName,workEmail,homeEmail,workPhone,mobilePhone," +
		"jobTitle,department,division,location,status,hireDate,terminationDate," +
		"supervisorId,supervisor,payRate,payType,payPeriod,employmentHistoryStatus,photoUrl"

	url := buildURL(fmt.Sprintf("employees/%s/?fields=%s", employeeID, fields))

	body, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee %s: %w", employeeID, err)
	}

	var emp model.BambooEmployee
	if err := json.Unmarshal(body, &emp); err != nil {
		return nil, fmt.Errorf("failed to decode employee: %w", err)
	}
	return &emp, nil
}

// AddEmployee creates a new employee in BambooHR
func AddEmployee(employee model.BambooEmployee) (*model.BambooEmployee, error) {
	url := buildURL("employees/")

	payload, err := json.Marshal(employee)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee: %w", err)
	}

	body, err := doPost(url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to add employee: %w", err)
	}

	var created model.BambooEmployee
	if err := json.Unmarshal(body, &created); err != nil {
		return nil, fmt.Errorf("failed to decode created employee: %w", err)
	}
	return &created, nil
}

// UpdateEmployee updates an existing employee's fields
func UpdateEmployee(employeeID string, fields map[string]string) error {
	url := buildURL(fmt.Sprintf("employees/%s/", employeeID))

	payload, err := json.Marshal(fields)
	if err != nil {
		return fmt.Errorf("failed to marshal fields: %w", err)
	}

	_, err = doPost(url, payload)
	if err != nil {
		return fmt.Errorf("failed to update employee %s: %w", employeeID, err)
	}
	return nil
}

// ---------- Time Off Operations ----------

// GetTimeOffRequests returns time off requests for a date range
func GetTimeOffRequests(start, end string) (*model.BambooTimeOffRequests, error) {
	url := buildURL(fmt.Sprintf("time_off/requests/?start=%s&end=%s&status=approved,pending,denied", start, end))

	body, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get time off requests: %w", err)
	}

	var requests model.BambooTimeOffRequests
	if err := json.Unmarshal(body, &requests); err != nil {
		return nil, fmt.Errorf("failed to decode time off requests: %w", err)
	}
	return &requests, nil
}

// ---------- Applicant Tracking (Hiring) Operations ----------

// GetJobOpenings returns all active job openings
func GetJobOpenings() ([]model.BambooJobOpening, error) {
	url := buildURL("applicant_tracking/job_summaries")

	body, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get job openings: %w", err)
	}

	var result struct {
		JobSummaries []model.BambooJobOpening `json:"jobSummaries"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode job openings: %w", err)
	}
	return result.JobSummaries, nil
}

// GetApplicants returns applicants for a specific job opening
func GetApplicants(jobOpeningID string) (*model.BambooApplicants, error) {
	url := buildURL(fmt.Sprintf("applicant_tracking/jobs/%s/applicants", jobOpeningID))

	body, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get applicants: %w", err)
	}

	var applicants model.BambooApplicants
	if err := json.Unmarshal(body, &applicants); err != nil {
		return nil, fmt.Errorf("failed to decode applicants: %w", err)
	}
	return &applicants, nil
}

// ChangeApplicantStatus moves an applicant to a new stage in the hiring pipeline
func ChangeApplicantStatus(applicantID string, statusID string) error {
	url := buildURL(fmt.Sprintf("applicant_tracking/applicants/%s/status", applicantID))

	payload, _ := json.Marshal(map[string]string{"status": statusID})
	_, err := doPost(url, payload)
	if err != nil {
		return fmt.Errorf("failed to change applicant status: %w", err)
	}
	return nil
}

// ---------- Goals / Performance Operations ----------

// GetGoals returns performance goals for an employee
func GetGoals(employeeID string) (*model.BambooGoals, error) {
	url := buildURL(fmt.Sprintf("performance/employees/%s/goals", employeeID))

	body, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get goals for employee %s: %w", employeeID, err)
	}

	var goals model.BambooGoals
	if err := json.Unmarshal(body, &goals); err != nil {
		return nil, fmt.Errorf("failed to decode goals: %w", err)
	}
	return &goals, nil
}

// ---------- Training Operations ----------

// GetTrainingRecords returns training records for an employee
func GetTrainingRecords(employeeID string) (*model.BambooTrainingRecords, error) {
	url := buildURL(fmt.Sprintf("training/record/%s", employeeID))

	body, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get training records: %w", err)
	}

	var records model.BambooTrainingRecords
	if err := json.Unmarshal(body, &records); err != nil {
		return nil, fmt.Errorf("failed to decode training records: %w", err)
	}
	return &records, nil
}

// ---------- Aggregate / HR Agent Dispatcher ----------

// HandleHRAction is the main dispatcher for the HR sub-agent. It routes
// incoming HRAgentRequest actions to the appropriate BambooHR API call.
func HandleHRAction(req model.HRAgentRequest) model.HRAgentResponse {
	switch req.Action {
	case "list_employees":
		dir, err := GetEmployeeDirectory()
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, dir)

	case "get_employee":
		id, _ := req.Parameters["employee_id"].(string)
		emp, err := GetEmployee(id)
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, emp)

	case "add_employee":
		payload, _ := json.Marshal(req.Parameters)
		var emp model.BambooEmployee
		json.Unmarshal(payload, &emp)
		created, err := AddEmployee(emp)
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, created)

	case "update_employee":
		id, _ := req.Parameters["employee_id"].(string)
		fields := make(map[string]string)
		if f, ok := req.Parameters["fields"].(map[string]interface{}); ok {
			for k, v := range f {
				fields[k] = fmt.Sprintf("%v", v)
			}
		}
		if err := UpdateEmployee(id, fields); err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, map[string]string{"updated": id})

	case "get_time_off":
		start, _ := req.Parameters["start"].(string)
		end, _ := req.Parameters["end"].(string)
		requests, err := GetTimeOffRequests(start, end)
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, requests)

	case "list_job_openings":
		openings, err := GetJobOpenings()
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, openings)

	case "get_applicants":
		jobID, _ := req.Parameters["job_opening_id"].(string)
		applicants, err := GetApplicants(jobID)
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, applicants)

	case "change_applicant_status":
		applicantID, _ := req.Parameters["applicant_id"].(string)
		statusID, _ := req.Parameters["status_id"].(string)
		if err := ChangeApplicantStatus(applicantID, statusID); err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, map[string]string{"applicant": applicantID, "new_status": statusID})

	case "get_goals":
		empID, _ := req.Parameters["employee_id"].(string)
		goals, err := GetGoals(empID)
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, goals)

	case "get_training":
		empID, _ := req.Parameters["employee_id"].(string)
		records, err := GetTrainingRecords(empID)
		if err != nil {
			return errorResponse(req.Action, err)
		}
		return successResponse(req.Action, records)

	default:
		return model.HRAgentResponse{
			Action:  req.Action,
			Status:  "error",
			Message: fmt.Sprintf("Unknown HR action: %s", req.Action),
		}
	}
}

// ---------- HTTP Helpers ----------

func buildURL(path string) string {
	companyDomain := config.Get("bamboohrDomain")
	return fmt.Sprintf("https://api.bamboohr.com/api/gateway.php/%s/%s/%s", companyDomain, apiVersion, path)
}

func doGet(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(config.Get("bamboohrApiKey"), "x")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("BambooHR API returned status %d: %s", resp.StatusCode, string(body))
	}

	return ioutil.ReadAll(resp.Body)
}

func doPost(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(config.Get("bamboohrApiKey"), "x")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("BambooHR API returned status %d: %s", resp.StatusCode, string(body))
	}

	return ioutil.ReadAll(resp.Body)
}

// ---------- Response Helpers ----------

func successResponse(action string, data interface{}) model.HRAgentResponse {
	return model.HRAgentResponse{
		Action: action,
		Status: "success",
		Data:   data,
	}
}

func errorResponse(action string, err error) model.HRAgentResponse {
	return model.HRAgentResponse{
		Action:  action,
		Status:  "error",
		Message: err.Error(),
	}
}
