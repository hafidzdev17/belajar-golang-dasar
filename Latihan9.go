package main

import (
	"fmt"
	"strings"
)

// TODO: function normal
func printName(message string, arr []string) {
	getName := strings.Join(arr, "")
	fmt.Println(message, getName)
}

// TODO: return value
func printMessage(message string) string {
	return "hello " + message
}

// TODO: function return multiple value
func printFullName() (string, string, string) {
	return "Mohamad", "Hafid", "Masruri"
}

// TODO: function return array
func printArray(arr [5]string) [5]string {
	return arr
}

// TODO: function return map
func printMap(member map[string]string) map[string]string {
	return member
}

// TODO: function return named
func printNamed() (firstName, middleName, lastName string) {
	firstName = "Fatih"
	middleName = "Nauval"
	lastName = "Azzidan"
	return
}

// TODO: function value
func test(name string) string {
	return name
}

func main() {

	showName := []string{"hafid", "masruri"}
	printName("hello", showName)

	showMessage := printMessage("golang")
	fmt.Println(showMessage)

	_, middleName, _ := printFullName()
	fmt.Println(middleName)

	showArray := printArray([5]string{"hafid", "nauval", "firli", "fila"})
	fmt.Println(showArray)

	showMap := printMap(map[string]string{
		"name":    "Hafid",
		"age":     "23",
		"address": "Konoha",
	})
	fmt.Println(showMap)

	_, m, l := printNamed()
	fmt.Println(m, l)

	showValue := test
	fmt.Println(showValue("hello"))
}
