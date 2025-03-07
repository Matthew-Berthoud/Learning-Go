package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var english string

//go:embed spanish_rights.txt
var spanish string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./BINARY LANGUAGE")
		return
	}
	switch os.Args[1] {
		case "english":
			fmt.Print(english)
		case "spanish":
			fmt.Print(spanish)
	}
}
