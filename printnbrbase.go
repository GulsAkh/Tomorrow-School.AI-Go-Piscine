package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	based := []rune(base)
	len_base := len(based)

	str := ""
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	for i := 0; i < len_base; i++ {
		for j := i + 1; j < len_base; j++ {
			if based[i] == based[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			} else if based[i] == '+' || based[i] == '-' {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if len_base < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == 0 {
		str = "0"
	} else {
		for nbr > 0 {
			str = string(base[nbr%len_base]) + str
			nbr /= len_base
		}
	}

	for _, el := range str {
		z01.PrintRune(el)
	}
}
