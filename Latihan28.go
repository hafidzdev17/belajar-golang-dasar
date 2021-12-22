package main

import (
	"fmt"
	"strings"
)

// TODO: Package String

func main() {

	// TODO: String Contains
	str1 := strings.Contains("Hafid Masruri", "Hafid")
	if str1 {
		fmt.Println("Available")
	} else {
		fmt.Println("Not Available")
	}
	// TODO: String Split
	fmt.Println(strings.SplitAfter("hafid masruri", " "))
	// TODO: String Tolower & Upper
	fmt.Println(strings.ToLower("HAFID"))
	fmt.Println(strings.ToUpper("masruri"))
	// TODO: Trim
	fmt.Println(strings.Trim("p hafid masruri p", "p"))
	// TODO: Replace
	fmt.Println(strings.ReplaceAll("hafid", "hafid", "masruri"))
}
