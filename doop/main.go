package main

import (
	"os"
)

func Atoi(s string) (int, bool) {
	num := 0
	sign := 1
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	for _, el := range s {
		if el < '0' || el > '9' {
			return 0, false
		}
		num = num*10 + int(el-'0')
	}
	return num * sign, true
}

func Itoa(n int) string {
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	str := ""
	if n == 0 {
		str = "0"
	} else {
		for n > 0 {
			digit := n % 10
			str = string(digit+'0') + str
			n /= 10
		}
	}
	if negative {
		str = "-" + str
	}
	return str
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
	num1, valid1 := Atoi(args[0])
	if valid1 {
		num1 = num1
	} else {
		return
	}
	num2, valid2 := Atoi(args[2])
	if valid2 {
		num2 = num2
	} else {
		return
	}

	sign := args[1]
	var result int

	if sign == "*" {
		result = num1 * num2
	} else if sign == "+" {
		result = num1 + num2
	} else if sign == "-" {
		result = num1 - num2
	} else if sign == "/" {
		if num2 == 0 {
			mesDiv := "No division by 0"
			os.Stdout.Write([]byte(mesDiv))
			os.Stdout.Write([]byte("\n"))
			return
		} else {
			result = num1 / num2
		}
	} else if sign == "%" {
		if num2 == 0 {
			mesMod := "No modulo by 0"
			os.Stdout.Write([]byte(mesMod))
			os.Stdout.Write([]byte("\n"))
			return
		} else {
			result = num1 % num2
		}
	}

	str := Itoa(result)
	os.Stdout.Write([]byte(str))
	os.Stdout.Write([]byte("\n"))
}
