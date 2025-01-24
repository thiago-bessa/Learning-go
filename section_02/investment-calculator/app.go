package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Println("")

	var futureValue, futureRealValue = calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var formattedFV = fmt.Sprintf("Future Value: $%.2f\n", futureValue)
	fmt.Print(formattedFV)

	//var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	var formattedRFV = fmt.Sprintf("Future Real Value (adjusted for Inflation): $%.2f\n", futureRealValue)
	fmt.Print(formattedRFV)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}
