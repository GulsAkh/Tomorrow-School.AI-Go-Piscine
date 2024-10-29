package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var runes []rune
	if n < 0 {
		return
	}
	// converting int to runes
	if n == 0 {
		runes = append(runes, '0')
	} else {
		for n > 0 {
			digit := n % 10
			runes = append(runes, rune(digit+'0'))
			n = n / 10
		}
	}

	// sorting runes
	num := len(runes)
	for i := 0; i < num; i++ {
		for j := 0; j < num-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	// printing sorted runes
	for _, el := range string(runes) {
		z01.PrintRune(el)
	}
	z01.PrintRune(10)
}
