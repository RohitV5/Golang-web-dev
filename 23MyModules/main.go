// mod means module tooling
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":4000", r)

	log.Fatal(http.ListenAndServe(":4000",r))

	//go list -m all
	//go mod verify
	//go list -m versions github.com/gorilla/mux
	//go mod tidy
	//go get package-from-web

}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series</h1>"))
}
