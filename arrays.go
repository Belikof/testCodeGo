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
// ================================
// func message() string {
// 	arr := [4]string{"ready", "Get", "Go", "to"}
// 	arr[1] = "It's"
// 	arr[0] = "time"
// 	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
// }

// func main() {
// 	fmt.Print(message())
// }
// ===============================
// func message() string {
// 	m := ""
// 	arr := [4]int{1,2,3,4}

// 	for i := 0; i < len(arr); i++ {
// 		arr[i] = arr[i] * arr[i]
// 		m += fmt.Sprintf("%v: %v\n", i, arr[i])
// 	}
// 	return m
// }

// func main() {
// 	fmt.Print(message())
// }
// =============================
// func fillArray(arr [10]int) [10]int {
// 	for i := 0; i < len(arr); i++ {
// 		arr[i] = i + 1
// 	}
// 	return arr
// }

// func opArray(arr [10]int) [10]int {
// 	for i := 0; i < len(arr); i++ {
// 		arr[i] = arr[i] * arr[i]
// 	}
// 	return arr
// }

// func main() {
// 	var arr [10]int
// 	arr = fillArray(arr)
// 	arr = opArray(arr)
// 	fmt.Println(arr)
// }
// ==============================
// func main() {
// 	var arr [10]int
// 	for i := 0; i < len(arr); i++ {
// 		arr[i] = i + 1
// 	}
// 	fmt.Println(arr)
// }