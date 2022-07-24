package main

import "fmt"

func main() {
	fmt.Println("variables")

	var username string = "rohit" 
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false 
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var intVal int = 25576786 
	fmt.Println(intVal)
	fmt.Printf("Variable is of type: %T \n", intVal)

	var smallFloat float32 = 25576786.876866   //retains only 5 decimal point 
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float32 = 25576786.876866   //supports only 5 decimal point 
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)


	// default values
	var anotherVariable int;
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)


	//walrus operator , no var style
	website := "Hi I am a string automatic"
	anotherNUmber := 300
	anotherbool := false
	fmt.Printf("%T %T %T", website, anotherNUmber, anotherbool)


	//public variable starts with capital letter  and private with small.
	Website := "Hi I am public";
	fmt.Printf("%T",Website);


	//declaring constants
	const hello = "Hello, 世界"






}