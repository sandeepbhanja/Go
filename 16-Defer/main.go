package main

import "fmt"

func main() {
	defer fmt.Println("Maverick")
	defer fmt.Println("Maverick2")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v ", i)
	}
}
