package main

import "fmt"
func main()  {
		fmt.Println("welcome to maps")
		languages:=make(map[int]string)

		languages[0]="Javascript"
		languages[1]="Ruby"
		languages[2]="Go"
		languages[3]="C++"

		fmt.Println(languages)
		fmt.Println(languages[0])

		//delete 
		delete(languages,1)
		fmt.Println(languages)

		//loops are interesting

		for key,value :=range languages{
			fmt.Printf("For Key %v,value is %v\n",key,value)
		}
}