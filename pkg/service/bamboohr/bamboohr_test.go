package bamboohr

import (
	"testing"
	"time"

	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
)

// TestFilterEmployeesByHireDate tests the employee filtering logic
func TestFilterEmployeesByHireDate(t *testing.T) {
	// Create sample employees with different hire dates
	employees := []model.BambooHREmployee{
		{
			ID:         "101",
			FirstName:  "Alice",
			LastName:   "Johnson",
			Email:      "alice.johnson@company.com",
			Department: "Engineering",
			JobTitle:   "Senior Software Engineer",
			HireDate:   "2026-02-01",
		},
		{
			ID:         "102",
			FirstName:  "Bob",
			LastName:   "Smith",
			Email:      "bob.smith@company.com",
			Department: "Marketing",
			JobTitle:   "Marketing Manager",
			HireDate:   "2026-02-15",
		},
		{
			ID:         "103",
			FirstName:  "Carol",
			LastName:   "Davis",
			Email:      "carol.davis@company.com",
			Department: "Sales",
			JobTitle:   "Sales Representative",
			HireDate:   "2026-01-20",
		},
		{
			ID:         "104",
			FirstName:  "David",
			LastName:   "Wilson",
			Email:      "david.wilson@company.com",
			Department: "Engineering",
			JobTitle:   "DevOps Engineer",
			HireDate:   "2026-02-28",
		},
		{
			ID:         "105",
			FirstName:  "Emma",
			LastName:   "Brown",
			Email:      "emma.brown@company.com",
			Department: "HR",
			JobTitle:   "HR Specialist",
			HireDate:   "2026-03-10",
		},
		{
			ID:         "106",
			FirstName:  "Frank",
			LastName:   "Taylor",
			Email:      "frank.taylor@company.com",
			Department: "Finance",
			JobTitle:   "Financial Analyst",
			HireDate:   "2025-12-15",
		},
	}

	// Test filtering for February 2026
	filtered := filterEmployeesByHireDate(employees, "02", "2026")

	// Expected: Alice, Bob, and David (hired in Feb 2026)
	expectedCount := 3
	if len(filtered) != expectedCount {
		t.Errorf("Expected %d employees, got %d", expectedCount, len(filtered))
	}

	// Verify the correct employees were filtered
	expectedIDs := map[string]bool{"101": true, "102": true, "104": true}
	for _, emp := range filtered {
		if !expectedIDs[emp.ID] {
			t.Errorf("Unexpected employee ID %s in filtered results", emp.ID)
		}

		// Verify hire date is in February 2026
		hireDate, _ := time.Parse("2006-01-02", emp.HireDate)
		if hireDate.Year() != 2026 || hireDate.Month() != time.February {
			t.Errorf("Employee %s has incorrect hire date: %s", emp.ID, emp.HireDate)
		}
	}
}

// TestFilterEmployeesByHireDateEdgeCases tests edge cases
func TestFilterEmployeesByHireDateEdgeCases(t *testing.T) {
	employees := []model.BambooHREmployee{
		{
			ID:         "201",
			FirstName:  "Test",
			LastName:   "User",
			Email:      "test@company.com",
			Department: "IT",
			JobTitle:   "Developer",
			HireDate:   "", // Empty hire date
		},
		{
			ID:         "202",
			FirstName:  "Invalid",
			LastName:   "Date",
			Email:      "invalid@company.com",
			Department: "IT",
			JobTitle:   "Developer",
			HireDate:   "invalid-date", // Invalid date format
		},
	}

	// Should return empty list for employees with invalid/empty dates
	filtered := filterEmployeesByHireDate(employees, "02", "2026")

	if len(filtered) != 0 {
		t.Errorf("Expected 0 employees with invalid dates, got %d", len(filtered))
	}
}

// SimulateGetNewEmployees simulates the GetNewEmployees function with mock data
func SimulateGetNewEmployees() map[string]interface{} {
	// Mock employee data for February 2026
	mockEmployees := []model.BambooHREmployee{
		{
			ID:              "101",
			EmployeeNumber:  "EMP-001",
			FirstName:       "Alice",
			LastName:        "Johnson",
			DisplayName:     "Alice Johnson",
			Email:           "alice.johnson@herbalgoodness.com",
			MobilePhone:     "+1-555-0101",
			WorkPhone:       "+1-555-0100",
			Department:      "Engineering",
			JobTitle:        "Senior Software Engineer",
			Location:        "New York, NY",
			Division:        "Technology",
			HireDate:        "2026-02-01",
			Status:          "Active",
			EmploymentStatus: "Full-time",
			Address1:        "123 Main St",
			City:            "New York",
			State:           "NY",
			Zip:             "10001",
			Country:         "United States",
		},
		{
			ID:              "102",
			EmployeeNumber:  "EMP-002",
			FirstName:       "Bob",
			LastName:        "Smith",
			DisplayName:     "Bob Smith",
			Email:           "bob.smith@herbalgoodness.com",
			MobilePhone:     "+1-555-0201",
			WorkPhone:       "+1-555-0200",
			Department:      "Marketing",
			JobTitle:        "Marketing Manager",
			Location:        "Los Angeles, CA",
			Division:        "Marketing & Sales",
			HireDate:        "2026-02-15",
			Status:          "Active",
			EmploymentStatus: "Full-time",
			Address1:        "456 Oak Ave",
			City:            "Los Angeles",
			State:           "CA",
			Zip:             "90001",
			Country:         "United States",
		},
		{
			ID:              "104",
			EmployeeNumber:  "EMP-004",
			FirstName:       "David",
			LastName:        "Wilson",
			DisplayName:     "David Wilson",
			Email:           "david.wilson@herbalgoodness.com",
			MobilePhone:     "+1-555-0401",
			WorkPhone:       "+1-555-0400",
			Department:      "Engineering",
			JobTitle:        "DevOps Engineer",
			Location:        "Austin, TX",
			Division:        "Technology",
			HireDate:        "2026-02-28",
			Status:          "Active",
			EmploymentStatus: "Full-time",
			Address1:        "789 Pine Rd",
			City:            "Austin",
			State:           "TX",
			Zip:             "73301",
			Country:         "United States",
		},
	}

	return map[string]interface{}{
		"employees": mockEmployees,
		"count":     len(mockEmployees),
		"month":     "02",
		"year":      "2026",
	}
}

