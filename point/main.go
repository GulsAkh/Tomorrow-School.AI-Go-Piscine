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

	num1 := [2]rune{4, 2}

	for _, digit := range num1 {
		z01.PrintRune(48 + digit)
	}

	for _, chr := range ", y = " {
		z01.PrintRune(chr)
	}
	num2 := [3]rune{2, 1, -38}

	for _, digit := range num2 {
		z01.PrintRune(48 + digit)
	}
}
