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
	n1 := points.x
	n2 := points.y
	str1 := string(rune('0'+n1/10)) + string(rune('0'+n1%10))
	str2 := string(rune('0'+n2/10)) + string(rune('0'+n2%10))

	message := "x = " + str1 + ", " + "y = " + str2
	for _, r := range message {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
