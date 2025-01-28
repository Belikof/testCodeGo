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
	// user := User {
	// 	Name:  "Alice",
	// 	Age:   30,
	// 	Email: "alice@example.com",
	// }

	// jsonData, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("Ошибка кодирования JSON:", err)
	// 	return
	// }

	// fmt.Println(string(jsonData))
	jsonString := `{"name":"Alice","age":30,"email":"alice@example.com"}`

	var user User

	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON", err)
		return
	}

	fmt.Println("Имя:", user.Name)
	fmt.Println("Возраст:", user.Age)
	fmt.Println("Email:", user.Email)
}