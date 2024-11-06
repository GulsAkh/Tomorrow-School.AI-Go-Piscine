package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error handling file", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	contentStr := string(content)
	contentStr = strings.TrimSpace(contentStr)

	fmt.Println(contentStr)
}
