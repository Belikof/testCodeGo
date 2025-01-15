package main

import "fmt"

// func addArr(arr[10]int) [10]int {
// 	fmt.Println("Введите 10 чисел:")
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i])
// 	}
// 	return arr
// }

// func orderArr(arr[10]int) [10]int {
// 	for {
// 		swapped := false
// 		for j := 0; j < len(arr) - 1; j++ {
// 			if arr[j] > arr[j + 1] {
// 				arr[j], arr[j + 1] = arr[j + 1], arr[j]
// 				swapped = true
// 			} else {
// 				continue
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// func reverseArr(arr[10]int) [10]int {
// 	for {
// 		swapped := false
// 		for j := 0; j < len(arr) - 1; j++ {
// 			if arr[j] < arr[j + 1] {
// 				arr[j], arr[j + 1] = arr[j + 1], arr[j]
// 				swapped = true
// 			} else {
// 				continue
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// func sumArr(arr [10]int) int {
// 	sum := 0
// 	for i := 0; i < len(arr); i++ {
// 		sum += arr[i] // Добавляем каждый элемент массива к сумме
// 	}
// 	return sum
// }

// func minArr(arr [10]int) int {
// 	minimum := arr[0]
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] < minimum {
// 			minimum = arr[i]
// 		} else {
// 			continue
// 		}
// 	}
// 	return minimum
// }

// func maxArr(arr [10]int) int {
// 	maximum := arr[0]
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] > maximum {
// 			maximum = arr[i]
// 		} else {
// 			continue
// 		}
// 	}
// 	return maximum
// }

// func main() {
// 	var orderArray[10] int
// 	var reverseArray[10] int
// 	var sumArray int
// 	var minArray int
// 	var maxArray int
// 	orderArray = addArr(orderArray)
// 	orderArray = orderArr(orderArray)
// 	reverseArray = reverseArr(orderArray)
// 	sumArray = sumArr(orderArray)
// 	minArray = minArr(orderArray)
// 	maxArray = maxArr(orderArray)
// 	fmt.Println("От меньшего к большему:", orderArray)
// 	fmt.Println("От большего к меньшему:", reverseArray)
// 	fmt.Println("Сумма значений массива:", sumArray)
// 	fmt.Println("минимальное значение:", minArray)
// 	fmt.Println("Максимальное значение:", maxArray)
// }
func newArr(arr [15]int) [15]int {
	arr = [15]int{}
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func oddNumb(arr [15]int) []int{
	oddArr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 0 {
			oddArr = append(oddArr, arr[i])
		} else {
			continue
		}
	} 
	return oddArr
}

func evenNumb(arr [15]int) []int{
	evenArr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 != 0 {
			evenArr = append(evenArr, arr[i])
		} else {
			continue
		}
	} 
	return evenArr
}

func mathArr(arr[15]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum / len(arr)
}

func positiveNumb(arr[15]int) int {
	positive := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive += 1
		} else {
			continue
		}
	}
	return positive
}

func negativeNumb(arr[15]int) int {
	negative := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative += 1
		} else {
			continue
		}
	}
	return negative
}

func reverArr(arr [15]int) [15]int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func main() {
	var newArray [15]int
	var oddArray []int
	var evenArray []int
	var mathArray int
	var positiveArray int
	var negativeArray int
	var reverseArray [15]int
	newArray = newArr(newArray)
	oddArray = oddNumb(newArray)
	evenArray = evenNumb(newArray)
	mathArray = mathArr(newArray)
	positiveArray = positiveNumb(newArray)
	negativeArray = negativeNumb(newArray)
	reverseArray = reverArr(newArray)
	fmt.Println("Массив чисел:", newArray)
	fmt.Println("Четные числа:", oddArray)
	fmt.Println("Не четные числа:", evenArray)
	fmt.Println("Среднее арифметическое:", mathArray)
	fmt.Println("Положительные числа:", positiveArray)
	fmt.Println("Отрицательные числа:", negativeArray)
	fmt.Println("Перевернутый массив:", reverseArray)
}