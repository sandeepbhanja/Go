package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Banana", "Mango"}
	fmt.Printf("Type of fruitlist is: %T\n", fruitList)

	fruitList = append(fruitList, "Peach", "Tomato")
	fmt.Println("Type of fruitlist is: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("Type of fruitlist is: ", fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 432
	highScore[2] = 45
	highScore[3] = 3232
	highScore = append(highScore, 778)
	fmt.Println(len(highScore))

	sort.Ints(highScore)
	fmt.Println(highScore)

	//how to remove value from slice based on index
	var courses = []string{"Reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
