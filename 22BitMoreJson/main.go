package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses:= []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}
	finalJson,err:= json.MarshalIndent(lcoCourses,"","\t")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}
func DecodeJson(){
	jsonDataFromWeb:=[]byte(`

{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "Tags": [
                        "web-dev",
                        "js"
                ]
        }

	`)
	 var lcoCourse course

	checkValid:=json.Valid(jsonDataFromWeb)
	if checkValid{
		fmt.Println("Json was valid ")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)
		//following is the struct format conversion
		fmt.Printf("%#v\n", lcoCourse)

	}else{
		fmt.Println("Not in json format")
	}
	//key-value format
	var myOnlineData map[string] interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	for k,v:=range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is %T\n",k,v,v)
	}

}