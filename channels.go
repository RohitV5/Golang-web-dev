package main

import (
	"fmt"
	"time"
)

func printData(c chan *int) {
	time.Sleep(time.Second * 3)
	data := <-c
	fmt.Println("Data in channel is: ", *data)
}

func main() {
	fmt.Println("Main started...")
	var a = 10
	b := &a
	//create channel
	c := make(chan *int)
	go printData(c)
	fmt.Println("Value of b before putting into channel", *b)
	c <- b
	a = 20
	fmt.Println("Updated value of a:", a)
	fmt.Println("Updated value of b:", *b)
	time.Sleep(time.Second * 2)
	fmt.Println("Main ended...")
}

//Article : https://www.velotio.com/engineering-blog/understanding-golang-channels#:~:text=So%2C%20what%20are%20the%20channels,put%20or%20read%20the%20data.
