package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	
	data := "Hello, world!"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error for write:", err)
		return
	}

	file, err = os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error for open:", err)
		return
	}
	defer file.Close()

	additionalData := "\nAdding more data..."
	_, err = file.WriteString(additionalData)
	if err != nil {
		fmt.Println("Error for adding:", err)
		return
	}

	file, err = os.Open("data.txt")
	if err != nil {
		fmt.Println("Error for opening:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	count, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error for reading:", err)
		return
	}

	fmt.Println(buffer[:count])

	err = os.Remove("data.txt")
	if err != nil {
		fmt.Println("Error for remove:", err)
		return
	}
	
}