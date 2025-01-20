package main

import ("fmt"
		
)

// type user struct {
// 	name string
// 	age int
// 	city string
// }

// func show(newUser user) {
// 	fmt.Println(newUser)
// }

// func main() {
// 	newUser := user {
// 		"John",
// 		24,
// 		"Moscow",
// 	}
// 	show(newUser)
// }

type contact struct {
	name  string
	phone string
	email string
}


func addContact(users *[]contact) {
	var name, phone, email string

	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	fmt.Println("Введите телефон:")
	fmt.Scan(&phone)
	fmt.Println("Введите email:")
	fmt.Scan(&email)


	*users = append(*users, contact{name, phone, email})
	fmt.Println("Контакт добавлен!")
}


func show(users []contact) {
	if len(users) == 0 {
		fmt.Println("Список контактов пуст.")
		return
	}
	fmt.Println("Список контактов:")
	for i, user := range users {
		fmt.Printf("%d. Имя: %s, Телефон: %s, Почта: %s\n", i+1, user.name, user.phone, user.email)
	}
}


func menu() {
	fmt.Println(`Выберите действие:
	1 - Добавить контакт
	2 - Показать контакты
	3 - Выход`)
}

func main() {
	var users []contact 
	var answer int

	fmt.Println("Добро пожаловать в книгу контактов!")
	for {
		menu()
		fmt.Scan(&answer)

		switch answer {
		case 1:
			addContact(&users)
		case 2:
			show(users)
		case 3:
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Некорректный ввод, попробуйте снова.")
		}
	}
}