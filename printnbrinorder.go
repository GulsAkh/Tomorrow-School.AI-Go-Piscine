package main

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

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
