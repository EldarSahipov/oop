package entity

import (
	"fmt"
)

type Cart struct {
	Products []Product
}

func NewCart() *Cart {
	return &Cart{Products: make([]Product, 0)}
}

func (c *Cart) GetInfo() {
	for _, v := range c.Products {
		fmt.Println(v)
	}
}

func (c *Cart) AddProduct(p Product) bool {
	// Проверяем, нет ли уже такого продукта в корзине
	if c.containsProduct(&p) {
		fmt.Println("Продукт уже существует в корзине")
		return false
	}

	c.Products = append(c.Products, p)
	fmt.Println("Продукт добавлен!")
	return true
}

func (c *Cart) DeleteProduct(p *Product) bool {
	productIndex := -1
	for i, v := range c.Products {
		if v.ID == p.ID {
			productIndex = i
			break
		}
	}

	if productIndex == -1 {
		fmt.Println("Продукт не найден")
		return false
	}

	c.Products = append(c.Products[:productIndex], c.Products[productIndex+1:]...)
	fmt.Println("Продукт удален!")
	return true
}

func (c *Cart) containsProduct(p *Product) bool {
	for _, v := range c.Products {
		if v.ID == p.ID {
			return true
		}
	}
	return false
}

func (c *Cart) TotalPrice() int {
	sum := 0
	for _, v := range c.Products {
		sum += v.Price * v.Number
	}
	return sum
}
