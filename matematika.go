package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var d = 30
	var e = 40
	var c = a + b - d * e
	fmt.Println(c)


	var i = 10
	i += 10
	fmt.Println(i)


	i += 5
	fmt.Println(i)

	var j = 10
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}
