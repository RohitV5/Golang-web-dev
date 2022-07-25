package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) // standard  formatting syntax

	createdDate := time.Date(2022, time.July, 25, 3, 13, 0, 0, time.UTC)  //year,month,date,hour,min,sec,nanoseconds,timezone
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format(("01-02-2006 Monday")))
}
