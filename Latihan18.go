package main

import "fmt"

//	TODO: Interface Kosong

func Ups(i int) interface{} {
	if i != 0 {
		return true
	} else {
		return false
	}
}

func main() {
	ups := Ups(1)
	fmt.Println(ups)
}
