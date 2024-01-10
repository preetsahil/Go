package main

import "fmt"

func main()  {
	fmt.Println("welcome to function,methods,defer")

	// greeter();
	
	// result :=adder(3,5)

	// fmt.Println("sum is",result)

	// proRes := proAdder(2,3,4,4,5)

	// two return values
	// proRes,_ := proAdder(2,3,4,4,5)
	// proRes,nameIs := proAdder(2,3,4,4,5)
    // fmt.Println("prosum is",proRes)

	// sahil := User{"sahil","preetsahil@gmail.com",true}

	// sahil.GetStatus()
	// sahil.newMail()


	// fmt.Println(sahil.Email)

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
   defer fmt.Println("world")
   defer fmt.Println("one")
   defer fmt.Println("two")

   fmt.Println("Hello")
    myDefer()
//    print  hello two one world

}
// hello 43210 two one world
func myDefer(){
	for i:=0;i<5;i++ {
		defer fmt.Println(i)
	}
}

func greeter(){
	fmt.Println("welcome")
}
func adder(a int,b int) int {
	return a+b;
}
//values is slsice
func proAdder(values ...int) (int, string) {
	total :=0

	for _,value := range values{
		total +=value
	}
	return total,"hi my name is"
}

type User struct {
	Name string
	Email string
	Status bool
}

func (u User) GetStatus(){
	fmt.Println("is user Active",u.Status)
}

func (u *User) newMail(){
	u.Email="test@dev.go"
	fmt.Println("email  is",u.Email)
}