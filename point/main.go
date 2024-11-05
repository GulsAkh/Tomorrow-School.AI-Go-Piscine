package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) *point {
	ptr.x = 42
	ptr.y = 21
	return ptr
}

func main() {
	points := point{}
	setPoint(&points)
	for _, chr := range "x = " {
		z01.PrintRune(chr)
	}
	for i := 0; i < 2; i++ {
		if i == 0 {
			z01.PrintRune(rune('0' + points.x/10))
		} else {
			z01.PrintRune(rune('0' + points.x%10))
		}
	}

	for _, chr := range ", y = " {
		z01.PrintRune(chr)
	}
	for i := 0; i < 2; i++ {
		if i == 0 {
			z01.PrintRune(rune('0' + points.y/10))
		} else {
			z01.PrintRune(rune('0' + points.y%10))
		}
	}
}
