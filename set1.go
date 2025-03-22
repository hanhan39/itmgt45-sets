package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Savings:")
	fmt.Println(savings(1000, 0.12, 80))

	fmt.Println("Material Waste:")
	fmt.Println(materialWaste(460, "kg", 8, 24))

	fmt.Println("Interest:")
	fmt.Println(interest(12000, 0.03, 8))
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
