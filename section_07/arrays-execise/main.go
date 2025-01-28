package main

import "fmt"

func main() {

	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	var hobbies = [3]string{"Games", "Movies", "Electronics"}
	fmt.Println("Mission 1")
	fmt.Println(hobbies)
	fmt.Println()

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	var selectedHobbies = hobbies[1:]
	fmt.Println("Mission 2")
	fmt.Println(hobbies[0])
	fmt.Println(selectedHobbies)
	fmt.Println()

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)

	var slice1 = hobbies[0:2]
	var slice2 = hobbies[:2]
	fmt.Println("Mission 3")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println()

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	var slice3 = slice1[1:cap(slice1)]
	fmt.Println("Mission 4")
	fmt.Println(slice3)
	fmt.Println()

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

	var courseGoals = []string{"Learn-GO", "Compile-crashrepo"}
	fmt.Println("Mission 5")
	fmt.Println(courseGoals)
	fmt.Println()

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Compile-monorepo"
	courseGoals = append(courseGoals, "Create an Agent")
	fmt.Println("Mission 6")
	fmt.Println(courseGoals)
	fmt.Println()

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	fmt.Println("Mission 7: The Bonus!")
	var products = []Product{
		{
			Id:    1,
			Title: "Book",
			Price: 9.99,
		},
		{
			Id:    2,
			Title: "Pen",
			Price: 5.99,
		},
	}

	fmt.Println(products)

	var anotherProduct = Product{
		Id:    3,
		Title: "Eraser",
		Price: 3.99,
	}

	products = append(products, anotherProduct)
	fmt.Println(products)
}

type Product struct {
	Id    int
	Title string
	Price float64
}
