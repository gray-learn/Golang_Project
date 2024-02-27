package main

// file is member of package called main
import (
	"fmt"
	"bufio" // input
	"os"
	"strconv" // string conversion
	"strings"  // manipulate strings
) // package

const aConst int = 66 // outside public

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ") // prompt
	input, _ := reader.ReadString('\n')
	// underscore ignore variable
	// align feed \n
	fmt.Print("You text: ", input)


	fmt.Print("Enter number: ") 
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("Value of number", aFloat)
	}




// Learning Go

}
