package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open or run")
	case 2:
		fmt.Println("Dice value is 1 and you can run")
	case 3:
		fmt.Println("Dice value is 3 and you can run")
	case 4:
		fmt.Println("Dice value is 4 and you can run")
	case 5:
		fmt.Println("Dice value is 5 and you can run")
		fallthrough  // if this case got hit then the next one will be triggered automatically
	case 6:
		fmt.Println("Dice value is 6 and you can open or run and roll dice again")
		fallthrough
	default:
		fmt.Println("Bug encountered. Report to dev. Must be a fallthrough")
		

	}

}

// fallthrough
