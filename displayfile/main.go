package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
	}
	fileName := os.Args[1]
	file, err := os.OpenFile(fileName, os.O_RDWR, 0o666)
	if err != nil {
		fmt.Println("Error handling file", err)
		return
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	fmt.Println(string(content))
	defer file.Close()
}
