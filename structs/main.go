package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang

	kunal := User{"Kunal", "kunal@godev.com", true, 20}
	fmt.Println(kunal)
	fmt.Printf("Name is %v and email is %v", kunal.Name, kunal.Email)
}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
