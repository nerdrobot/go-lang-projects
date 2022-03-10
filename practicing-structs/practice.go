package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func NewProduct(id string, title string, description string, price float64) *Product {
	return &Product{id, title, description, price}
}

func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func main() {
	createdProduct := getProduct()
	createdProduct.printData()
	createdProduct.store()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data")
	fmt.Println("------------------------------")
	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)

	return product

}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func (product *Product) store() {
	file, _ := os.Create(product.id + ".txt")
	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDesctiption %v\nPrice %.2f\n", product.id, product.title, product.description, product.price)
	file.WriteString(content)
	file.Close()
}
