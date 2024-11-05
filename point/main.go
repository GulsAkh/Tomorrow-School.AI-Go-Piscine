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
	z01.PrintRune(52)
	z01.PrintRune(50)
	for _, chr := range "y = " {
		z01.PrintRune(chr)
	}
	z01.PrintRune(50)
	z01.PrintRune(49)
}
