package bamboohr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"
)

// GetNewEmployees fetches employees from BambooHR and filters for those hired in a specific month/year
func GetNewEmployees(month, year string) (map[string]interface{}, error) {
	// Fetch all employees
	employees, err := getAllEmployees()
	if err != nil {
		return nil, err
	}

	// Filter employees hired in the specified month and year
	newEmployees := filterEmployeesByHireDate(employees, month, year)

	return map[string]interface{}{
		"employees": newEmployees,
		"count":     len(newEmployees),
		"month":     month,
		"year":      year,
	}, nil
}

// GetNewEmployeesForYear fetches all employees hired in a specific year
func GetNewEmployeesForYear(year string) (map[string]interface{}, error) {
	// Fetch all employees once
	employees, err := getAllEmployees()
	if err != nil {
		return nil, err
	}

	// Filter employees hired in the specified year
	newEmployees := filterEmployeesByYear(employees, year)

	// Group by month for better organization
	employeesByMonth := make(map[string][]model.BambooHREmployee)
	for _, emp := range newEmployees {
		hireDate, err := time.Parse("2006-01-02", emp.HireDate)
		if err != nil {
			continue
		}
		month := fmt.Sprintf("%02d", int(hireDate.Month()))
		employeesByMonth[month] = append(employeesByMonth[month], emp)
	}

	return map[string]interface{}{
		"employees":        newEmployees,
		"employeesByMonth": employeesByMonth,
		"count":            len(newEmployees),
		"year":             year,
	}, nil
}

// getAllEmployees fetches all employees from BambooHR API
func getAllEmployees() ([]model.BambooHREmployee, error) {
	url := getAPIURL() + "employees/directory"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// BambooHR uses API key as username, password is left blank
	apiKey := config.Get("bamboohrApiKey")
	req.SetBasicAuth(apiKey, "x")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("BambooHR API returned status %d", resp.StatusCode)
	}

	var result struct {
		Employees []model.BambooHREmployee `json:"employees"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Employees, nil
}

// filterEmployeesByHireDate filters employees by hire date (month and year)
func filterEmployeesByHireDate(employees []model.BambooHREmployee, month, year string) []model.BambooHREmployee {
	var filtered []model.BambooHREmployee

	for _, emp := range employees {
		if emp.HireDate == "" {
			continue
		}

		// Parse hire date (expected format: YYYY-MM-DD)
		hireDate, err := time.Parse("2006-01-02", emp.HireDate)
		if err != nil {
			continue
		}

		// Check if hire date matches the specified month and year
		if fmt.Sprintf("%d", hireDate.Year()) == year && fmt.Sprintf("%02d", int(hireDate.Month())) == month {
			filtered = append(filtered, emp)
		}
	}

	return filtered
}

// filterEmployeesByYear filters employees hired in a specific year
func filterEmployeesByYear(employees []model.BambooHREmployee, year string) []model.BambooHREmployee {
	var filtered []model.BambooHREmployee

	for _, emp := range employees {
		if emp.HireDate == "" {
			continue
		}

		// Parse hire date (expected format: YYYY-MM-DD)
		hireDate, err := time.Parse("2006-01-02", emp.HireDate)
		if err != nil {
			continue
		}

		// Check if hire date matches the specified year
		if fmt.Sprintf("%d", hireDate.Year()) == year {
			filtered = append(filtered, emp)
		}
	}

	return filtered
}

// getAPIURL constructs the BambooHR API base URL
func getAPIURL() string {
	domain := config.Get("bamboohrDomain")
	return fmt.Sprintf("https://api.bamboohr.com/api/gateway.php/%s/v1/", domain)
}

// GetEmployeeFields fetches specific employee fields from BambooHR
func GetEmployeeFields(fields []string) (map[string]interface{}, error) {
	fieldsList := strings.Join(fields, ",")
	url := getAPIURL() + "employees/directory?fields=" + fieldsList

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	apiKey := config.Get("bamboohrApiKey")
	req.SetBasicAuth(apiKey, "x")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("BambooHR API returned status %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
