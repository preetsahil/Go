package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("welcome to slices")

	var fruitlist=[]string{"papaya","kiwi"}
	fmt.Printf("type of fruitlist is %T\n",fruitlist)

	fruitlist=append(fruitlist,"Apple","Banana","Mango")
	fmt.Println("Fruit List is",fruitlist)

	fmt.Println(fruitlist[1:])
	fmt.Println(fruitlist[1:3])
	fmt.Println(fruitlist[:3])
	fruitlist= append(fruitlist[1:3],"nk")
	fmt.Println(fruitlist)

	highScores :=make([]int,4)
	fmt.Printf("type of highscores %T\n",highScores)
	//slices   have array abstraction

	highScores[0]=231
	highScores[1]=797
	highScores[2]=453
	highScores[3]=500
	// highScores[4]=121


	fmt.Println(highScores)
   //reallocation happens again
	highScores=append(highScores,555,1000,234)
	fmt.Println(highScores)
    fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove a value from slice based on index in golang
	// can be used in database operation

	var courses=[]string{"JS","python","swift","GO","C++"}
	fmt.Println(courses)

	index:=2
	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}
