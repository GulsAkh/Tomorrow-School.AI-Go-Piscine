package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for y := '0'; y <= '9'; y++ {
			z01.PrintRune(i)
			z01.PrintRune(y)
			if !(i == '9' && y == '9') {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
