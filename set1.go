// Placeholder
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Savings:")
	first := savings(1000, 0.12, 80)
	fmt.Println(first)
	fmt.Println("Material Waste:")
	second := materialWaste(460, "kg", 8, 24)
	fmt.Println(second)
	fmt.Println("Interest:")
	third := interest(12000, 0.03, 8)
	fmt.Println(third)
}

// Savings
// Calculate the money remaining for an employee after taxes and expenses.
//
// To get the take-home pay of an employee:
// 1. Apply the tax rate to the gross pay of the employee, round down.
// 2. Subtract the expenses from the after-tax pay of the employee.
//
// Params:
// - grossPay, the gross pay of an employee for a certain time period, expressed in centavos
// - taxRate, the tax rate for a certain time period, expressed as a number between 0 and 1 (e.g., 0.12)
// - expenses, the expenses of an employee for a certain time period, expressed in centavos
//
// Returns:
// - the number of centavos remaining from an employee's pay after taxes and expenses
func savings(grossPay int, taxRate float64, expenses int) int {
	taxDeduct := int(math.Floor(float64(grossPay) * taxRate))
	takeHome := grossPay - taxDeduct - expenses
	return takeHome
}

// Material waste
// Calculate how much material input will be wasted after running a number of jobs that consume a set amount of material.
//
// To get the waste of a set of jobs:
// 1. Multiply the number of jobs by the material consumption per job
// 2. Subtract the total material consumed from the total material available
//
// Format the output as a string, annotated with the units in which the material is expressed. Do not add a space between the number and the unit.
//
// Params:
// - totalMaterial, how much material you have at the start
// - materialUnits, the unit used to express an amount of material, e.g., "kg"
// - numJobs, how many jobs to run
// - jobConsumption, how much material each job consumes
//
// Returns:
// - the amount of remaining material expressed with its unit (e.g., "10kg")
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	waste := totalMaterial - (jobConsumption * numJobs)
	fmtWaste := strconv.Itoa(waste) + materialUnits
	return fmtWaste
}

// Interest
// Calculate the final value of an investment after gaining simple interest over a number of periods.
//
// To get sample interest, simply multiply the principal to the quantity (rate * time). Add this amount to the principal to get the final value.
//
// Round down the final amount.
//
// Params:
// - principal, the principal, or starting, amount invested, expressed in centavos
// - rate, the interest rate per period, expressed as a decimal representation of a percentage (e.g., 3% is 0.03)
// - periods, the number of periods invested
//
// Returns:
// - the final value of the investment
func interest(principal int, rate float64, periods int) int {
	interest := float64(principal) * rate * float64(periods)
	investment := principal + int(interest)
	return investment
}
