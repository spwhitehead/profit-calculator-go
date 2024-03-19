package main

import (
	"fmt"
	"os"
)

func writeCalcuationsToFile(ebt, profit, ratio float64){
	calculationText := fmt.Sprintf("EBT: $%.2f\nProfit: $%.2f\nRatio: %.4f", ebt, profit, ratio)
	os.WriteFile("calculationsFile.txt", []byte(calculationText), 0644)
}

func getValue(infoText string) float64 {
	for {
		var userInput float64
		fmt.Print(infoText)
		_ , err := fmt.Scan(&userInput)
		if err != nil || userInput <= 0{
			fmt.Println()
			fmt.Println("Invalid input. Enter a number greater than zero.")
		} else {
			return userInput
		}
	}
}

func calculations(annualRevenue, annualExpenses, taxRate float64) (ebt float64, profit float64, ratio float64){
	ebt = annualRevenue - annualExpenses
	profit = ebt * (1-taxRate/100)
	ratio = ebt / profit
	writeCalcuationsToFile(ebt, profit, ratio)
	return ebt, profit, ratio
}

func main(){


	annualRevenue := getValue("Annual Revenue: ")
	annualExpenses := getValue("Annual Expenses: ")
	taxRate := getValue("Tax Rate: ")

	ebt, profit, ratio := calculations(annualRevenue, annualExpenses, taxRate)

	fmt.Printf("Expenses before tax: $%.2f\n", ebt)
	fmt.Printf("Profit after tax: $%.2f\n", profit)
	fmt.Printf("Ratio: %.4f\n", ratio)

}