package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	n := len(arg)
	for j := 0; j < n-1; j++ {
		for k := 0; k < n-j-1; k++ {
			if arg[k] > arg[k+1] {
				arg[k], arg[k+1] = arg[k+1], arg[k]
			}
		}
	}
	for _, sort := range arg {
		for _, char := range sort {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
