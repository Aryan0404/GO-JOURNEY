package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()
	result:=adder(3,5)
	proResult,strResult:=proAdder(1,2,3,4,5,6)
	//can ignore strResult by adding _
	fmt.Println("The proresult is " ,proResult)
	fmt.Println(strResult)
	fmt.Println(result)

}
//we cant nest functions
//But we can certainly return more then 1 value in go functions
func proAdder(val ...int) (int,string)  {
	total:=0
	for _,i := range val{
		total+=i
	}
	

	return total,"The total is correct"
}
func adder(valOne int,valTwo int) int{
	return valOne+valTwo
}
func greeter()  {
	fmt.Println("Hey there")
}

