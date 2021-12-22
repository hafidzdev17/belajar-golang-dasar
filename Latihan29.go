package main

import (
	"container/list"
	"fmt"
)

// TODO: Container List

func main() {

	data := list.New()

	data.PushBack("Mohamad")
	data.PushBack("Hafid")
	data.PushBack("Masruri")
	data.PushFront("Nauval")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	//	iterate depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	//	irerate belakang
	for el := data.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Value)
	}
}
