package main

import "fmt"

const LoginToken string = "jdkksosop"
// Capital 'L' signifies it is public

func main() {
	var username string = "kunal"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.5646748393934865
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type

	var website = "happycoding.in"
	fmt.Println(website)
    
    // no var style

	numberOfUsers:=3000
	fmt.Println(numberOfUsers)



}
