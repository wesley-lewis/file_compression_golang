package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {
	name_of_file := "file1.txt"

	// The os.Open() method just opens the file and does not read the file
	// It returns two values first one is the pointer to the file and the second one is the error. If no error encountered then error is nil
	f, err := os.Open("C:\\Users\\wesle\\OneDrive\\Desktop\\go\\src\\file_compression\\" + name_of_file)
	// Checking if an error has occurred or not
	if err != nil {
		log.Fatal(err)
	}

	// Creates a new buffered reader
	// Here we pass the file to the NewReader() method
	read := bufio.NewReader(f)

	// data stores the contents of the file
	data, _ := io.ReadAll(read)

	// Creating the new file name which will have extension gzip
	name_of_file = strings.Replace(name_of_file, ".txt", ".gz", -1)

	// Creating a new file with .gzip extension
	f, _ = os.Create("C:\\Users\\wesle\\OneDrive\\Desktop\\go\\src\\file_compression\\" + name_of_file)

	// Creating a new gzip writer to write to the gzip file which we created above
	w := gzip.NewWriter(f)

	// Writing to name_of_file.gzip
	w.Write(data)

	// Closing the writer
	w.Close()

	fmt.Println("File is converted succesfully")
}