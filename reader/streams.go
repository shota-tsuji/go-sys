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
	lReaderC := io.LimitReader(computer, 1)
	lReaderS := io.LimitReader(system, 1)
	sectionReader := io.NewSectionReader(programming, 5, 4)
	var bw bytes.Buffer// bytes.Reader
	tee := io.TeeReader(sectionReader, &bw)

	r, w := io.Pipe()
	var br bytes.Reader
	go func() {
	}

	//lr1 := io.LimitReader(&buffer, 
	lr2 := io.NewSectionReader(tee, 3, 1)
	//stream = io.MultiReader(programming, lReaderS, lReaderC)
	stream = io.MultiReader(sectionReader, lReaderS, lReaderC, lr2)
	io.Copy(os.Stdout, stream)
}
