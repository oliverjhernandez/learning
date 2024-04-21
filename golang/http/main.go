package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	logger := logWriter{}

	io.Copy(logger, resp.Body)
	// myByte := make([]byte, 99999)
	// resp.Body.Read(myByte)
	// fmt.Println(string(myByte))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
