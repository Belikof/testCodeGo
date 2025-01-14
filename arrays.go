package main

import "fmt"

func main() {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}
	// Формируем сообщение из массива
	message := fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
	fmt.Print(message) // Вывод: Get ready to Go
}