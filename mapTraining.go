package main

import "fmt"

func addMap(info map[string][]int) map[string][]int{
	var name string
	var mark int
	var answer string
	for {
		fmt.Println("введите имя:")
		fmt.Scan(&name)
		for {
			fmt.Println("введите оценку:")
			fmt.Scan(&mark)
			info[name] = append(info[name], mark)
			fmt.Println("добавить еще оценок ", name, "?(да/нет)")
			fmt.Scan(&answer)
			if answer == "да" {
				continue
			} else if answer == "нет" {
				break
			}
		}
		fmt.Println("добавить еще ученика?(да/нет)")
		fmt.Scan(&answer)
		if answer == "да" {
			continue
		} else if answer == "нет" {
			break
		}
	}
	return info
}

func main() {
	info := make(map[string][]int)
	info = addMap(info)
	for key, value := range info {
		fmt.Println(key, "-", value)
	}
}