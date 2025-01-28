package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	productNames[2] = "A Carpet"

	fmt.Println(productNames)
	fmt.Println(prices)

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99

	hightlightedPrices := featuredPrices[:1]

	fmt.Println(hightlightedPrices)
	fmt.Println(len(hightlightedPrices), cap(hightlightedPrices))

	hightlightedPrices = hightlightedPrices[:3]

	fmt.Println(hightlightedPrices)
	fmt.Println(len(hightlightedPrices), cap(hightlightedPrices))
}
