package main

import "fmt"

func main()  {

	fmt.Println("understanding pointers in go")

	var ptr *int
	fmt.Println("value of ptr is ",ptr)  //default value of pointer is  nil

	myusers:=12
	ptr=&myusers  //we have reference myusers variable to ptr & it contains the address of myusers variable

	fmt.Println("value of ptr is ",ptr)
	fmt.Println("value of myusers is ",*ptr)
 
	*ptr=*ptr * 3
	fmt.Println("value of myusers is",*ptr)
}