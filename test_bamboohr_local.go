// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
	"github.com/herbal-goodness/inventoryflo-api/pkg/service/bamboohr"
)

func main() {
	fmt.Println("=============================================================")
	fmt.Println("BambooHR New Employees Test - February 2026")
	fmt.Println("=============================================================")
	fmt.Println()

	// Simulate the API response
	fmt.Println("Simulating GET /bamboohr/new-employees/02/2026")
	fmt.Println()

	result := bamboohr.SimulateGetNewEmployees()

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
	fmt.Printf("Summary: Found %v new employee(s) in %s/%s\n",
		result["count"], result["month"], result["year"])
	fmt.Println("=============================================================")
	fmt.Println()

	// Display employee details in a table format
	if employees, ok := result["employees"].([]model.BambooHREmployee); ok && len(employees) > 0 {
		fmt.Println("New Employees Hired in February 2026:")
		fmt.Println("----------------------------------------")
		for i, emp := range employees {
			fmt.Printf("%d. %s %s (Employee #%s)\n", i+1, emp.FirstName, emp.LastName, emp.EmployeeNumber)
			fmt.Printf("   Email: %s\n", emp.Email)
			fmt.Printf("   Department: %s\n", emp.Department)
			fmt.Printf("   Job Title: %s\n", emp.JobTitle)
			fmt.Printf("   Hire Date: %s\n", emp.HireDate)
			fmt.Printf("   Location: %s\n", emp.Location)
			fmt.Printf("   Mobile: %s\n", emp.MobilePhone)
			fmt.Println()
		}
	}
}
