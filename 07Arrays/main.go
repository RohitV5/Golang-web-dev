// Array is very less compared to Slice because of its fixed length restrictions.
// But Arrays are more efficient compared to slice.

package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string;

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList)
	//console puts a blank space if any index is empty

	fmt.Println("Fruit list length is", len(fruitList))
	// Will give the defined value of array irrespective of what is stored in array.

	var vegList = [5]string{"potato", "beans", "squash"}

	fmt.Println("Vegetable list is: ", vegList)
	//console puts a blank space if any index is empty

	fmt.Println("Vegetable list length is", len(vegList))
	// Will give the defined value of array irrespective of what is stored in array.
	
}