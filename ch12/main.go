package main

import (
	"fmt"
	"math"
	"sync"
)

/*
1. Create a function that launches three goroutines that communicate using a channel.
The first two goroutines each write 10 numbers to the channel.
The third goroutine reads all the numbers from the channel and prints them out.
The function should exit when all values have been printed out.
Make sure that none of the goroutines leak.
You can create additional goroutines, if needed.
*/

func ex1() {
	fmt.Println("Exercise 1")
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			ch <- j * 2 // write even numbers
		}

	}()
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			ch <- j*2 + 1 // write odd numbers
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}

/*
2. Create a function that launches two goroutines.
Each goroutine writes 10 numbers to its own channel.
Use a for-select loop to read from both channels, printing out the number and the goroutine that wrote the value.
Make sure that your function exits after all values are read and that none of your goroutines leak.
*/
func ex2() {
	fmt.Println("Exercise 2")
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		for j := 0; j < 10; j++ {
			ch <- j * 2 // write even numbers
		}
		close(ch)
	}()
	go func() {
		for j := 0; j < 10; j++ {
			ch2 <- j*2 + 1 // write odd numbers
		}
		close(ch2)
	}()
	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
				count--
				break
			}
			fmt.Printf("%d, from channel 1\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			fmt.Printf("%d, from channel 2\n", v)
		}
	}
}

/*
3. Write a function that builds a map[int]float64 where the keys are
the numbers from 0 (inclusive) to 100,000 (exclusive) and the values are the square roots of those numbers
(use the math.Sqrt function to calculate square roots).
Use sync.OnceValue to generate a function that caches the map returned by this function
and use the cached value to look up square roots for every 1,000th number from 0 to 100,000.
*/

func getRoots() map[int]float64 {
	fmt.Println("Exercise 3")
	roots := make(map[int]float64)
	for i := 0; i < 100_000; i++ {
		roots[i] = math.Sqrt(float64(i))
	}
	return roots 
}
func ex3() {
	var getRootsCached = sync.OnceValue(getRoots)
	roots := getRootsCached()
	for i := 0; i < 100_000; i += 1_000 {
		fmt.Printf("%d: %f\n", i, roots[i])
	}
}

func main() {
	ex1()
	ex2()
	ex3()
}

