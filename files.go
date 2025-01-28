package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
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

	data := "Hello world!"
	count, err := file.WriteString(data)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Записано %d байт\n", count)
}