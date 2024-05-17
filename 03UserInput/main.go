package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader:= bufio.NewReader(os.Stdin)

	fmt.Println("Enter the reading for our pizza: ")


	//comma ok || err ok//

	input,_:= reader.ReadString('\n')
	fmt.Println("Thanks for reading," ,input)
	fmt.Printf("Type of this rating is %T", input)

}