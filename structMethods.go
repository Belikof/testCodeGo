package main

import "fmt"

type Book struct {
	title  string
	author string
	price  int
}


func (b *Book) showInfo() {
	fmt.Println("Title:", b.title)
	fmt.Println("Author:", b.author)
	fmt.Println("Price:", b.price)
}


func (b *Book) setDiscount(percent float64) {
	discountedPrice := float64(b.price) - (float64(b.price) * percent / 100)
	b.price = int(discountedPrice)
	fmt.Println("New Price after discount:", b.price)
}

func main() {

	book := Book{
		title:  "Gogo",
		author: "Gigi",
		price:  100,
	}

	book.showInfo()      
	book.setDiscount(10) 
	book.showInfo()      
}