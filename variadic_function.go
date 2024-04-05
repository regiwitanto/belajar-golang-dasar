package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 20, 30, 40, 50)) // 150
	fmt.Println(sumAll(10, 20, 30)) // 60
	fmt.Println(sumAll(10, 20)) // 30

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(sumAll(numbers...)) // 150
}