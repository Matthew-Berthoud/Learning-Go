package main

import "fmt"

// Write a generic function that doubles the value of any integer or float that’s passed in to it.
// Define any needed generic interfaces.

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func doubleNumber[T Numeric](number T) T {
	return number * 2
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(doubleNumber(i))
	}
	fmt.Println(doubleNumber(14.8))
	// fmt.Println(doubleNumber(true))
	// fmt.Println(doubleNumber("not Numeric"))
}
