package main

import (
	"fmt"
	"strconv"
)

// Define a generic interface called Printable that matches a type that
// implements fmt.Stringer and has an underlying type of int or float64.
// Define types that meet this interface.
// Write a function that takes in a Printable and prints its value to the screen using fmt.Println.

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type IntPrint int

func (i IntPrint) String() string {
	return strconv.Itoa(int(i))
}

type FloatPrint float64

func (f FloatPrint) String() string {
	return strconv.FormatFloat(float64(f), 'f', 4, 64)
}

func printPrintable[T Printable](p T) {
	fmt.Println(p.String())
}

func main() {
	a := IntPrint(678)
	b := FloatPrint(2.34098)

	printPrintable(a)
	printPrintable(b)
}
