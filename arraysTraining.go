package main

import "fmt"

func addArr(arr[10]int) [10]int {
	fmt.Println("Введите 10 чисел:")
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func orderArr(arr[10]int) [10]int {
	for {
		swapped := false
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				swapped = true
			} else {
				continue
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func reverseArr(arr[10]int) [10]int {
	for {
		swapped := false
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] < arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				swapped = true
			} else {
				continue
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func sumArr(arr [10]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i] // Добавляем каждый элемент массива к сумме
	}
	return sum
}

func minArr(arr [10]int) int {
	minimum := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < minimum {
			minimum = arr[i]
		} else {
			continue
		}
	}
	return minimum
}

func maxArr(arr [10]int) int {
	maximum := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maximum {
			maximum = arr[i]
		} else {
			continue
		}
	}
	return maximum
}

func main() {
	var orderArray[10] int
	var reverseArray[10] int
	var sumArray int
	var minArray int
	var maxArray int
	orderArray = addArr(orderArray)
	orderArray = orderArr(orderArray)
	reverseArray = reverseArr(orderArray)
	sumArray = sumArr(orderArray)
	minArray = minArr(orderArray)
	maxArray = maxArr(orderArray)
	fmt.Println("От меньшего к большему:", orderArray)
	fmt.Println("От большего к меньшему:", reverseArray)
	fmt.Println("Сумма значений массива:", sumArray)
	fmt.Println("минимальное значение:", minArray)
	fmt.Println("Максимальное значение:", maxArray)
}

