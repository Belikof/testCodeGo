package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// buffer := make([]byte, 1024)
	// count, err := file.Read(buffer)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }

	// fmt.Printf("Read %d byte: %s\n", count, buffer[:count])

	data := "\nHello world!"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Succesfully")
}