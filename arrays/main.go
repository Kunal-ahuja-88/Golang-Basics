package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("Fruits present are :", fruitList)
	fmt.Println("Fruits present are :", len(fruitList))

	var vegList = [3]string{"potato", "tomato", "beans"}
	fmt.Println("Veg list is", vegList)
}
