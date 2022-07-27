package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance of golang; No super or parent

	//first letter cap to make it public. Applies to all eg. The name of struct or property inside struct
	rohit := User{"Rohit", "verma.rohit@gmail.com", true, 16}

	var ak User;  //unless you assign value each propert in struct will have the default value
	fmt.Println("Value before init \n",ak)
	ak.Age = 31
	ak.Email = "ak@gmail.com"
	ak.Name = "AK"
	



	fmt.Println(rohit)
	fmt.Printf("Value after init: %v \n",ak)
	fmt.Printf("Value after init with each type: %R \n",ak)
	fmt.Printf("Type ok ak %T \n",ak)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
