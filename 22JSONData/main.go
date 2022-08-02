package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"source"`
	Password string   `json:"-"`              //will remove from json
	Tags     []string `json:"tags,omitempty"` //remove any data which doesn't have tag
}

func main() {
	// EncodeJson()
	DecodeJson()
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

	finalJson, err := json.MarshalIndent(course1, "", "\t") //formatted json

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of finalJSon is %T and value is %s\n", finalJson, finalJson)
}

func DecodeJson() {
	JsonDataFromWeb := []byte(`[{
		"coursename": "ReactJS",
		"Price": 299,
		"Platform": "Learn",
		"Password": "password",
		"Tags": [
				"web-dev",
				"js"
				]
		},
		{
				"coursename": "Angular",
				"Price": 399,
				"Platform": "Learn",
				"Password": "password",
				"Tags": [
						"web-dev",
						"js"
				]
		},
		{
				"coursename": "Vue",
				"Price": 199,
				"Platform": "Learn",
				"Password": "password1234",
				"Tags": null
		}
		]`)

	var courseone []course

	checkValid := json.Valid(JsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(JsonDataFromWeb, &courseone)
		fmt.Printf("%#v\n", courseone)
	}

	//some cases where you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(JsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T", k, v, v)
	}
}
