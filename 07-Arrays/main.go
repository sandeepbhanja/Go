package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "mango"

	fmt.Println("Fruit List is :", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Vegies List is :", vegList)
}
