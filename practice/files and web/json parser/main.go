package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"io"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	// fmt.Print(content)

	tours := toursFromJson(content)
	for _, tour := range tours { // _ dont care index
		fmt.Println(tour.Name)
	}
}

func toursFromJson(content string) []Tour { // [] slice
	tours := make([]Tour, 0, 20) // create slice

	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() { // read next availabel objcet
		err := decoder.Decode(&tour) // memory address &tour
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
