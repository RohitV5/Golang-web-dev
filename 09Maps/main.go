package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	//gives a key value datasructure
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["TS"] = "Typecript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages) //sorted map
	fmt.Println("JS is short for", languages["JS"])

	delete(languages, "RB")

	fmt.Println(languages)

	// loops are interesting in golang
	// similat to javascript foreach
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
