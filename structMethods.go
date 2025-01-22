package main

import (
	"errors"
	"fmt"
	"unicode"
)

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

func newBankAccount (owner, accountNumber string, balance float64) (*BankAccount, error) {
	if len(owner) == 0 {
		return nil, errors.New("имя владельца не может быть пустым")
	}

	for _, char := range accountNumber {
		if !unicode.IsDigit(char) {
			return nil, errors.New("номер счета должен содержать только цифры")
		}
	}

	if balance < 0 {
		return nil, errors.New("баланс не может быть отрицательным")
	}

	return &BankAccount {
		owner: owner,
		accountNumber: accountNumber,
		balance: balance,
	}, nil
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
	//корректные данные
	bankAccount, err := newBankAccount("owner", "123456", 1000.0)
	if err !=nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("Счет успешно создан для:", bankAccount.owner)

	//некорректные данные
	bankAccount, err = newBankAccount("", "123abc", -100)
	if err != nil {
		fmt.Println("Ошибка:", err)  // Выводит: "Ошибка: имя владельца не может быть пустым"
		return
	}
	bankAccount.deposit(500)
	bankAccount.withdraw(200)
	bankAccount.showBalance()
}