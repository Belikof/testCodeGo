package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func main() {
	user := User {
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}