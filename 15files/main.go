package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - xyz.edu"

	file, err := os.Create("./myfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkErrNil(err)

	len, err := io.WriteString(file, content)

	checkErrNil(err)

	fmt.Println("length is :", len)
	defer file.Close()
	readFile("./mylocfile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkErrNil(err)
	fmt.Println("Text data inside the file is:", string(dataByte))
}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
