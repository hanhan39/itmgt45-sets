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

func savings(grossPay int, taxRate float64, expenses int) int {
	taxDeduct := int(math.Floor(float64(grossPay) * taxRate))
	takeHome := grossPay - taxDeduct - expenses
	return takeHome
}

func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	waste := totalMaterial - (jobConsumption * numJobs)
	fmtWaste := strconv.Itoa(waste) + materialUnits
	return fmtWaste
}

func interest(principal int, rate float64, periods int) int {
	interest := float64(principal) * rate * float64(periods)
	investment := principal + int(interest)
	return investment
}
