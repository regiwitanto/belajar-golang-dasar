package main

import "fmt"

func main() {
	name := "Gilda"

	switch name {
	case "John":
		fmt.Println("Hello John")
	case "Jane":
		fmt.Println("Hello Jane")
	default:
		fmt.Println("Hello Stranger")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Your name is long")
	case false:
		fmt.Println("Your name is short")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Your name is long")
	case length < 5:
		fmt.Println("Your name is short")
	default:
		fmt.Println("Your name is exactly 5 characters")
	}
}