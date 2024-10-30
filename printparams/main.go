package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		for _, el := range arg {
			z01.PrintRune(el)
		}
		z01.PrintRune('\n')
	}
}
