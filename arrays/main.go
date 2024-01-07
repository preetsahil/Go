package main

import "fmt"

func main()  {

	fmt.Println("welcome to arrays")

	var players[4]string
	players[0]="messi"
	players[1]="cr7"
	players[3]="luka"

	fmt.Println("array is :",players)
	fmt.Println("len of array is:",len(players))

	var lang=[5]string{"GO","JAVA","C++","JS","JSX"}
	fmt.Println(lang)
	fmt.Println(len(lang))
	
	
}