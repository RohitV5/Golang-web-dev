package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount < 30 {
		result = "In Between user"
	} else {
		result = "Something else"
	}

	fmt.Println(result)


	if num := 3;num < 10 {
		fmt.Println("Num is less than 10");
	} else {
		fmt.Println("Num is NOT less than 10")
	}


	//Error handling if
	// if err != nil {

	// }
}
