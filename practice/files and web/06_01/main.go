package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from GO"
	// file reference
	file, err := os.Create("./fromString.txt")
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Printf("Wrote a file with %v characters\n",length )
	defer file.Close()// "defer" wait until everything is done 

	defer readFile("./fromString.txt");

}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string){
	data, err := ioutil.ReadFile(fileName)
	checkErr(err)
	fmt.Println("Text read from file", string(data))
}