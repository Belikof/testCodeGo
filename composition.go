package main

import "fmt"


type Product struct {
	Name     string
	Price    float64
	Quantity int
}


type Customer struct {
	Name  string
	Email string
	Phone string
}

type Order struct {
	Customer   Customer  
	Products   []Product 
	TotalPrice float64    
}

func (o *Order) calculateTotal() {
	o.TotalPrice = 0
	for _, product := range o.Products {
		o.TotalPrice += product.Price * float64(product.Quantity)
	}
}

func (o *Order) printOrderDetails() {
	fmt.Printf("Заказ для %s (Email: %s, Телефон: %s)\n\n",
		o.Customer.Name, o.Customer.Email, o.Customer.Phone)

	fmt.Println("Список товаров:")
	for i, product := range o.Products {
		fmt.Printf("%d. %s - %.2f x %d = %.2f\n", 
			i+1, product.Name, product.Price, product.Quantity, product.Price*float64(product.Quantity))
	}

	fmt.Printf("\nОбщая сумма заказа: %.2f\n", o.TotalPrice)
}

func main() {
	customer := Customer{
		Name:  "Alice Johnson",
		Email: "alice@example.com",
		Phone: "+123456789",
	}

	products := []Product{
		{Name: "Laptop", Price: 1000.00, Quantity: 1},
		{Name: "Mouse", Price: 25.50, Quantity: 2},
		{Name: "Keyboard", Price: 75.00, Quantity: 1},
	}

	order := Order{
		Customer: customer,
		Products: products,
	}

	order.calculateTotal()

	order.printOrderDetails()
}