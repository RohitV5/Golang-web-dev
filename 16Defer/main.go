package main

import "fmt"

func main() {
	//Defer functions are invoked
	//in reverse order they are deferred LIFO
	fmt.Println("Hello")
	defer fmt.Println("1st Defer")
	fmt.Println("World")
	defer fmt.Println("2nd Defer")
	myDefer()
}

func myDefer() {
	for i := 0; i <5 ;i++{
		defer fmt.Println(i)
	}
}