package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


const url = "https://puppygifs.tumblr.com/api/read/json"

func main() {
	fmt.Println("WElcome to http web request");

	response,err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n",response)
	fmt.Println(response) //this wont be of much use


	defer response.Body.Close() //callers responsibility to close the connection

	databytes,err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
	

}

