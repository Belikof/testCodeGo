package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// buffer := make([]byte, 1024)
	// count, err := file.Read(buffer)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }

	// fmt.Printf("Read %d byte: %s\n", count, buffer[:count])

	// data := "\nHello world!"
	// _, err = file.WriteString(data)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }

	fmt.Printf("Succesfully")
}