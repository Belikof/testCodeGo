package main

import "fmt"

// type Book struct {
// 	title  string
// 	author string
// 	price  int
// }


// func (b *Book) showInfo() {
// 	fmt.Println("Title:", b.title)
// 	fmt.Println("Author:", b.author)
// 	fmt.Println("Price:", b.price)
// }


// func (b *Book) setDiscount(percent float64) {
// 	discountedPrice := float64(b.price) - (float64(b.price) * percent / 100)
// 	b.price = int(discountedPrice)
// 	fmt.Println("New Price after discount:", b.price)
// }

// func main() {

// 	book := Book{
// 		title:  "Gogo",
// 		author: "Gigi",
// 		price:  100,
// 	}

// 	book.showInfo()      
// 	book.setDiscount(10) 
// 	book.showInfo()      
// }


type BankAccount struct {
	owner string
	accountNumber string
	balance float64
}

func newBankAccount (owner, accountNumber string, balance float64) *BankAccount {
	return &BankAccount {
		owner: owner,
		accountNumber: accountNumber,
		balance: balance,
	}
}

func (ba *BankAccount) deposit(amount float64) {
	ba.balance += amount
	fmt.Println("Пополняем счет на", amount)
	fmt.Println("Итого", ba.balance)
}

func (ba *BankAccount) withdraw(amount float64) {
	if ba.balance > amount {
		ba.balance -= amount
		fmt.Println("Снимаем со счета", amount)
		fmt.Println("Итого", ba.balance)
	} else if ba.balance < amount {
		fmt.Println("На вашем счету недостаточно средств!")
	}
}

func (ba *BankAccount) showBalance() {
	fmt.Println("Текущий баланс:", ba.balance)
}

func main() {
	bankAccount := newBankAccount(owner, accountNumber, balance)

	bankAccount.deposit(500)
	bankAccount.withdraw(200)
	bankAccount.showBalance()
}