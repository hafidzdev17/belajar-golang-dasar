package main

import (
	"fmt"
	"go-learn/database"
	//_ "go-learn/database"
)

// TODO: package initialization,blank identifier

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
