package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	Aryan:=User{"Aryan","aryan04@.dev",true,16}
	fmt.Println(Aryan)
	fmt.Printf("Aryan details are : %+v\n",Aryan)
	fmt.Printf("Name is : %v\n",Aryan.Name)

}
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
