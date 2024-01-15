package main

import (
	"fmt"
	"net/url"
)

const myurl string="https://lco.dev:3000/learn?coursename=react.js&paymentId=fmas2332"

func main()  {
	fmt.Println("understanding urls")
	fmt.Println(myurl)

	//parsing --to understand each word

	result,_:= url.Parse(myurl)
	// fmt.Println(result)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Port())
	// fmt.Println(result.Host)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Path)

	// fmt.Printf("type is %T\n",result.RawQuery)

	qparams := result.Query();  //for better format

	fmt.Printf("type of qparams is %T\n",qparams)
	fmt.Println(qparams["coursename"])

	for i,value :=range qparams{
		fmt.Printf("key is %v value is %v\n",i,value)
	}

	//to create a url out of info
	//you have to pass reference of the url not the copy

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)






	
}