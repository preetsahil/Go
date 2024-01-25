package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"net/url"
)

func main(){
	fmt.Println("this is get request")
    // PerformGetRequest("http://localhost:8000/get")
	// PerformpPostJsonRequest("http://localhost:8000/post")
	PerformpPostFormRequest("http://localhost:8000/postform")

}

func PerformGetRequest(url string){
	response,err:=http.Get(url)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code",response.StatusCode)
	fmt.Println("content length",response.ContentLength)
    
	// var responseString strings.Builder
	content,_:=io.ReadAll(response.Body)
	// byteCount,_:=responseString.Write(content)
	// fmt.Println(byteCount)
	// fmt.Println(responseString.String())

	fmt.Println(string(content))
}

func PerformpPostJsonRequest(myurl string){

	reqBody:=strings.NewReader(`{
		"message":"hello",
		"work":"golang"
	}`)
   response,err:=http.Post(myurl,"application/json",reqBody)

   if err!=nil{
	panic(err)
   }
defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformpPostFormRequest(myurl string){

	data:=url.Values{}
	data.Add("firstname","sahil")
	data.Add("lastname","preet")
	data.Add("city","jalandhar")

	response,err:=http.PostForm(myurl,data)
	   if err!=nil{
	panic(err)
   }
defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))
}