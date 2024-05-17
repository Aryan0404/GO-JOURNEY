package main

import "fmt"

const LoginToken string = "ghabbhhgh"
//initiakization of first letter of variable as capital 
//makes the variable public


func main() {
	var username string = "aryan";
	fmt.Println(username)
	fmt.Printf("Varaible is of type: %T \n",username)


	var isloggedin bool = false;
	fmt.Println(isloggedin)
	fmt.Printf("Varaible is of type: %T \n",isloggedin)

	var smallVal uint8=255
	fmt.Println((smallVal))
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var website = "learn code online"
	fmt.Println(website)

	//no user style

	numberOfuser := 30000.0
	fmt.Println(numberOfuser)

	fmt.Println(LoginToken)
}
