package main

import "fmt"

func main() {
	fmt.Println("Welcome to this class of pointers")

	myNumber:=23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is " ,ptr)
	fmt.Println("Value of actual pointer is " ,*ptr)

	*ptr=*ptr+2
	fmt.Println("New value is ",myNumber)
}