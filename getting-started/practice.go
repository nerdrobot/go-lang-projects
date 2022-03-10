package main

import "fmt"

func main() {
	var firstName string = "Rao"
	lastName := "Hassan"

	fmt.Println(firstName)
	fmt.Println(lastName)

	var currentYear int
	currentYear = 2022
	birthYear := 1998

	fmt.Println(currentYear)
	fmt.Println(birthYear)

	age := currentYear - birthYear

	fmt.Println(age)

	currentYear = currentYear + 1

	age = currentYear - birthYear

	fmt.Println(age)

	pi := 3.14
	radius := 5

	circumference := 2 * pi * float64(radius)

	outputString := fmt.Sprintf("For a radius of %v, the circle cirumference is %0.2f", radius, circumference)

	fmt.Println(outputString)

}
