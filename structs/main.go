package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// fmt.Println(alex)
	var alex person

	alex.firstName = "Rex"
	alex.lastName = "Parker"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
