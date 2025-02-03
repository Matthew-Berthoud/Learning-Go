// Write a function called fileLen that has an input parameter of type string and returns an int and an error.
// The function takes in a filename and returns the number of bytes in the file.
// If there is an error reading the file, return the error.
// Use defer to make sure the file is closed properly.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	data := make([]byte, 2048)
	total := 0
	for {
		count, err := f.Read(data)
		total += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return total, nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
