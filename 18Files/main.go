package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the files")
	content:="This needs to go in a file-LCO.in"
	file,err:=os.Create("./mygoFile.txt")
	checkNilError(err)
	length,err:= io.WriteString(file,content)
	checkNilError(err)
	fmt.Println("The lenght of the file is",length)

	defer file.Close()
	readFile("./mygoFile.txt")
}
func readFile(filename string){
	databyte, err:= os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Data:",string(databyte))
}
func checkNilError(err error){
	if err!=nil{
		panic(err)
	}
}