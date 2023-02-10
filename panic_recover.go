package main

import "fmt"

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func division(num1, num2 float32) {

	// execute the handlePanic even after panic occurs
	defer handlePanic()

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		errorString := fmt.Sprintf("Cannot divide %v by %v", num1, num2)
		panic(errorString)
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}

}

func main() {

	division(4, 2)
	division(8, 0)
	division(2, 8)

}
