package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string="http://lco.dev"
func main()  {
	fmt.Println("Welcome to web Request")

	response,err :=http.Get(url)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("Response type is %T\n",response)

	defer response.Body.Close() 

	//read response

	datatype,err := io.ReadAll(response.Body)

	content :=string(datatype)

	fmt.Println(content)
	
}