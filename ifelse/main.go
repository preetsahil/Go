package main

import "fmt"


func main()  {

	loginCount:=23

	if loginCount<10{
		fmt.Println("loginCount less than 10")
	}else if loginCount>10{
		fmt.Println("loginCount more than 10")
	}else{
		fmt.Println("loginCount equal to 10")
	}

	if 9%2==0{
		fmt.Println("even number")
	}else{
		fmt.Println("odd number")
	}

	if num:=3; num<10{
		fmt.Println("num less than 10")
	}else{
		fmt.Println("num more than 10")
	}

	// if err!=nil{

	// }
}