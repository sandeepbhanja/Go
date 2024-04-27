package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}
	// for i:=0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }
	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		if rougeValue == 8 {
			break
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}
}
