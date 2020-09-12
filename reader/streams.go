package main

import (
	"strings"
	"io"
	"os"
	//"bytes"
)

var (
	computer	= strings.NewReader("COMPUTER")
	system		= strings.NewReader("SYSTEM")
	programming	= strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	sr1 := io.NewSectionReader(computer, 0, 1)
	sr2 := io.NewSectionReader(system, 0, 1)
	sr3 := io.NewSectionReader(programming, 5, 1)
	sr4 := io.NewSectionReader(programming, 8, 1)
	sr5 := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(sr3, sr2, sr1, sr4, sr5)

	io.Copy(os.Stdout, stream)
}
