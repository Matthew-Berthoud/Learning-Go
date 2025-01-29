package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	// 1. struct literal style, no names
	julia := Employee{"Julia", "Oliveira", 1}

	// 2. struct literal style with names
	beth := Employee{
		firstName: "Beth",
		lastName:  "Smith",
		id:        2,
	}

	// 3. var declaration, dot notation to populate
	var fred Employee
	fred.firstName = "Fred"
	fred.lastName = "Johnson"
	fred.id = 3

	// print all 3
	fmt.Println(julia)
	fmt.Println(beth)
	fmt.Println(fred)
}
