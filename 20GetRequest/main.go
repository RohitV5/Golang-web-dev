package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Perform Get Request")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000"

	response, err := http.Get(myurl)

	

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	//parse using strings package
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body) // content is in byte format
	byteCount,_  := responseString.Write(content)
	fmt.Println("Byte count is ",byteCount)
	fmt.Println(responseString.String())

	//simple way
	// content, _ :=  ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

}
