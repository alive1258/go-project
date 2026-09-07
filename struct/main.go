package main

import "fmt"

type addition struct {
	phone   int
	address string
}
type user struct {
	name string
	age  int
	metainformation addition
}

func main() {
	// jon := user{name: "Jon", age: 30}
	// jon.age = 31
	// jon.name = "Jon Snow"
	// println(jon.name, jon.age)

	// var user1 user
	// user1.name = "Arya"
	// user1.age = 18
	// println(user1.name, user1.age)

	jhon := user{name: "Jon",
		age: 30,
		metainformation: addition{
			phone:   123456789,
			address: "Winterfell"}}
	fmt.Println(jhon)
}
