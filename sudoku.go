//package main

import (
	"fmt"
	"os"
)

func main() {
	var digit []int
	for i := 1; i < len(os.Args); i++ {
		args := os.Args[i]
		for _, el := range args {
			if el == '.' {
				digit = append(digit, 0)
			} else {
				digit = append(digit, int(el-'0'))
			}
		}
	}
	for i := 0; i < len(digit); i++ {
		if i > 0 && i%3 == 0 {
			fmt.Println()
		}
		fmt.Print(digit[i])
	}
	fmt.Println()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := []int{1, 2, 3}
			if i == 0 {
				if j == 0 {
					fmt.Print(digit[i])
					fmt.Println()
				}
			}
		}
	}
}
