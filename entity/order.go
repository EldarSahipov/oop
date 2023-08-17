package entity

import "fmt"

type Order struct {
	Cart   Cart
	Status string
}

func NewOrder(cart *Cart) *Order {
	return &Order{*cart, "Оформляется"}
}

func (o Order) GetInfo() {
	fmt.Println("Статус заказа: ", o.Status)
	fmt.Println("Итого: ", o.Cart.TotalPrice())
	fmt.Println("Продукты: ")
	for i, v := range o.Cart.Products {
		fmt.Printf("%d) Name: %s, Price: %d, Number: %d\n", i+1, v.Name, v.Price, v.Number)
	}
}

func (o *Order) PayOrder() bool {
	o.Status = "Оплачено"
	return true
}
