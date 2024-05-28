package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web erb video - LCO")
	performGetRequst()
	performPostRequest()
performFormRequest()
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
func performPostRequest(){
		const myUrl = "http://localhost:8000/post"
		requestBody := strings.NewReader(`
        {
            "coursename":"Let's go with golang",
            "price": 0,
            "platform":"learnCodeOnline.in"
        }
    `)
	response,err := http.Post(myUrl,"application/json",requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))

	
}
func performFormRequest(){
	const myUrl= "http://localhost:8000/postform"

	data:= url.Values{

	}
	data.Add("firstName","Hithesh")
	data.Add("LastName","Choudhary")
	data.Add("email","hithest.go@dev")
	response,err:=http.PostForm(myUrl,data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))		
}