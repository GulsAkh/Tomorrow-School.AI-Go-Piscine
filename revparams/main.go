package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i >= 1; i-- {
		arg := os.Args[i]
		for _, el := range arg {
			z01.PrintRune(el)
		}
		z01.PrintRune('\n')
	}
}
