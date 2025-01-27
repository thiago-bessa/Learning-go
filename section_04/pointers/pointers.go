package main

import "fmt"

func main() {
	var age int = 39
	var agePointer *int = &age

	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years:", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
