package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "Doe"
	lastName = "Smith"

	return firstName, middleName, lastName
}

func main () {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}