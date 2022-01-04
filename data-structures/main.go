package main

import "fmt"

type Employee struct {
	firtsName, lastName string
}

//interfaces are contracts that are mostly used to design the behaviour of structs
type Shape interface {
	Area() float64
	Premimeter() float64
}

type Triangle struct {
	base   float64
	heigth float64
}

//triangle recievers
func (triangle Triangle) returnTriangleArea() float64 {
	return 0.5 * triangle.base * triangle.heigth
}

func (triangle Triangle) returnTrianglePerimeter() float64 {
	return triangle.base * triangle.heigth
}

var myTriangle Shape

func main() {
	//Arrays, There are two array types in go, normal arrays and slices
	// myArray := [5]int{1, 2, 3, 4, 5} defauly Array declaration, type and size need to be specified

	// var mySlice []int array with a dynamic size

	// 2.2 Or like this
	// d := []int{10, 20, 30, 40, 50}
	// fmt.Println("d:", d)

	// 2.3 Or like this
	// newArray := make([]int, 4) //make function is used to create slice and maps, when a slice is created it is filled with 0 or other data, so when append is used it adds to the already existing data
	// newArray = append(newArray, 1, 2, 3, 4)
	// newArray[2] = 100          //  array elements can be accessed using bracket notation
	// fmt.Println(newArray[0:4]) //sub slices or an array can also be accessed like this
	// fmt.Println("s:", newArray)

	// //looping through an array
	// for k, v := range newArray {
	// 	fmt.Println(k, v)
	// }

	//Maps

	// var myMap map[string]string  map declaration

	// personalData := map[string]string{
	// 	"firstName": "George",
	// 	"lastName":  "Ansong",
	// 	"residence": "Accra",
	// }

	// //adding to a map
	// personalData["middleName"] = "Wiredu"

	// //deleting from a map
	// delete(personalData, "middleName")

	// //replacing values in the map
	// personalData["firstName"] = "Kofi"

	// //looping through a map
	// for k, v := range personalData {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println(personalData)

	//Structs
	//Struct are data types in go langs that behanve like classes
	//structs can also be passed into structs

	// type person struct {
	// 	name      string `json:firstName yaml:firstName` // this is known as tagging, this can be used when serializing
	// 	residence string
	// 	age       int
	// }

	// person1 := person{
	// 	name:      "George Anson",
	// 	residence: "Spintex",
	// 	age:       22,
	// }

	// type animal struct {
	// 	name     string
	// 	features []string
	// }

	// newAnimal := animal{
	// 	name: "Zebra",
	// 	features: []string{
	// 		"Head", "Body", "Tail",
	// 	},
	// }

	// //promotion, passing a struct into another struct
	// type bigStruct struct {
	// 	animal   animal
	// 	eatHuman bool
	// }

	// nextAnimal := bigStruct{
	// 	animal: animal{
	// 		name: "Lion",
	// 		features: []string{
	// 			"Head", "Waist", "Tail",
	// 		},
	// 	},
	// 	eatHuman: true,
	// }

	// fmt.Println(newAnimal.name) // data in structs can be accessed using dot notation
	// fmt.Println(nextAnimal)
	// fmt.Println(person1, newAnimal)

	//Ananymous structs
	//can be used to create structs without specifying struct types

	// myProfile := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "George",
	// 	age:  40,
	// }

	// fmt.Println(myProfile)

	//Recievers
	// These are functions that can be attached to structs like class methods

	myEmployee := Employee{
		firtsName: "George",
		lastName:  "Ansong",
	}

	//We can use recievers to attach functions to the struct

	fmt.Println(myEmployee)

	c := Triangle{heigth: 0.40, base: 0.99}

	c.returnTrianglePerimeter()

}

func (employee Employee) printEmployee() { // this is reciever by value , we can read the value but we cant change it
	fmt.Println(employee.firtsName)
}

func (employee *Employee) changeEmployee() { // this is function reciever by refrence this we pass a pointer to the struct
	employee.firtsName = "Jackson"
}
