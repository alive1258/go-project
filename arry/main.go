package main

import "fmt"

func main() {
	
	var numbers [6]int
	// numbers[1] = 10
	// fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}
	fmt.Println(numbers)
}