package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "This is text that will be stored in a file"
	file, err := os.Create("./textFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("File Lenght is ", length)
	readFile("./textFile.txt")
	defer file.Close()
}

func readFile(filePath string) {
	databyte, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}
