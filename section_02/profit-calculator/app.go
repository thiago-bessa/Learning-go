package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var err error

	revenue, err = getUserInput("Revenue")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err = getUserInput("Expenses")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err = getUserInput("taxRate")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("")

	var ebt, profit, ratio = calculateEarningsAndProfit(revenue, expenses, taxRate)

	writeCalculationsToFile(ebt, profit, ratio)

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserInput(text string) (float64, error) {
	var value float64

	fmt.Print(text + ": ")
	fmt.Scan(&value)

	if value <= 0.0 {
		return 0, errors.New("invalid value: should be greater than 0")
	}

	return value, nil
}

func calculateEarningsAndProfit(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	var ebt = revenue - expenses
	var profit = ebt * (1 - taxRate/100)
	var ratio = ebt / profit

	return ebt, profit, ratio
}

func writeCalculationsToFile(ebt, profit, ratio float64) {
	var balanceText = fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("calculated.txt", []byte(balanceText), 0644)
}
