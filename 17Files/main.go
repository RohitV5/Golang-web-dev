package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in golang")

	writeFile("./mygolangfile.txt")
	readFile("./mygolangfile.txt")

}

func writeFile(filename string) {
	content := "This needs to go in a file - RohitVerma"

	file, err := os.Create(filename)

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)
	file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}
