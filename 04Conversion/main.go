package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")
	
	//in go we dont have try catch, so we have comma,ok syntax
	input, _ := reader.ReadString('\n')  //reading till newline is encountered
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T \n", input)

	removeslashN := strings.TrimSpace(input)

	numRating,error  := strconv.ParseFloat(removeslashN,64);
	fmt.Printf("Type after conversion of rating is %T \n", numRating)

	if error != nil {
		fmt.Println(error)
	}else{
		fmt.Println("Added 1 to your rating", numRating + 1)
	}




}