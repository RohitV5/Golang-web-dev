package main

import "fmt"

func main() {
	fmt.Println("variables")

	var username string = "hitesh" 
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false 
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
}