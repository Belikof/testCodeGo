package main

import (
	"fmt"
)


func menu() {
	fmt.Println(`Добро пожаловать в меню закладок! Что вы хотите сделать?
	- 1. Посмотреть закладки
	- 2. Добавить закладку
	- 3. Удалить закладку
	- 4. Выход`)
}

func add(bookmark map[string]int) map[string]int {
	fmt.Println("Вы в разделе добавления закладок.")
		for {
			var name string
			var page int
			var answer string
			fmt.Println("Введите название книги:")
			fmt.Scan(&name)
			fmt.Println("Введите номер страницы:")
			fmt.Scan(&page)
			bookmark[name] = page
			fmt.Println("Вы добавили:", name, "-", page)
			fmt.Println("Добавить еще?(да/нет)?")
			fmt.Scan(&answer)
			if answer == "да" {
				continue
			} else if answer == "нет" {
				break
			}
		}
	return bookmark
}

func deleting(bookmark map[string]int) map[string]int {
	var name string
	var answer string
	for {
		fmt.Println("Введите название книги, которую хотите удалить:")
		fmt.Scan(&name)
		delete(bookmark, name)
		fmt.Println("Удалить еще?(да/нет)")
		fmt.Scan(&answer)
		if answer == "да" {
			continue
		} else if answer == "нет" {
			break
		}
	}
	return bookmark
}



func main() {
	bookmark := make(map[string]int)
	for {
		var menuChoose int
		menu()
		fmt.Scan(&menuChoose)
		if menuChoose == 1 {
			for key, value := range bookmark {
				fmt.Println(key, "-", value)
				fmt.Println("==============")
			}
		} else if menuChoose == 2 {
			bookmark = add(bookmark)
		} else if menuChoose == 3 {
			bookmark = deleting(bookmark)
		} else if menuChoose == 4 {
			break
		}
	}
}