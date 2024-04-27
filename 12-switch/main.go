package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello to Switch Statements")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)
	switch diceNumber{
	case 1:
		fmt.Println("Dice Value is 1")
	case 2:
		fmt.Println("Dice value is",diceNumber)
	default:
		fmt.Println("Not valid dice number")
	}
}
