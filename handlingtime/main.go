package main

import (
	"fmt"
	"time"
)

//Package time provides functionality for measuring and displaying time.
func main()  {

	fmt.Println("understanding time handling in go")

	presentTime := time.Now()
	fmt.Println("present time is ", presentTime)
	fmt.Println("present time in string format is ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	//you can create date time object using date time function
	//func time.Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time
	createdDate :=time.Date(2018,time.December,4,6,30,0,0,time.UTC)

	fmt.Println("created date is, ",createdDate)
	fmt.Println(createdDate.Format(("01-02-2006 Monday")))

	//go env command will give you the environment variables and bunch of other things  and can we used in the command 
	// GOOS="windows" go build
	
	
}