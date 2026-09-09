package main

import "fmt"

// Define the nested struct type
type addition struct {
	phone int
}

// Define the main user struct type
type user struct {
	name            string
	age             int
	metainformation addition
}

func main() {
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["userScore"] = 7
	fmt.Println(myMap)

	myMap2 := map[string]user{
		"jon": {
			name: "Jon",
			age:  30,
			metainformation: addition{
				phone: 123456789,
			},
		},
	}
	fmt.Println(myMap2)
}