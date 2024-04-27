package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter Age:")
	legalAge := bufio.NewReader(os.Stdin)
	input, _ := legalAge.ReadString('\n')
	age, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	var result string

	if age < 21 {
		result = "Not Legal Age"
	} else {
		result = "Legal Age"
	}
	fmt.Println(result)
}
