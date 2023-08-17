package main

import (
	"OOP/entity"
)

func main() {
	product1 := entity.NewProduct(1516, "Сок", 119, 50)
	product2 := entity.NewProduct(1564, "Хлеб", 155, 40)
	product3 := entity.NewProduct(9465, "Шоколад", 200, 10)
	listProducts := entity.NewCart()
	listProducts.Products = append(listProducts.Products, *product1, *product2, *product3)
	listProducts.GetInfo()
	product1.GetInfo()
	order := entity.NewOrder(listProducts)
	order.GetInfo()
	order.PayOrder()
	order.GetInfo()
}
