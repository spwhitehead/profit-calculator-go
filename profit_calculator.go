package main

import (
	"fmt"
)

func main(){

	var annualRevenue float64
	var annualExpenses float64
	var taxRate float64

	fmt.Print("What is your annual revenue? ")
	fmt.Scan(&annualRevenue)

	fmt.Print("What are your annual expenses (in $)? ")
	fmt.Scan(&annualExpenses)

	fmt.Print("What is your tax rate? ")
	fmt.Scan(&taxRate)

	// Calculations

	earningsBeforeTax := annualRevenue - annualExpenses
	profit := earningsBeforeTax * (1-taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Print("Earnings Before Tax: ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Profit: ")
	fmt.Println(profit)

	fmt.Print("Ratio (EBT/Profit): ")
	fmt.Println(ratio)

}