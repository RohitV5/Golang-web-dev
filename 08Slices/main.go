package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")

	// Syntax 1
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Value of fruitList is", fruitList)
	fmt.Println("Length of fruitList is", len(fruitList))

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("Value of fruitList is", fruitList)
	fmt.Println("Length of fruitList is", len(fruitList))

	// get element from index (including from) till index (excluding till)
	fruitList = append(fruitList[1:3])

	fmt.Println("Value of fruitList after deleting first time", fruitList)
	fmt.Println("Length of fruitList after deleting first time is", len(fruitList))

	fruitList = append(fruitList[1:])

	fmt.Println("Value of fruitList after deleting 2nd time", fruitList)
	fmt.Println("Length of fruitList after deleting 2nd time is", len(fruitList))

	fruitList = append(fruitList[:0])

	fmt.Println("Value of fruitList after deleting 3rd time", fruitList)
	fmt.Println("Length of fruitList after deleting 3rd time is", len(fruitList))

	//Syntax 2
	highScores := make([]int, 4) //default length 4 but can be increased using append

	highScores[0] = 324
	highScores[1] = 225
	highScores[2] = 426
	highScores[3] = 227
	highScores = append(highScores, 928, 129)

	fmt.Println("Value of highScores is", highScores)
	fmt.Println("Length of highScores is", len(highScores))

	sort.Ints(highScores)

	fmt.Println("Value of highScores after sorting is", highScores)

	//IMPORTANT: Removing value from slice based on index
	var courses = []string{"reactjs", "angular", "vue", "svelte", "ruby","nodejs"}

	var index int = 2;
	courses = append(courses[:index],courses[index+1:]...)

	fmt.Println("Value after deleting at index", courses)


	//NOTE: square bracket [] can be used to give index and get value at that index or can be used to pass two arguments to slice a value

	//Append replace existing value and keeps everything that is inside append function




}
