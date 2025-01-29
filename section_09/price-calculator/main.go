package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	var taxRates = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {

		outputFilename := fmt.Sprintf("data/result_%0.f.json", taxRate*100)
		fileManager := filemanager.New("data/prices.txt", outputFilename)

		job := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		err := job.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

		cmdManager := cmdmanager.New()
		job = prices.NewTaxIncludedPriceJob(cmdManager, taxRate)
		job.Process()
	}

}
