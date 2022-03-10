package main

import (
	"fmt"
)

type Product struct {
	id    int
	title string
	price float64
}

func NewProduct(id int, title string, price float64) *Product {
	return &Product{id, title, price}
}

func main() {
	// Step 1
	hobbies := [3]string{"Learning New Programming Languages", "Golf", "Gym"}
	fmt.Println(hobbies)

	// Step 2
	fmt.Println(hobbies[0])
	slicedArray := hobbies[1:]
	fmt.Println(slicedArray)

	// Step 3
	newSlice := hobbies[:2]
	//differentWayOfSlicing := hobbies[0:2]

	// Step 4
	anotherSlice := newSlice
	anotherSlice[0] = hobbies[0]
	anotherSlice[1] = hobbies[2]
	fmt.Println(anotherSlice)

	// Step 5
	courseGoals := []string{"Learn Go lang syntax", "Improve Understanding"}
	// Step 6
	courseGoals[1] = "Become an expert in Go"

	courseGoals = append(courseGoals, "Understand different things to achieve with Go")

	fmt.Println(courseGoals)

	// Step 7
	productsList := []Product{*NewProduct(1, "Test", 1.99), *NewProduct(2, "Test-2", 2.99)}

	productsList = append(productsList, *NewProduct(3, "test-3", 4.99))

	fmt.Println(productsList)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
