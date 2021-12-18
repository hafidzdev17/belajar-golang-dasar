package main

import "fmt"

func main() {

	// TODO: Create Closure
	counter := 0

	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	fmt.Println(counter)

}
