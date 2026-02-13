// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
	"github.com/herbal-goodness/inventoryflo-api/pkg/service/bamboohr"
)

func main() {
	fmt.Println("=============================================================")
	fmt.Println("BambooHR New Employees Test - All of 2026")
	fmt.Println("=============================================================")
	fmt.Println()

	// Simulate the API response
	fmt.Println("Simulating GET /bamboohr/new-employees/2026")
	fmt.Println()

	result := bamboohr.SimulateGetNewEmployeesForYear()

	// Convert to JSON for pretty printing
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	fmt.Println("Response:")
	fmt.Println(string(jsonData))
	fmt.Println()

	// Display summary
	fmt.Println("=============================================================")
	fmt.Printf("Summary: Found %v new employee(s) in %s\n",
		result["count"], result["year"])
	fmt.Println("=============================================================")
	fmt.Println()

	// Display employee details grouped by month
	if employees, ok := result["employees"].([]model.BambooHREmployee); ok && len(employees) > 0 {
		if employeesByMonth, ok := result["employeesByMonth"].(map[string][]model.BambooHREmployee); ok {
			monthNames := map[string]string{
				"01": "January",
				"02": "February",
				"03": "March",
				"04": "April",
				"05": "May",
				"06": "June",
				"07": "July",
				"08": "August",
				"09": "September",
				"10": "October",
				"11": "November",
				"12": "December",
			}

			// Sort months and display
			months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
			for _, month := range months {
				if emps, exists := employeesByMonth[month]; exists {
					fmt.Printf("\n%s 2026 (%d employee(s)):\n", monthNames[month], len(emps))
					fmt.Println("----------------------------------------")
					for i, emp := range emps {
						fmt.Printf("%d. %s %s (Employee #%s)\n", i+1, emp.FirstName, emp.LastName, emp.EmployeeNumber)
						fmt.Printf("   Email: %s\n", emp.Email)
						fmt.Printf("   Department: %s\n", emp.Department)
						fmt.Printf("   Job Title: %s\n", emp.JobTitle)
						fmt.Printf("   Hire Date: %s\n", emp.HireDate)
						fmt.Printf("   Location: %s\n", emp.Location)

						// Calculate days since hire
						hireDate, _ := time.Parse("2006-01-02", emp.HireDate)
						now := time.Date(2026, 2, 13, 0, 0, 0, 0, time.UTC)
						daysSince := int(now.Sub(hireDate).Hours() / 24)
						if daysSince >= 0 {
							fmt.Printf("   Days since hire: %d days\n", daysSince)
						} else {
							fmt.Printf("   Starts in: %d days\n", -daysSince)
						}
						fmt.Println()
					}
				}
			}
		}

		// Summary by department
		fmt.Println("\n=============================================================")
		fmt.Println("Summary by Department:")
		fmt.Println("=============================================================")
		deptCount := make(map[string]int)
		for _, emp := range employees {
			deptCount[emp.Department]++
		}
		for dept, count := range deptCount {
			fmt.Printf("  %s: %d employee(s)\n", dept, count)
		}
	}
}
