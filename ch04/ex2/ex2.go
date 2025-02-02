// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
// Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
// If the value is divisible by 2, print “Two!”
// If the value is divisible by 3, print “Three!”
// If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// Otherwise, print “Never mind”.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randVals []int
	for i := 0; i < 100; i++ {
		n := rand.Intn(100)
		randVals = append(randVals, n)
	}

	for _, v := range randVals {
		switch {
			case v % 6 == 0:
				fmt.Println("Six!")
			case v % 2 == 0:
				fmt.Println("Two!")
			case v % 3 == 0:
				fmt.Println("Three!")
			default:
				fmt.Println("Never mind")
		}
	}
}

