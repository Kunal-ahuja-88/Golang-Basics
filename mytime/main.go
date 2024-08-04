package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handling time in Golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.June, 01, 06, 15, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
