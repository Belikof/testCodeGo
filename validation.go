package main

import (
	"fmt"
	"errors"
	"strings"
)


type User struct {
	name string
	age int
	email string
}

func newUser (name string, age int, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("имя не может быть пустым")
	} else if len(name) < 3 {
		return nil, errors.New("длина должна быть не менее 3 символов")
	}
	if age < 18 || age > 100 {
		return nil, errors.New("не подходящий возраст")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return nil, errors.New("некорректный email: должен содержать '@' и '.'")
	}
	
	return &User {
		name: name,
		age: age,
		email: email,
	}, nil
}

func printUserInfo(user *User) {
	fmt.Println("Имя", user.name)
	fmt.Println("Возраст", user.age)
	fmt.Println("Почта", user.email)
}

func main() {
	user, err := newUser("Alice", 25, "alice@example.com")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	printUserInfo(user)
}