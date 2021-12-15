package main

import "fmt"

func main()  {

	// TODO: block if
	var number = 90

	if number > 80 {
		fmt.Println("A")
	} else {
		fmt.Printf("E")
	}

	// TODO: if short statement
	var name = "uchiha madara"
	if length := len(name); length > 5 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// TODO: variabel temporary
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// TODO: block switch
	var value = 6

	switch value {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// TODO: style if-else
	var points = 90

	switch {
	case points == 80 :
		fmt.Println("perfect")
	case points > 80 && points < 100:
		fmt.Println("Super Perfect")
		fallthrough
	case (points > 80 && points == 90):
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// TODO: nested-if-switch
	var val = 9

	if val > 7 {
		switch val {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if val == 5 {
			fmt.Println("not bad")
		} else if val == 3 {
			fmt.Println("keep trying ok")
		} else {
			fmt.Println("you can do it")
			if val == 0 {
				fmt.Println("try harder")
			}
		}
	}

}