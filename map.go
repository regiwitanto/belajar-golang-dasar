package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "John Doe"
	// person["age"] = "25"

	person := map[string]string{
		"name": "John Doe",
		"age":  "25",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "John Doe"
	book["ups"] = "Salah"
	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
