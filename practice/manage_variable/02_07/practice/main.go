package main

// file is member of package called main
import (
	"fmt"
	"time"
) // package

func main() {

	n := time.Now()
	fmt.Println("this time video at", n)
	t := time.Date(2012, time.January, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("this time launch at", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Jan 10 23:00:00 2012")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}
