package main

import (
	"flag"
	"fmt"
)

// TODO: Package Flag

func main() {

	var host *string = flag.String("host", "localhost", "put your database host")
	var username *string = flag.String("username", "root", "put your database username")
	var password *string = flag.String("password", "root", "put your database password")

	flag.Parse()

	fmt.Println("host = ", *host)
	fmt.Println("username = ", *username)
	fmt.Println("password = ", *password)
}
