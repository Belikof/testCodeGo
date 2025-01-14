package main

import "fmt"

// func main() {
// 	arr := [...]string{
// 		"ready",
// 		"Get",
// 		"Go",
// 		"to",
// 	}
// 	// Формируем сообщение из массива
// 	message := fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
// 	fmt.Print(message) // Вывод: Get ready to Go
// }

func message() string {
	arr := [4]string{"ready", "Get", "Go", "to"}
	arr[1] = "It's"
	arr[0] = "time"
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Print(message())
}