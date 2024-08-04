package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)

	//	for i := range days {
	//fmt.Println(days[i])
	// }
	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumping at LearncodeOnline.in")
}
