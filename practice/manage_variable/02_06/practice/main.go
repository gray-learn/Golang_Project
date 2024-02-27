package main

// file is member of package called main
import (
	"fmt"
	"math"
) // package

func main() {

	fmt.Print("Math")
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 12.5, 45.1, 68.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum) 
	// may change precision just like java => floating values in Go are stored in binary format
	floatSum = math.Round(floatSum*100) /100 // rounding fraction number(safe)
	fmt.Println("Then sum is now:", floatSum) 
	// Learning Go

	circleRadius := 15.5
	circumference := circleRadius *2 * math.Pi
	fmt.Printf("Circumference %.2f\n" , circumference)

}
