package main

import "fmt"

func main() {
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["userScore"] = 7
	fmt.Println(myMap)
}
