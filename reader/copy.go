package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file2, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	io.Copy(file2, file)
}
