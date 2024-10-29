package main

import "fmt"

func Capitalize(s string) {
	var str []string

	for _, el := range s {
		if (el >= '0' && el <= '9') || (el >= 'A' && el <= 'Z') || (el >= 'a' && el <= 'z') {
			str = append(str, string(el))
		}
	}
	fmt.Print(str[5])
}

func main() {
	Capitalize("Hello! How are you?")
}
