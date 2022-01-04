package main

import "fmt"

func main() {
	var testVar int = 12

	if testVar%2 == 0 {
		fmt.Println("This number is even")
	} else {
		fmt.Println("This number is not even")
	}

	//switch case
	var testSwitch string = "Hello"

	switch testSwitch {
	case "Nello":
		fmt.Println("Nello")
	case "Hello":
		fmt.Println("Hello")
	default:
		fmt.Println("I dunno")
	}

	//for loops
	for i := 1; i < 10; i++ {
		print("George \n")
	}
}
