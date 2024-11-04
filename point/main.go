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
	// n1 := points.x
	// n2 := points.y

	// digit := []int{}

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
	z01.PrintRune('4')
	z01.PrintRune('2')
	z01.PrintRune(commaCh)
	z01.PrintRune(spaceCh)
	z01.PrintRune(yChar)
	z01.PrintRune(spaceCh)
	z01.PrintRune(equalCh)
	z01.PrintRune(spaceCh)
	z01.PrintRune('2')
	z01.PrintRune('1')

	z01.PrintRune(newLine)
}
