package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var regi Customer
	fmt.Println(regi)

	regi.Name = "Regi"
	regi.Address = "Jakarta"
	regi.Age = 25
	fmt.Println(regi)
	fmt.Println(regi.Name)
	fmt.Println(regi.Address)
	fmt.Println(regi.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Surabaya", 35}
	fmt.Println(budi)

	budi.sayHello("Agus")
	regi.sayHello("Agus")
	joko.sayHello("Agus")
}