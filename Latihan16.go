package main

import "fmt"

// TODO: Struct
type Customer struct {
	Name, Address, Gender string
	Age                   int
}

func getCustomer() {

	// Struct Literal
	printCustomer := Customer{
		Name:    "deddy",
		Address: "kumogakure",
		Gender:  "male",
		Age:     100,
	}
	fmt.Println(printCustomer)
}

// TODO: Struct Method
func (customer Customer) sayName(name string) {
	fmt.Println("hai ", name, "namaku adalah ", customer.Name)
}

func main() {
	// TODO: Struct
	getCustomer()

	var data Customer
	data.Name = "hafid"
	data.Address = "konoha"
	data.Age = 23
	data.Gender = "male"

	// TODO: Struct Method
	data.sayName("deddy")
}
