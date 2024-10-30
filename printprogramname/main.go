package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]

	for i := range programName {
		z01.PrintRune(rune(programName[i]))
	}
	z01.PrintRune('\n')
}
