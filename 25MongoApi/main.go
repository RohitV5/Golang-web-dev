package main

import (
	"fmt"
	"net/http"

	"github.com/rohitv5/Golang-web-dev/25MongoApi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server getting started")
	http.ListenAndServe(":4000",r)
	fmt.Println("Listening at port 4000")
}