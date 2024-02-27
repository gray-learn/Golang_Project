package main

import (
	"fmt"
)

func main() {
	// Group related values in struct
	// Go do not have inheritance model(java) -
	fmt.Println("Struct is independent")
	poodle := Dog{"Poodle", 20} //like construcor define object
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle) // all the data from object
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	// %v placeholdre
	// \n align feed
	poodle.Weight = 12
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}
