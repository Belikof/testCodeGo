package main

import ("fmt"

)

// func getPassedArgs(minArgs int) []string {
// 	if len(os.Args) < minArgs {
// 		fmt.Printf("At least %v arguments are needed\n", minArgs)
// 		os.Exit(1)
// 	}
// 	var args []string
// 	for i := 1; i < len(os.Args); i++ {
// 		args = append(args, os.Args[i])
// 	}
// 	return args
// }

// func findLongest(args []string) string {
// 	var longest string
// 	for i := 0; i < len(args); i++ {
// 		if len(args[i]) > len(longest) {
// 			longest = args[i]
// 		}
// 	}
// 	return longest
// }

// func main() {
// 	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
// 		fmt.Println("The longest word passed was:", longest)
// 	} else {
// 		fmt.Println("There was an error")
// 		os.Exit(1)
// 	}
// }

// func getPassedArgs() []string {
// 	var args []string
// 	for i := 1; i < len(os.Args); i++ {
// 		args = append(args, os.Args[i])
// 	}
// 	return args
// }

// func getLocals(extraLocals []string) []string {
// 	var locales []string
// 	locales = append(locales, "en_US", "fr_FR")
// 	locales = append(locales, extraLocals...)
// 	return locales
// }

// func main() {
// 	locales := getLocals(getPassedArgs())
// 	fmt.Println("Locales to use:", locales)
// }

// func message() string {
// 	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	m := fmt.Sprintln("First:", s[0], s[0:1], s[:1])
// 	m += fmt.Sprintln("Last:", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
// 	m += fmt.Sprintln("First 5:", s[:5])
// 	m += fmt.Sprintln("Last 4 :", s[5:])
// 	m += fmt.Sprintln("Middle 5:", s[2:7])
// 	return m
// }

// func main() {
// 	fmt.Print(message())
// }

func genSlices() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)
	return s1, s2, s3
}

func main() {
	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
}