package main

import "fmt"

// TODO: Nil

func checkNil(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			name: name,
		}
	}
}

func main() {

	var newNil map[string]string = checkNil("hafid")

	if newNil == nil {
		fmt.Println("data null")
	} else {
		fmt.Println(newNil)
	}

}
