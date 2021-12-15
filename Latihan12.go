package main

import "fmt"

// TODO: variadic function
func sumAll(numbers ...int)int {
	total := 0
	for _,value := range numbers{
		total += value
	}
	return total
}

func main(){

	// TODO: total sum
	total := sumAll(10,10,10,10,10)
	fmt.Println(total)

	// TODO: slice as params
	slice := []int{10,20,30,40,50}
	total = sumAll(slice...)
	fmt.Println(total)
}