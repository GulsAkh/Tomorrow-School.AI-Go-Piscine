package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) *point {
	ptr.x = 42
	ptr.y = 21
	xChar := 'x'
	yChar := 'y'
	equalCh := '='
	spaceCh := ' '
	commaCh := ','
	newLine := '\n'

	z01.PrintRune(xChar)
	z01.PrintRune(spaceCh)
	z01.PrintRune(equalCh)
	z01.PrintRune(spaceCh)
	z01.PrintRune(52)
	z01.PrintRune(50)
	z01.PrintRune(commaCh)
	z01.PrintRune(spaceCh)
	z01.PrintRune(yChar)
	z01.PrintRune(spaceCh)
	z01.PrintRune(equalCh)
	z01.PrintRune(spaceCh)
	z01.PrintRune(50)
	z01.PrintRune(49)

	z01.PrintRune(newLine)
	return ptr
}

func main() {
	points := point{}
	setPoint(&points)
}
