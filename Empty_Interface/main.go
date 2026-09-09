package main

import "fmt"

func Print(data any){
	
	strDta,ok := data.(string)
	if !ok {
		fmt.Println("Data is not string:")

	}else{
	    fmt.Println(strDta)
	}
}

func main(){
	Print("10")
}
