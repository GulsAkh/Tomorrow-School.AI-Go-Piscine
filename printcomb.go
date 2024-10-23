package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var i, x, y int
	for i = 0; i <= 1; i++ {
		for x = 1; x <= 2; x++ {
			for y = 2; y <= 3; y++ {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(x + '0'))
				z01.PrintRune(rune(y + '0'))

				if i != 7 && x != 8 && y != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
			y = y + 1
		}
		x = x + 1
	}
	z01.PrintRune('\n')
}
