package main

import "fmt"

func main() {
	const firstName string = "Regi"
	const lastName = "Witanto"

	// firstName = "Witanto" // error
	// lastName = "Regi" // error

	const (
		firstName2 string = "Regi"
		lastName2         = "Witanto"
	)

	fmt.Println(firstName2)
	fmt.Println(lastName2)
}
