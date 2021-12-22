package main

import "fmt"

// TODO: Pointer Method
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr . " + man.Name
}

func main() {
	hafet := Man{"hafidz"}
	hafet.Married()

	fmt.Println(hafet.Name)
}
