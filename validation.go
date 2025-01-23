package main

import (
	"fmt"
	"errors"
	"time"
	"strings"
)


// type User struct {
// 	name string
// 	age int
// 	email string
// }

// func newUser (name string, age int, email string) (*User, error) {
// 	if name == "" {
// 		return nil, errors.New("имя не может быть пустым")
// 	} else if len(name) < 3 {
// 		return nil, errors.New("длина должна быть не менее 3 символов")
// 	}
// 	if age < 18 || age > 100 {
// 		return nil, errors.New("не подходящий возраст")
// 	}
// 	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
// 		return nil, errors.New("некорректный email: должен содержать '@' и '.'")
// 	}
	
// 	return &User {
// 		name: name,
// 		age: age,
// 		email: email,
// 	}, nil
// }

// func printUserInfo(user *User) {
// 	fmt.Println("Имя", user.name)
// 	fmt.Println("Возраст", user.age)
// 	fmt.Println("Почта", user.email)
// }

// func main() {
// 	user, err := newUser("Alice", 25, "alice@example.com")
// 	if err != nil {
// 		fmt.Println("Ошибка:", err)
// 		return
// 	}
// 	printUserInfo(user)
// }

type Car struct {
	brand string
	year int
	licencePlate string
}

func newCar (brand string, year int, licencePlate string) (*Car, error) {
	if brand == "" || len(brand) < 2 {
		return nil, errors.New("марка  машинв не может быть пустой или содержать меньше 2 символов")
	}
	if year < 1886 || year > time.Now().Year() {
		return nil, errors.New("не подходящий год выпуска")
	}
	if len(licencePlate) < 6 || !strings.ContainsAny(licencePlate, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") || !strings.ContainsAny(licencePlate, "1234567890") {
		return nil, errors.New("неправильный формат лицензии: должен содержать буквы и цифры")
	}

	return &Car{
		brand: brand, 
		year: year,
		licencePlate: licencePlate,
	}, nil
}

func printCarInfo(car *Car) {
	fmt.Println("Марка", car.brand)
	fmt.Println("Год", car.year)
	fmt.Println("Лицензия", car.licencePlate)
}

func main() {
	car, err := newCar("toyota", 1900, "asdasfasf")
	if err != nil {
		fmt.Println("Ошибка", err)
		return 
	}
	printCarInfo(car)
}