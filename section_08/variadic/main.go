package main

import "fmt"

func main() {
	numbers := []int{10, 15, 40, -5}
	sum := sumNumbers(1, 10, 15, 40, -5)
	anotherSum := sumNumbers(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumNumbers(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}
