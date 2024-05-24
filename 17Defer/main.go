package main

import "fmt"

func main() {
	
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello there")
	defer myDefer()
}
//print Sequence: Hello 4 3 2 1 0 there two one world
//defer follows LIFO sequence

func myDefer(){
	for i:=0 ; i<5; i++{
		defer fmt.Println(i)
	}
}