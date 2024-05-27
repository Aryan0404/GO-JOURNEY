package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web erb video - LCO")
	performGetRequst()

}
func performGetRequst(){
	const myurl = "http://localhost:8000/get"
	response,err:=http.Get(myurl)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)
	var responseString strings.Builder
	content,_ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println(byteCount)
	

}