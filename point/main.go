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

func printNum(num int) {
	digit := []int{}
	for num > 0 {
		digit = append(digit, num%10)
		num /= 10
	}

	for i := len(digit) - 1; i >= 0; i-- {
		z01.PrintRune(rune(48 + digit[i]))
	}
}

func main() {
	points := point{}
	setPoint(&points)
	n1 := points.x
	n2 := points.y
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
	printNum(n1)
	z01.PrintRune(commaCh)
	z01.PrintRune(spaceCh)
	z01.PrintRune(yChar)
	z01.PrintRune(spaceCh)
	z01.PrintRune(equalCh)
	z01.PrintRune(spaceCh)
	printNum(n2)

	z01.PrintRune(newLine)
}
