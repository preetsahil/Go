package main

import "fmt"

func main()  {

	fmt.Println("welcome to loops")

	days:=[]string{"mon","tue","wed","thru","fri"}

	// for d:=0; d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	
	// for each 

	// for index,day :=range days{
	// 	fmt.Printf("index is %v and value is %v\n",index,day)
	// }
	for _,day :=range days{
		fmt.Printf("index is v and value is %v\n",day)
	}

	 rougueValue := 1

	 for rougueValue < 10 {

		// if rougueValue == 3{
		// 	break
		// }
		// if rougueValue == 3{
		// 	rougueValue++
		// 	continue
		// }
		 if rougueValue == 2{
			goto sahil
		 }


		fmt.Println("Value is: ",rougueValue)
		rougueValue++
	 }

    sahil:fmt.Println("using goto")


}