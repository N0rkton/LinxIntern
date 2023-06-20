package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type product struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
	Rating  int    `json:"rating"`
}

func main() {
	jsonFlag := flag.String("j", "Developer - Backend - Test-2/db.json", "path to json file")
	csvFlag := flag.String("c", "Developer - Backend - Test-2/db.csv", "path to csv file")
	flag.Parse()
	maxRating, maxPrice, err := readJSON(*jsonFlag)
	if err == nil {
		fmt.Printf("max price from json file: %d, max raiting: %d\n", maxPrice, maxRating)
	}
	maxRating, maxPrice, err = readCSV(*csvFlag)
	if err == nil {
		fmt.Printf("max price from csv file: %d, max raiting: %d", maxPrice, maxRating)
	}

}
func readJSON(path string) (int, int, error) {
	f, errF := os.Open(path)
	defer f.Close()
	if errF != nil {
		return 0, 0, errF
	}
	s := bufio.NewScanner(f)
	var tmp product
	var maxPrice, maxRating int
	for s.Scan() {
		if err := json.Unmarshal(s.Bytes(), &tmp); err != nil {
			continue
		}
		if tmp.Rating > maxRating {
			maxRating = tmp.Rating
		}
		if tmp.Price > maxPrice {
			maxPrice = tmp.Price
		}
	}
	return maxPrice, maxRating, nil
}
func readCSV(path string) (int, int, error) {
	f, errF := os.Open(path)
	if errF != nil {
		return 0, 0, errF
	}
	defer f.Close()
	var maxPrice, maxRating int
	csvReader := csv.NewReader(f)
	csvReader.Read()
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			return maxPrice, maxRating, nil
		}
		if err != nil {
			return 0, 0, err
		}

		tmpRating, ok := strconv.Atoi(rec[2])
		if ok != nil {
			return 0, 0, ok
		}
		tmpPrice, ok := strconv.Atoi(rec[1])
		if ok != nil {
			return 0, 0, ok
		}
		if tmpRating > maxRating {
			maxRating = tmpRating
		}
		if tmpPrice > maxPrice {
			maxPrice = tmpPrice
		}
	}
}
