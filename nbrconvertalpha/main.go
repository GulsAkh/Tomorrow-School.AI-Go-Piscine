package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	isUpper := false
	startIndex := 1

	// check for the --upper flag
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		isUpper = true
		startIndex = 2
	}
	for j := startIndex; j < len(os.Args); j++ {
		arg := os.Args[j]
		num := 0
		valid := true

		for i := 0; i < len(arg); i++ {
			digit := arg[i] - '0'

			if digit < 0 || digit > 9 {
				valid = false
				break
			}
			num = num*10 + int(digit)
		}
		if !valid || num < 1 || num > 26 {
			z01.PrintRune(' ')
			continue
		}
		var letter rune
		if isUpper {
			letter = rune(num + 'A' - 1)
		} else {
			letter = rune(num + 'a' - 1)
		}
		z01.PrintRune(letter)

		// z01.PrintRune('\n')
	}
}
