package main

import "fmt"

func main() {

	// TODO: create map
	foodPrice := map[string]int{
		"chicken":   150,
		"noodle":    50,
		"hamburger": 300,
		"hotdog":    100,
	}

	// TODO: map terpisah
	courseList := make(map[string]string)
	courseList["name"] = "Golang Fundamental"
	courseList["category"] = "Back End"
	courseList["mentor"] = "Hafidz Masruri"
	fmt.Println(courseList)

	// TODO: map iteration
	for key, value := range foodPrice {
		fmt.Println(key, " \t:", value)
	}

	// TODO: delete map
	delete(foodPrice, "noodle")
	fmt.Println(foodPrice)

	// TODO: multiple map
	koridorMembers := []map[string]string{
		map[string]string{
			"name": "hafid",
			"nim":  "17010197",
		},
		map[string]string{
			"name": "deddy",
			"nim":  "17010199",
		},
	}

	// TODO: multiple map iteration
	for _, getKoridor := range koridorMembers {
		fmt.Println(
			"name:", getKoridor["name"],
			"nim:", getKoridor["nim"],
		)
	}

}
