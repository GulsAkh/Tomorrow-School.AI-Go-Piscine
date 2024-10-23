package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var i, x, y int
	for i = 0; i <= 7; i++ {
		for x = i + 1; x <= 8; x++ {
			for y = x + 1; y <= 9; y++ {

				firstDigit := '0' + rune(i)
				secondDigit := '0' + rune(x)
				thirdDigit := '0' + rune(y)

				z01.PrintRune(firstDigit)
				z01.PrintRune(secondDigit)
				z01.PrintRune(thirdDigit)

				if i < 7 || (i == 7 && x < 8) || (i == 7 && x == 8 && y < 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
