package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Perform Get Request")
	PerformFormDataRequest()
}

func PerformPostRequest() {
	const myurl = "http://localhost:3000/post"

	//fake json payload
	requestBody := strings.NewReader((`
		{
			"coursename":"Golang course",
			"price":0,
			"platform":"web"
		}
	`))

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	//parse using strings package
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body) // content is in byte format
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is ", byteCount)
	fmt.Println(responseString.String())

	//simple way
	// content, _ :=  ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

}


func PerformFormDataRequest() {
	const myurl = "http://localhost:3000/postform"

	//fake formdata payload
	formData := url.Values{}
	formData.Add("firstname","rohit")
	formData.Add("moredata","data")

	response, err := http.PostForm(myurl, formData)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	//parse using strings package
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body) // content is in byte format
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is ", byteCount)
	fmt.Println(responseString.String())

	//simple way
	// content, _ :=  ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

}



