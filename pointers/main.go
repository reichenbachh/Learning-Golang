package main

import "fmt"

func main() {
	//In golang when an argument is passed into a function, the function creates a copy of
	// the variable and uses it, it doesnt alter the original variable
	var myFavFruit string = "Pear"
	fmt.Println((myFavFruit))
	changeFavFruit(&myFavFruit, "Orange")
}

func changeFruitName(favFruit, newFruit string) { // this wouldnt work because of the above reason
	// to actually change the variabe the pointer should be passed in
	favFruit = newFruit
}

func changeFavFruit(favFruit *string, newFruit string) {
	fmt.Println("The Adress for favfruit is:", favFruit)
	*favFruit = newFruit // this will actually alter the data stored in the variable
	// favFruit is the pointer
	// *favFruit is the value
}
