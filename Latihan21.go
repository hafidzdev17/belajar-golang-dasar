package main

import "fmt"

// TODO: type assertion
func random() interface{} {
	return "200 OK"
}

func main() {

	var random interface{} = random()

	switch value := random.(type) {
	case string:
		fmt.Println("value:", value, "is string")
	case int:
		fmt.Println("value:", value, "is string")
	default:
		fmt.Println("unknown type")
	}

}
