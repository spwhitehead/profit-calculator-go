package main

import (
	"fmt"
)

func main(){


	annualRevenue := getValue("Annual Revenue: ")
	annualExpenses := getValue("Annual Expenses: ")
	taxRate := getValue("Tax Rate: ")

	ebt, profit, ratio := calculations(annualRevenue, annualExpenses, taxRate)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.4f\n", ratio)
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

func getValue(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculations(annualRevenue, annualExpenses, taxRate float64) (ebt float64, profit float64, ratio float64){
	ebt = annualRevenue - annualExpenses
	profit = ebt * (1-taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio

}