package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitList = []string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fruitList = append(fruitList,"Mango","Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)


	highScores :=make([]int,4)
	highScores[0]=231
	highScores[1]=232
	highScores[2]=233
	
	highScores[3]=234
	//highScores[4]=235//this will crash 
	//thorugh indexing will can only define elements for the 
	//already defined number in make


	fmt.Println(highScores)
	highScores=append(highScores, 555,666,721)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted((highScores)))//tells wether array is sorted or not



	//how to remove a value from a slice based on index

	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)
	var index int = 2;
	courses = append(courses[:index] , courses[index+1:]...)
	fmt.Println(courses)
	


}