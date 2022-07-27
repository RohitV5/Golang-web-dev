package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")
	greeter()

	val := adder(1, 2)

	proval, stringVal := proAdder(1, 2, 3, 4)

	fmt.Println(val, "===", proval, stringVal)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Hello from golang")
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total = total + v
	}

	return total, "string also returned"
}

// Anonymous functions and IIF are available in go
// Function cannout be declared inside another function
