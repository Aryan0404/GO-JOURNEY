package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling Urls in goLang")
	fmt.Println(myUrl)
	result,_ :=url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams:=result.Query()
	fmt.Printf("The type of params here is : %T\n",qparams)
	for _,val:= range qparams {
		fmt.Println("Param value is :",val)
	} 
	partsOfUrl:= &url.URL{
		Scheme:	"https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherUrl:=partsOfUrl.String()
	fmt.Println(anotherUrl)
}