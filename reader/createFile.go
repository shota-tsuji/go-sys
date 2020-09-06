package main

import (
	"io"
	"os"
	"crypto/rand"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//buffer := make([]byte, 1024)
	_, err = io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
}
