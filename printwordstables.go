package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, el := range a {
		for _, char := range el {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
