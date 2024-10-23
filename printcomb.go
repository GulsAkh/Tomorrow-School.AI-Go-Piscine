package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for x := i + 1; x <= 8; x++ {
			for y := x + 1; y <= 9; y++ {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(x + '0'))
				z01.PrintRune(rune(y + '0'))

				if i < 7 || (i == 7 && x < 8) || (i == 7 && x == 8 && y < 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
