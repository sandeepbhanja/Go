package main

import "fmt"

func main() {
	fmt.Println("Hello, Maps!")
	languages := make(map[string]string,3)
	languages["Go"] = "https://golang.org"
	languages["Python"] = "https://python.org"
	languages["Java"] = "https://java.org"

	fmt.Println(languages["Go"]);

	delete(languages,"Python");

	//looping thorugh maps
	for key, value := range languages {
        fmt.Println(key, value)
    }
}
