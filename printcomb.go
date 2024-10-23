package piscine

import piscine

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for x := i + 1; x <= 8; x++ {
			for y := x + 1; y <= 9; y++ {
				PrintRune(rune(i) + 48)
				PrintRune(rune(x) + 48)
				PrintRune(rune(y) + 48)

				if i < 7 || (i == 7 && x < 8) || (i == 7 && x == 8 && y < 9) {
					PrintRune(',')
					PrintRune(' ')
				}
			}
		}
	}
	PrintRune('\n')
}
