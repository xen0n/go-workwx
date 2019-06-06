package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// TODO: error handling
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open input file failed: %+v\n", err)
		os.Exit(1)
		return // unreachable
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input failed: %+v\n", err)
		os.Exit(1)
		return // unreachable
	}

	mdRoot := parseDocument(content)
	analyzeDocument(mdRoot)
}
