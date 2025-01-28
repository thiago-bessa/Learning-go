package main

import (
	"fmt"
)

func main() {

	printSomething(1)
	printSomething(1.1)
	printSomething(true)
	printSomething("Welcome to Interfaces!")

	printSomething2(1)
	printSomething2(1.1)
	printSomething2(true)
	printSomething2("Welcome to Interfaces!")

	resultInt := add(1, 2)
	fmt.Printf("1 + 2 = %v\n", resultInt)
}

func add[T int | float64 | string](a T, b T) T {
	return a + b
}

func printSomething(value interface{}) {

	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println(value)
	default:
		fmt.Println("Type unknown")
	}
}

func printSomething2(value interface{}) {

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println(stringVal)
	}

	_, ok = value.(bool)

	if ok {
		fmt.Println("This is a boolean!")
	}
}
