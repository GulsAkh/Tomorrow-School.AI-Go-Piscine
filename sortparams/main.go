package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var str []rune
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		for _, el := range arg {
			str = append(str, el)
		}
	}
	for j := 0; j < len(str)-1; j++ {
		for k := 0; k < len(str)-j-1; k++ {
			if str[k] > str[k+1] {
				str[k], str[k+1] = str[k+1], str[k]
			}
		}
	}
	for _, sort := range str {
		z01.PrintRune(sort)
		z01.PrintRune('\n')
	}
}
