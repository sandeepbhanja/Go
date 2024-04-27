package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello to Files")
	content := "This needs to go in a file"
	file, err := os.Create("./myFile1.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println(length)
	defer file.Close()
	readFile("./myFile1.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text Data\n", string(databyte))
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}