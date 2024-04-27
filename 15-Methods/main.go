package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// no inheritance in goLang
	// no super or parent
	maverick := User{"Maverick", "maverick@go.dev", true, 22}
	// fmt.Printf("Maverick details:%+v\n", maverick)
	// fmt.Printf("Name:%v and Email:%v\n", maverick.Name, maverick.Email)
	result := maverick.GetStatus()
	fmt.Println(result)
	maverick.NewEmail()
	fmt.Println(maverick)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() int {
	fmt.Println(u)
	return u.Age
}

func (u User) NewEmail() {
	u.Email = "test@google.com"
	// fmt.Println(u)
}
