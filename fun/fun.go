package main

import "fmt"

func makeCoffee(kind string){
	fmt.Printf("making  %s coffee.........\n", kind)
}

func main(){
	makeCoffee("balck")
	makeCoffee("Cold")
}