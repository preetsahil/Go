package main

import "fmt"

type User struct {
	Name string
	Email string
	Status bool
	Age     int
}

func main()  {
		fmt.Println("welcome to structs")

		sahil:=User{"sahil","preet@gmail.com",true,16}
		fmt.Println(sahil)
		fmt.Printf("Sahil details is %+v\n",sahil)
		fmt.Printf("Name is %v and email is %v\n",sahil.Name,sahil.Email)
}