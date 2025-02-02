// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x []int
	// mySlice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		n := rand.Intn(100)
		x = append(x, n)
		// mySlice = append(mySlice, rand.Intn(100))
	}
	fmt.Println(x)
	// fmt.Println(mySlice)
}
