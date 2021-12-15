package main

import "fmt"

func main()  {
	
	// TODO: array with value declare
	listName := [3]string{
		"hafid",
		"deddy",
		"hasyim",
	}
	fmt.Println(listName)

	// TODO: array with value destruct
	listFruit := [...]string{
		"avocade",
		"manggo",
		"banana",
		"grape",
		"water melon",
	}
	fmt.Println(listFruit)

	// TODO: looping array
	for i := 1; i < len(listFruit); i++ {
		fmt.Println(i,".",listFruit[i])
	}

	// TODO: for range
	for _, printFruit := range listFruit {
		fmt.Printf("name fruit: %s\n",printFruit)
	}

}