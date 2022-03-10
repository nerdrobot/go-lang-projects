package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func NewProduct(id int, name string, desc string, price float64, sku string, createdOn string, updatedOn string) *Product {
	return &Product{id, name, desc, price, sku, createdOn, updatedOn, ""}
}

func GetProducts() []*Product {
	return productsList
}

var productsList = []*Product{
	NewProduct(1, "Latte", "Frothy Milky Coffee", 2.45, "abc323", time.Now().UTC().String(), time.Now().UTC().String()),
	NewProduct(1, "Latte", "Frothy Milky Coffee", 2.45, "abc323", time.Now().UTC().String(), time.Now().UTC().String()),
}
