package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married()  {
	man.Name = "Mr. " + man.Name
}

func main() {
	regi := Man{"Regi"}
	regi.Married()

	fmt.Println(regi)
}