package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(7) + 1
	// fmt.Println("Day", dow)
	// ceiling of 7

	var result string
	// switch dow { // no break
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Sun"
	case 2:
		result = "Mon"
	default:
		result = "other day"

	}
	fmt.Println(result)

	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Sun"
		fallthrough
	case 2:
		result = "Mon"
		fallthrough // go to bottom excute to next case once true
	default:
		result = "other day"

	}
	fmt.Println(result)

}
