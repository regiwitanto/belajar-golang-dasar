package main

import "fmt"

func main() {

	type NoKTP string

	var noKTP NoKTP = "1111111111111111"

	var contoh string = "2222222222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(noKTP)
	fmt.Println(contohKtp)
}
