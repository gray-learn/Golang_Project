package main

import (
	"fmt"
)

func main() {
	colors := []string{"R", "G", "B"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// consise and readable
	for i := range colors {
		fmt.Println(colors[i])
	}

	//index only need value --> "_"

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("value :", value)
		value++
	}

	sum := 1
	for sum < 10 {

		sum += sum
		fmt.Println("sum :", sum)
		if sum > 10 {
			goto theEnd
		}
	}

	// label
theEnd:
	fmt.Println("end of program")

}
