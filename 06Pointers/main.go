package main

import "fmt"

func main() {
	fmt.Println("welcome to the class on pointers")

	//way 1
	var pointerOne *int;  // *  makes this a pointer.

	fmt.Println("Value of pointer is ",pointerOne)


	//way 2
	myNumber := 23;

	var numPointer = &myNumber;
	fmt.Println("Address of numPointer is ",numPointer)
	fmt.Println("Value of numPointer is ",*numPointer)

	
	*numPointer = *numPointer * 2;
	fmt.Println("Value of numPointer after multiply",*numPointer)

	//making one pointer to point to another pointer
	pointerOne = numPointer;


	*pointerOne = *pointerOne * 2;

	//now both pointers point to same memory location and manipulating
	// the data using one pointer will affect fetching data using another one.
	fmt.Println("Value of numPointer after multiplication of pointerOne",*numPointer)





}

//pointer is a reference to the address of the value in memory.

