package main

import (
	"fmt"
	"os"
)

// TODO: Package OS

func main() {
	args := os.Args
	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname:", name)
	} else {
		fmt.Println("error:", err)
	}
}
