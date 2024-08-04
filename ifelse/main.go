package main

import "fmt"

func main() {
	login := 10

	var result string

	if login < 10 {
		result = "Regular user"
	} else if login > 10 {
		result = "Watch out"

	} else {
		result = "Excatly 10"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Number is gretaer than 10")
	}

}
