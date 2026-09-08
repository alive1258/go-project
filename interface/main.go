package main

import "fmt"

type Dong struct {}
type Cat struct {}

type Speaker interface {
	speak()
}

func (d *Dong) speak() {
	fmt.Println("woof")
}
func (c *Cat) speak() {
	fmt.Println("meow")
}

func mainSound(s Speaker) {
	s.speak()
}

func main() {
	d := Dong{}
	c := Cat{}

	mainSound(&d)
	mainSound(&c)
}
