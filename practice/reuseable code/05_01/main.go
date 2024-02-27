package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum := addValues(2, 5)
	fmt.Println(sum)
	total, valuecnt := addValuess(2, 1, 56, 3) // arbitray
	fmt.Println(total)
	fmt.Println("valuecnt", valuecnt)

}

func doSomething() {
	fmt.Println("do Something")
}

func addValues(i1, i2 int) int {
	return i1 + i2

}

func addValuess(values ...int) (int, int) {
	totoal := 0
	for _, v := range values {
		totoal += v
	}
	return totoal, len(values)
}
