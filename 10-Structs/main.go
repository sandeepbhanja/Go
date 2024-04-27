package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// no inheritance in goLang
	// no super or parent
	maverick := User{Email: "Maverick", Name: "maverick@go.dev", Status: true, Age: 22}
	fmt.Printf("Maverick details:%+v\n", maverick)
	fmt.Printf("Name:%v and Email:%v\n", maverick.Name, maverick.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
