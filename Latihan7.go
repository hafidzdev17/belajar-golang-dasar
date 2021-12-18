package main

import "fmt"

func main() {

	// TODO: create array
	koridorMembers := [...]string{
		"lutfi",
		"usama oyo",
		"ilham",
		"sholeh",
		"indra",
		"hafid",
	}
	fmt.Println(koridorMembers)

	// TODO: slice index
	sliceKoridor := koridorMembers[5:6]
	fmt.Println(sliceKoridor)
	/* fmt.Println(len(sliceKoridor))
	fmt.Println(cap(sliceKoridor)) */

	// TODO: append slice
	appendKoridor := append(sliceKoridor, "deddy", "bombom", "akbar", "sam", "didik", "karim", "hafidjs")
	fmt.Println(appendKoridor)

	//	TODO: copy slice
	copySlice := make([]string, len(appendKoridor), cap(appendKoridor))
	copy(copySlice, appendKoridor)
	fmt.Println(copySlice)

}
