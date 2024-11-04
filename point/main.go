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

func convertInt(n int) string {
	var str string

	if n == 0 {
		str = "0"
	} else {
		if n < 0 {
			str = "-"
			n = -n
		}

		for n > 0 {
			digit := n % 10
			str = string(rune('0'+digit)) + str
			n = n / 10
		}
	}
	return str
}

func printWithRune(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	points := point{}

	setPoint(&points)
	str1 := convertInt(points.x)
	str2 := convertInt(points.y)

	message := "x = " + str1 + ", " + "y = " + str2

	printWithRune(message)
}
