package main

import "fmt"

func main() {
	fmt.Println(("Hello World")) //print to console

	// var helloText string = "Hello World" typed variable declaration
	// nextText := "Hello all"  infered variable decleration

	// george, joy, renee := returnNames() access multpiple return values from function

	// defer  this keyworkd is used to tell the go compiler to execute this function last

	var testString string = "Hey"
	fmt.Println(testString)
	fmt.Printf("type is %T", testString)
}

func printName() string { // function declaration
	return "George Ansong"
}

func returnNames() (string, string, int) { //  return multiple values from a function
	return "George", "Joy", 20
}

func icrementByOne(number int) int { //passing arguments into functions
	return 10
}

func returnFloats(number1, number2 float64) float64 { // shortcut for passing in arguments of the same type
	return 2.90
}
