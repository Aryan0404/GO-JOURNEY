package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switcg operator")

	rand.Seed(time.Now().UnixNano())
	diceNumber:=rand.Intn(6)+1;
	fmt.Println("Value of data is ",diceNumber)
	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("You can move 2 spot ")

	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
	case 5: 
		fmt.Println("You can move 5 spots")
	default:
		fmt.Println("Hey there")
	}
}