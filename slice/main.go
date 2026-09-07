package main

import "fmt"

func main() {
	var numbers = [6]int{1, 2, 3, 4, 5}
	slice := numbers[1:4]
	fmt.Println(slice)
}