package main

import "fmt"

func addValue(arr [4]string, sli[]int) ([4]string, []int) {
	fmt.Println("Заполните массив данными (4 значения типа string):")
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Заполните слайс данными (8 значений типа int):")
	for i := 0; i < len(sli); i++ {
		fmt.Scan(&sli[i])
	}
	return arr, sli
}

func show(arr [4]string, sli[]int) (string, int) {
	return arr[2], sli[2]
}

func redInfo(arr [4]string, sli[]int) ([4]string, []int) {
	fmt.Println("Введите новое значение для третьего индекса в массиве(строка):")
	fmt.Scan(&arr[3])
	fmt.Println("Введите новое значение для третьего индекса в слайсе(целое число):")
	fmt.Scan(&sli[3])
	return arr, sli
}

func main() {
	var arr [4]string
	var sli = make([]int, 8)
	arr, sli = addValue(arr, sli)
	fmt.Println(show(arr, sli))
	fmt.Println(redInfo(arr, sli))
}