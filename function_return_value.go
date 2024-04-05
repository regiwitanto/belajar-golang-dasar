package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name

	return hello
}

func main() {
	result := getHello("John")
	fmt.Println(result)

	fmt.Println(getHello("Doe"))
	fmt.Println(getHello("Smith"))
}