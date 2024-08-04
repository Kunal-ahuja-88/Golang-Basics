package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Go"] = "Golang"

	fmt.Println("Languages are : ", languages)

	// loops

	for key, value := range languages {
		fmt.Printf("For key %v,value is %v\n", key, value)
	}
}
