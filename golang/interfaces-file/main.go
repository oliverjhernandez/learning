package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	fileName := args[1]
	contentFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Could not open file.")
		os.Exit(1)
	}
	io.Copy(os.Stdout, contentFile)
}
