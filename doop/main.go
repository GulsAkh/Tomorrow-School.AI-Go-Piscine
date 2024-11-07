package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	num := 0
	sign := 1
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	for _, el := range s {
		if el < '0' || el > '9' {
			return 0
		}
		num = num*10 + int(el-'0')
	}
	return num * sign
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	if len(args[0]) >= 19 || len(args[2]) >= 19 {
		return
	}

	if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "*" && args[1] != "%" {
		return
	}
	var result int
	sign := args[1]
	num1 := Atoi(args[0])
	num2 := Atoi(args[2])
	if num1 == 0 || num2 == 0 {
		return
	}
	if sign == "*" {
		result = num1 * num2
	} else if sign == "+" {
		result = num1 + num2
	} else if sign == "-" {
		result = num1 - num2
	} else if sign == "/" {
		if num2 == 0 {
			fmt.Println("No division by 0")
			return
		} else {
			result = num1 / num2
		}
	} else if sign == "%" {
		if num2 == 0 {
			fmt.Println("No modulo by 0")
			return
		} else {
			result = num1 % num2
		}
	}
	fmt.Println(result)
}
