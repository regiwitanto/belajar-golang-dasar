package main

import "fmt"

func Ups() any {
	// return 1
	// return "Ups"
	return true
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}