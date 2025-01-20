package main

import ("fmt"
		
)

type user struct {
	name string
	age int
	city string
}

func show(newUser user) {
	fmt.Println(newUser)
}

func main() {
	newUser := user {
		"John",
		24,
		"Moscow",
	}
	show(newUser)
}