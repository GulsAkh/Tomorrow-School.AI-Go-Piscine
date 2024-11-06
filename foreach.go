package piscine

import (
	"fmt"
)

func ForEach(f func(int), a []int) {
	for _, el := range a {
		f(el)
	}
	fmt.Println()
}
