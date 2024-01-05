package main

import "fmt"

//public 
const Token string="fsafas"

func main()  {

	var username string = "sahil"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var a uint8 = 255
	fmt.Println(a)
	fmt.Printf("Variable is of type: %T \n",a)

	var decimal float32 = 342.32131
	fmt.Println(decimal)
	fmt.Printf("Variable is of type: %T \n",decimal)

	//default value and aliases

	var users string
	fmt.Println(users)
	fmt.Printf("Variable is of type: %T \n",users)

	//implicit type lexer decide the type of variable

	var score = 100
	fmt.Println(score)
	fmt.Printf("Variable is of type: %T \n",score)

	//no var type declaration

	numberOfUsers :=3000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n",numberOfUsers)

	fmt.Println(Token)
	fmt.Printf("Variable is of type: %T \n",Token)

}