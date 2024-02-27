package main

import (
	"fmt"
)

func main() {
	theAns := 42
	var result string
	// no parentheses
	if theAns < 0 {
		result = "less than 0"
	} else if theAns == 0 {
		result = "equal to 0"
	} else {
		result = "greater thant 0"
	}
	fmt.Println(result)

	if x := -42; x < 0 { // init variable in if statemtn
		result = "less than 0"
	} else if x == 0 {
		result = "equal to 0"
	} else {
		result = "greater thant 0"
	}

	fmt.Println(result)
}
