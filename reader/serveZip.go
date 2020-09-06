package main

import (
	"net/http"
	"io"
	"archive/zip"
	"os"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")
	//io.WriteString(w, "http.ResponseWriter sample")
	//w.Write([]byte("http.ResponseWriter sample"))
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	filenames := [...] string {"file1-for-zip.txt", "file2-for-zip.txt"}
	for _, fileName := range filenames {
		fmt.Println(fileName)
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer, err := zipWriter.Create(fileName)
		io.Copy(writer, file)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8085", nil)
}
