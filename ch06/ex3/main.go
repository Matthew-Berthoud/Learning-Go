package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}


func main() {
	numReps := 10_000_000

	people := []Person{}
	for i := 0; i < numReps ; i++ {
		people = append(people, Person{"Bart", "Simpson", 10})
	}
	fmt.Println("first test done")

	people = make([]Person, 0, numReps)
	for i := 0; i < numReps; i++ {
		people = append(people, Person{"Bart", "Simpson", 10})
	}
	fmt.Println("second test done")

}

// next garbage collection happens when heap size = CURRENT_HEAP_SIZE + CURRENT_HEAP_SIZE*GOGC/100
// if setting GOMEMLIMIT, set GOGC to off

