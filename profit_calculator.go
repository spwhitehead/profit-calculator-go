package main

import (
	"fmt"
)

func main(){

	var annualRevenue float64
	var annualExpenses float64
	var taxRate float64

	getValue("Annual Revenue: ")
	fmt.Scan(&annualRevenue)

	getValue("Annual Expenses: ")
	fmt.Scan(&annualExpenses)

	getValue("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculations(annualRevenue, annualExpenses, taxRate)

	fmt.Printf(`
	Earnings Before Tax:	 %.2f
	Profit: 		 %.2f
	Ratio (ebt/profit): 	 %.2f`, ebt, profit, ratio)
	// fmt.Print("What is your annual revenue? ")
	// fmt.Scan(&annualRevenue)

	// fmt.Print("What are your annual expenses (in $)? ")
	// fmt.Scan(&annualExpenses)

	// fmt.Print("What is your tax rate? ")
	// fmt.Scan(&taxRate)

	// Calculations

	// earningsBeforeTax := annualRevenue - annualExpenses
	// profit := earningsBeforeTax * (1-taxRate/100)
	// ratio := earningsBeforeTax / profit

	// fmt.Print("Earnings Before Tax: ")
	// fmt.Println(earningsBeforeTax)

	// fmt.Print("Profit: ")
	// fmt.Println(profit)

	// fmt.Print("Ratio (EBT/Profit): ")
	// fmt.Println(ratio)

}

func getValue(value string){
	fmt.Print(value)
}

func calculations(annualRevenue, annualExpenses, taxRate float64) (ebt float64, profit float64, ratio float64){
	ebt = annualRevenue - annualExpenses
	profit = ebt * (1-taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio

}