package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int
	Platform string  `json:"website"`
	Password string   `json:"-"`  //remove entire field from response
	Tags   []string    `json:"tags,omitempty"`
}

func main()  {
	fmt.Println("welcome on how to create json and handle json")
	// encodejson()
	DecodeJson()
}

func encodejson(){
	lcoCourses :=[]course{
		{"ReactJS Bootcamp",299,"learn.dev","abc123",[]string{"web-dev","js"}},
		{"MERN Bootcamp",199,"learn.dev","ab23",[]string{"nodejs","reactjs"}},
		{"Go Bootcamp",299,"learn.dev","abc13",nil},
	}

	// finalJson,err :=json.Marshal(lcoCourses)
	//to convert to json data
	finalJson,err :=json.MarshalIndent(lcoCourses,"","\t")
	// finalJson,err :=json.MarshalIndent(lcoCourses,"1","\t")


	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}

//how to decode json data

func DecodeJson(){
	jsonDataFromWeb :=[]byte(`
	   {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "learn.dev",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var lcoCourse course 
	checkValid :=json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)
		//to print interfaces "%#v"
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("JSON was not valid")
	}
    
	// some cases where you just want to add data to key value
	//not create struct create map 

	var myOnlineData map[string]interface{}  
	// since the data is coming from online web we don't know what type of data will come
	//  so we make it a interface and populate it 
	// key can string and value can be anything coming from web
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for key,value:=range myOnlineData{
		fmt.Printf("key is %v and value is %v and Type is %T\n",key,value,value)
	}
	 
	}