package main

import "fmt"

func main() {
	fmt.Println("Welcome to lops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0;d < len(days);d++ {
	// 	fmt.Println(days[d])
	// }

	// Alternative way
	// for d := range days {
	// 	fmt.Println(days[d])
	// }

	//Alternative way
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}


	anyValue := 1;

	for anyValue < 10 {

		if anyValue == 9 {
			goto rv;  //for loop ends after this goto runs.
		}

		if(anyValue == 3){
			anyValue++
			continue
		}

		if(anyValue == 10){
			break
		}
		fmt.Println("Value is ",anyValue)
		anyValue++
	}

	rv: fmt.Println("Jumping to rv.com")
}
