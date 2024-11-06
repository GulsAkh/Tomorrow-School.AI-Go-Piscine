package main

import "fmt"

func PrintNbr(num int) {
	fmt.Print(num)
}

func ForEach(f func(int), a []int) {
	for _, el := range a {
		f(el)
	}
	fmt.Println()
}
