package main

import "fmt"

// func main() {
// 	a := [4]int{1, 2, 3, 4}
// 	reverse(&a)
// 	fmt.Println(a)
// }

// func reverse(arr *[4]int) {
// 	for index, value := range *arr {
// 		(*arr)[len(arr) - 1 - index] = value
// 	}
// }

func newName(name *string) {
	*name = "Linda"
}

func newAge(age *int) {
	*age = 30
}

func newEmail(email *string) {
	*email = "belsta.nuk@gmail.com"
}

func main() {
	var name = "Alice"
	var age = 25
	var email = "alice@example.com"
	fmt.Println(name, age, email)
	newName(&name)
	newAge(&age)
	newEmail(&email)
	fmt.Println(name, age, email)
}
