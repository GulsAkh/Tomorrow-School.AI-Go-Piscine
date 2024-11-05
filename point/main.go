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
	for _, chr := range "x = 42, y = 21\n" {
		z01.PrintRune(chr)
	}
}
