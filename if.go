package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "John" {
		fmt.Println("Hello John")
	} else if name == "Jane" {
		fmt.Println("Hello Jane")
	} else {
		fmt.Println("Hello Stranger")
	}

	if length := len(name); length > 5 {
		fmt.Println("Your name is long")
	} else {
		fmt.Println("Your name is short")
	}
}
