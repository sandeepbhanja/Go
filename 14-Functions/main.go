package main

import "fmt"

func main() {
	fmt.Println("Hello to Functions")
	greeter()
	result := adder(3, 5)
	result1 := proAdder(1, 2, 3, 4, 5)
	fmt.Println(result)
	fmt.Println(result1)
}

func greeter() {
	fmt.Println("Hello Maverick")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
