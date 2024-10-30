package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	if len(name) >= 2 && name[0] == '.' && name[1] == '/' {
		name = name[2:]
	}
	for _, el := range name {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
}
