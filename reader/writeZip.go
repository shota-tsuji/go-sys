package main

import (
	"io"
	"os"
	"archive/zip"
	"fmt"
)

func main() {
	zipFile, err := os.Create("a.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	filenames := [...] string {"file1-for-zip.txt", "file2-for-zip.txt"}
	for _, fileName := range filenames {
		fmt.Println(fileName)
		//fileName := "file1-for-zip.txt"
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer, err := zipWriter.Create(fileName)
		io.Copy(writer, file)
	}
	//fileName := "file1-for-zip.txt"
	//file, err := os.Open(fileName)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//writer, err := zipWriter.Create(fileName)
	//io.Copy(writer, file)

	//fileName := "file2-for-zip.txt"
	//file, err := os.Open(fileName)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//writer, err := zipWriter.Create(fileName)
	//io.Copy(writer, file)
}
