package main

import (
	"fmt"
	"strconv"
)

func main() {
	//uses the strings package

	var myString string = "10"

	res, err := strconv.ParseInt(myString, 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
