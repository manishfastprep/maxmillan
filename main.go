package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter total revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Enter total expenses:")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate:")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	ratio := ebt / eat

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Earnings After Tax (EAT): %.2f\n", eat)
	fmt.Printf("Tax Efficiency Ratio: %.2f\n", ratio)
}
