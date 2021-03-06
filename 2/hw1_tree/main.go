package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	//"strings"
)

var printFiles bool

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, filePath string, printFiles bool) error {
	fmt.Fprintln(out, "hello cold")

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {

	})

	return nil
}