// SimulateGetNewEmployeesForYear simulates the GetNewEmployeesForYear function with mock data for all of 2026
func SimulateGetNewEmployeesForYear() map[string]interface{} {
	// Mock employee data across different months in 2026
	mockEmployees := []model.BambooHREmployee{
		// January 2026
		{
			ID:              "100",
			EmployeeNumber:  "EMP-000",
			FirstName:       "Sarah",
			LastName:        "Martinez",
			DisplayName:     "Sarah Martinez",
			Email:           "sarah.martinez@herbalgoodness.com",
			MobilePhone:     "+1-555-0001",
			WorkPhone:       "+1-555-0000",
			Department:      "Sales",
			JobTitle:        "Sales Director",
			Location:        "Miami, FL",
			Division:        "Marketing & Sales",
			HireDate:        "2026-01-15",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
		// February 2026
		{
			ID:              "101",
			EmployeeNumber:  "EMP-001",
			FirstName:       "Alice",
			LastName:        "Johnson",
			DisplayName:     "Alice Johnson",
			Email:           "alice.johnson@herbalgoodness.com",
			MobilePhone:     "+1-555-0101",
			WorkPhone:       "+1-555-0100",
			Department:      "Engineering",
			JobTitle:        "Senior Software Engineer",
			Location:        "New York, NY",
			Division:        "Technology",
			HireDate:        "2026-02-01",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
		{
			ID:              "102",
			EmployeeNumber:  "EMP-002",
			FirstName:       "Bob",
			LastName:        "Smith",
			DisplayName:     "Bob Smith",
			Email:           "bob.smith@herbalgoodness.com",
			MobilePhone:     "+1-555-0201",
			WorkPhone:       "+1-555-0200",
			Department:      "Marketing",
			JobTitle:        "Marketing Manager",
			Location:        "Los Angeles, CA",
			Division:        "Marketing & Sales",
			HireDate:        "2026-02-15",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
		{
			ID:              "104",
			EmployeeNumber:  "EMP-004",
			FirstName:       "David",
			LastName:        "Wilson",
			DisplayName:     "David Wilson",
			Email:           "david.wilson@herbalgoodness.com",
			MobilePhone:     "+1-555-0401",
			WorkPhone:       "+1-555-0400",
			Department:      "Engineering",
			JobTitle:        "DevOps Engineer",
			Location:        "Austin, TX",
			Division:        "Technology",
			HireDate:        "2026-02-28",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
		// March 2026
		{
			ID:              "105",
			EmployeeNumber:  "EMP-005",
			FirstName:       "Emma",
			LastName:        "Brown",
			DisplayName:     "Emma Brown",
			Email:           "emma.brown@herbalgoodness.com",
			MobilePhone:     "+1-555-0501",
			WorkPhone:       "+1-555-0500",
			Department:      "HR",
			JobTitle:        "HR Specialist",
			Location:        "Chicago, IL",
			Division:        "Human Resources",
			HireDate:        "2026-03-10",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
		{
			ID:              "106",
			EmployeeNumber:  "EMP-006",
			FirstName:       "Michael",
			LastName:        "Chen",
			DisplayName:     "Michael Chen",
			Email:           "michael.chen@herbalgoodness.com",
			MobilePhone:     "+1-555-0601",
			WorkPhone:       "+1-555-0600",
			Department:      "Engineering",
			JobTitle:        "Data Engineer",
			Location:        "San Francisco, CA",
			Division:        "Technology",
			HireDate:        "2026-03-25",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
		// April 2026
		{
			ID:              "107",
			EmployeeNumber:  "EMP-007",
			FirstName:       "Jessica",
			LastName:        "Lee",
			DisplayName:     "Jessica Lee",
			Email:           "jessica.lee@herbalgoodness.com",
			MobilePhone:     "+1-555-0701",
			WorkPhone:       "+1-555-0700",
			Department:      "Product",
			JobTitle:        "Product Manager",
			Location:        "Seattle, WA",
			Division:        "Product",
			HireDate:        "2026-04-05",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
		// May 2026
		{
			ID:              "108",
			EmployeeNumber:  "EMP-008",
			FirstName:       "Ryan",
			LastName:        "O'Connor",
			DisplayName:     "Ryan O'Connor",
			Email:           "ryan.oconnor@herbalgoodness.com",
			MobilePhone:     "+1-555-0801",
			WorkPhone:       "+1-555-0800",
			Department:      "Finance",
			JobTitle:        "Financial Analyst",
			Location:        "Boston, MA",
			Division:        "Finance",
			HireDate:        "2026-05-12",
			Status:          "Active",
			EmploymentStatus: "Full-time",
		},
	}

	// Group by month
	employeesByMonth := make(map[string][]model.BambooHREmployee)
	for _, emp := range mockEmployees {
		month := emp.HireDate[5:7] // Extract month from YYYY-MM-DD
		employeesByMonth[month] = append(employeesByMonth[month], emp)
	}

	return map[string]interface{}{
		"employees":        mockEmployees,
		"employeesByMonth": employeesByMonth,
		"count":            len(mockEmployees),
		"year":             "2026",
	}
}
