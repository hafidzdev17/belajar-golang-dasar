package main

import "fmt"

// TODO: defer, panic & recover

// TODO: create func defer
func endApp() {
	// handle recover
	message := recover()
	if message != nil {
		fmt.Println("message:", message)
	}
	fmt.Println("application done")
}

func runApp(error bool) {
	// Calling Defer
	defer endApp()
	if error {
		// Handle Panic
		panic("APPLICATION ERROR")
	}
	fmt.Println("APPLICATION RUNNING")
}

func main() {
	runApp(true)
	fmt.Println("hafid")
}
