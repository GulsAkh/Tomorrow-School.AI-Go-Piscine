package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	r := []rune(s)

	for _, el := range r {
		z01.PrintRune(el)
	}
}
