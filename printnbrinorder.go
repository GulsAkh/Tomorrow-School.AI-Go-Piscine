package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	for i := 0; i < n; i++ {
		if n > 0 {
			z01.PrintRune(rune(n%10) + '0')
			n = n / 10
		}
	}
}
