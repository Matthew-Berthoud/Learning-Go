package main

import "fmt"

func main() {
	var b byte = 255 // uint8
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Printf("b = %d, smallI = %d, bigI = %d\n", b, smallI, bigI)
}
