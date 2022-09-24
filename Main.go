package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	name_of_file := "file1.txt"
	f, _ := os.Open("C:\\Users\\wesle\\OneDrive\\Desktop\\go\\src\\file_compression\\" + name_of_file)

	read := bufio.NewReader(f)

	data, _ := ioutil.ReadAll(read)

	name_of_file = strings.Replace(name_of_file, ".txt", ".gz", -1)

	f, _ = os.Create("C:\\Users\\wesle\\OneDrive\\Desktop\\go\\src\\file_compression\\" + name_of_file)
}