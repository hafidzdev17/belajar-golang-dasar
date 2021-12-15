package main

import "fmt"

func main() {

	for i:= 0; i < 5; i++ {
		for j:= i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}

	fmt.Println("")

	for i:= 5; i > 0; i-- {
		for j:= i; j > 0; j-- {
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}

	// break continue

	// outer loop
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
