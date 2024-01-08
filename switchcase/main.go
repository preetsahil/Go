package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Println("welcome to switchcase")

	diceNumber:=rand.Intn(6)+1

	switch diceNumber{
	case 1:
		fmt.Println("you can open and dice value is 1")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll the dice again")
	default:
		fmt.Println("what was that")
	}
	
}