package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int  
	Platform string `json:"source"`
	Password string `json:"-"`  //will remove from json
	Tags     []string `json:"tags,omitempty"`  //remove any data which doesn't have tag
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	course1 := []course{
		{"ReactJS", 299, "Learn", "password", []string{"web-dev", "js"}},
		{"Angular", 399, "Learn", "password", []string{"web-dev", "js"}},
		{"Vue", 199, "Learn", "password1234", nil},
	}

	//package this data as JSON data

	// Basic json create 
	//finalJson, err := json.Marshal(course1)


	finalJson, err := json.MarshalIndent(course1,"","\t") //formatted json 

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of finalJSon is %T and value is %s\n", finalJson, finalJson)
}
