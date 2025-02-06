package main

import "fmt"

func UpdateSlice(slice []string, newString string) {
	slice[len(slice) - 1] = newString
	fmt.Println(slice)
}

func GrowSlice(slice []string, newString string) {
	slice = append(slice, newString)
	fmt.Println(slice)
}

func main() {
	slice := []string{"hello", "world", "from", "matthew"}
	fmt.Println(slice)
	UpdateSlice(slice, "justin")
	fmt.Println(slice)
	GrowSlice(slice, "eli")
	fmt.Println(slice)
}

