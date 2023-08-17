package entity

import "fmt"

type Product struct {
	ID     int
	Name   string
	Price  int
	Number int
}

func NewProduct(ID int, name string, price int, number int) *Product {
	return &Product{ID, name, price, number}
}

func (p Product) GetInfo() {
	fmt.Printf("Name: %s, Price: %d, Number: %d\n", p.Name, p.Price, p.Number)
}
