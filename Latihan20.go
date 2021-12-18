package main

import (
	"errors"
	"fmt"
)

// TODO: Error Interface
func handleError(value, bagi int) (int, error) {

	if bagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		result := value / bagi
		return result, nil
	}
}

func main() {

	var exError error = errors.New("oopss error")
	fmt.Println(exError)

	results, err := handleError(100, 10)
	if err == nil {
		fmt.Println("hasil: ", results)
	} else {
		fmt.Println("error: ", err)
	}

}
