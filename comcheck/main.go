package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, el := range args {
		if el == "01" || el == "galaxy" || el == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
