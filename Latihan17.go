package main

import "fmt"

// TODO: Interface
type HasName interface {
	GetName() string
}

// TODO: Create Contract Interface
func sayHello(hasName HasName) {
	fmt.Println("hello ", hasName.GetName())
}

// TODO: Implementasi Interface
type Person struct {
	name string
}

type Animal struct {
	name string
}

func (person Person) GetName() string {
	return person.name
}

func (animal Animal) GetName() string {
	return animal.name
}

func main() {

	var person Person
	person.name = "hafid"
	sayHello(person)

	var animal Animal
	animal.name = "dragon"
	sayHello(animal)

}
