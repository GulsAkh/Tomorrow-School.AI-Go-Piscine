package main

import (
	"github.com/01-edu/z01"
)

func NRune(s string, n int) rune {
	var runes rune
	if n >= 0 && n <= len(s) {
		runes = rune(s[n-1])
	}
	return runes
}

func main() {
	// z01.PrintRune(NRune("Hello!", 3))
	// z01.PrintRune(NRune("Salut!", 2))
	// z01.PrintRune(NRune("Bye!", -1))
	// z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
