package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("working with files")
	content :="Name: sahil\nRollNo: 21103129"

	file,err := os.Create("./myinfo.txt")

	// if err != nil{
	// 	panic(err)
	// }
	checknilError(err)

	length,err :=io.WriteString(file,content)

		checknilError(err)
	fmt.Println("length is: ",length)
	defer file.Close()
	readFile("./myinfo.txt")
	
}

func readFile(filename string){
	databyte ,err :=os.ReadFile(filename)
	checknilError(err)
// file have data [78 97 109 101 58 32 115 97 104 105 108 32 10 82 111 108 108 78 111 58 32 50 49 49 48 51 49 50 57]
//not string format
	fmt.Println("file have data",databyte)

	fmt.Println("text data inside file is\n",string(databyte))
}

func checknilError(err error){
	if err !=nil {
		panic(err)
	}
}