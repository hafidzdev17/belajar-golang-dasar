package main

import (
	"fmt"
)

// TODO: Anonymous Function
type Blacklist func(name string)bool

func registerUser(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println("welcome " , name)
	} else {
		fmt.Println("Blocked ", name)
	}
}

func main() {

	user := func(name string)bool {
		return name == "hafid"
	}

	registerUser("root",user)

}