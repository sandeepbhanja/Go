package main

import "fmt"

const LoginToken string = "lucifer"

// jwt := 3000 				not allowed

func main() {
	var username string = "Maverick"
	fmt.Println(username)
	fmt.Printf("Variables is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(!isLoggedIn)
	fmt.Printf("Variables is of type : %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variables is of type : %T \n", smallValue)

	var smallFloat float64 = 255.899283091830912830912830
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type : %T \n", smallFloat)

	// default Values and some aliases
	var anotherVariable float32
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of type : %T \n", anotherVariable)

	//implicit type

	var website = "http://localhost:8000"
	fmt.Println(website)

	//no var style
	numberOfUser := 30000.09
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variables is of type : %T \n", LoginToken)
}
