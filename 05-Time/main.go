package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time function")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2002, time.January, 18, 15, 25, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}