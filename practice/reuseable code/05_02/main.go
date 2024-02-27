package main

import (
	"fmt"
)

func main() {

	// Group related values in struct
	// Go do not have inheritance model(java) -
	fmt.Println("Struct is independent")
	poodle := Dog{"Poodle", 20, "woof"} //like construcor define object
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle) // all the data from object
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	// %v placeholdre
	// \n align feed
	poodle.Weight = 12
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "ohhh"
	poodle.Speak()

	poodle.SpeakThreeTime()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// method is member of type
// type assign to method
// Speak is how dogs speaks
func (d Dog) Speak() { // receiver(copy is made instead of reference(pointer))
	// In the function definition func (r Raster) Rastify(dat int) {}, what is the purpose of r Raster? receiver
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTime() {
	d.Sound = fmt.Sprintf("%v %v %v ", d.Sound, d.Sound, d.Sound) // print string
	fmt.Println(d.Sound)
} // reference of Dog object
