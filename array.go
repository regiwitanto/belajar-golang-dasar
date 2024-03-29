package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		95,
		80,
		70,
		85,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}
