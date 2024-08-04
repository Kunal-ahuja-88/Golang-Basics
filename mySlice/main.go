package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "peach", "mango"}

	fruitList = append(fruitList, "Chiku", "dryfruit")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 208
	highScores[1] = 209
	highScores[2] = 264
	highScores[3] = 200

	//reallocation of memory

	highScores = append(highScores, 548, 467, 221)

	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println((sort.IntsAreSorted(highScores)))

	//removing values through slice

	var courses = []string{"reactjs", "javascript", "python", "ruby", "rust"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
