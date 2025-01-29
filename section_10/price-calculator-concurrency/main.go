package main

import (
	"fmt"

	"example.com/price-calculator-concurrency/filemanager"
	"example.com/price-calculator-concurrency/prices"
)

func main() {
	var taxRates = []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {

		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		outputFilename := fmt.Sprintf("data/result_%0.f.json", taxRate*100)
		fileManager := filemanager.New("data/prices.txt", outputFilename)

		job := prices.NewTaxIncludedPriceJob(fileManager, taxRate)

		go job.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
