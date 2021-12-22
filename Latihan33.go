package main

import (
	"fmt"
	"os"
	"regexp"
)

//TODO: Regexp

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {

	//var regex *regexp.Regexp = regexp.MustCompile("h([a-z])d")
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("hafid"))
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("Fila"))

	/**
	^	Mulai dari text target
	\w	Setiap karakter kata [0-9A-Za-z_]
	+	Setidaknya satu dari karakter sebelumnya
	@	Secara harfiah karakter @
	\.	Karakter dot literal. Harus lolos dengan \
	$	Akhir dari text target
	*/

	emails := []string{
		"brown@fox",
		"brown@fox.",
		"brown@fox.com",
		"br@own@fox.com",
	}

	pattern := `^\w+@\w+\.\w+$`
	for _, email := range emails {
		matched, err := regexp.Match(pattern, []byte(email))
		check(err)
		if matched {
			fmt.Printf("√ '%s' is a valid email\n", email)
		} else {
			fmt.Printf("X '%s' is not a valid email\n", email)
		}
	}
}
