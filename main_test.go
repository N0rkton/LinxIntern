package main

import (
	"fmt"
	"log"
)

func ExampleReadCSV() {
	maxPrice, maxRating, err := readCSV("Developer - Backend - Test-2/db.csv")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(maxPrice, maxRating)
	//Output:
	//3 5
}
func ExampleReadJSON() {
	maxPrice, maxRating, err := readJSON("Developer - Backend - Test-2/db.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(maxPrice, maxRating)
	//Output:
	//200 5
}
