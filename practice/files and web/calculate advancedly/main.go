// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	var reslut float64
	v1 := convertInputToValue(input1)
	v2 := convertInputToValue(input2)
	switch operation {
	case "+":
		reslut = addValues(v1, v2)
	case "-":
		reslut = subtractValues(v1, v2)
	case "*":
		reslut = multiplyValues(v1, v2)
	case "/":
		reslut = divideValues(v1, v2)
	default:
		panic("invalid operation")
	}

	// Your code goes here.
	return reslut
}

func convertInputToValue(input string) float64 {
	v, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(fmt.Sprintf("%v must be a number", input))
	}
	return v
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
