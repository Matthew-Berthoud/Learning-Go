package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	personNormal := Person{firstName, lastName, age}
	return personNormal
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	personPointer := Person{firstName, lastName, age}
	return &personPointer
}

func main() {
	lucas := MakePerson("Lucas", "Teuber", 21)
	rolf := MakePersonPointer("Rolf", "Hsu", 22)

	fmt.Println(lucas)
	fmt.Println(rolf)
}
