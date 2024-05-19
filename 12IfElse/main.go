package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount:=25
	var result string
	if loginCount<10{
		result="regular User"
	}else if loginCount>10{
		result="watch out"
	}else{
		result =" something crazy"
	}
	fmt.Println(result)

	if 9%2==0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Num is odd")
	}

	if num:=3; num<10{
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is mor ethan 10")
	}
}

