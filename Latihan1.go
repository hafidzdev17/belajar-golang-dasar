package main

import "fmt"

// TODO: declare var
var name string = "hafid masruri"
var age int = 23

// TODO: declare multiple var
var (
	firstName = "moh hafid"
	lastName = "masruri"
) 

// TODO: declare constant var
const (
	example1 = "hello"
	example2 = "bro"
)

func main() {
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("Full Name = ", firstName + lastName)
	fmt.Println(example1 + example2)
}





