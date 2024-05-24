package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	days:=[]string{"Sunday","Monday","Tuesday","Wednesday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("The day is %v\n",days[i])
	// }

	// for i:= range days {
	// 		fmt.Println(days[i])
	// }

	// for i,day:=range days{
	// 	fmt.Printf("Index is %v and day is %v\n",i,day)
	// }
	for _,day:=range days{
		fmt.Printf("Index is  and day is %v\n",day)
	}
	rogueValue:=1
	for rogueValue<10{
		if rogueValue==2 {
			goto lco
		}
		if rogueValue==5{
			rogueValue++;
			continue
		}
		if rogueValue==9{
			break
		}
		fmt.Println("Rogue value is ",rogueValue)
		rogueValue++
	}

	lco:
	fmt.Println(rogueValue)
}
