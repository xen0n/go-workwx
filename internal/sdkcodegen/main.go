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

	hir, err := analyzeDocument(mdRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "syntax error in spec: %+v\n", err)
		os.Exit(1)
		return // unreachable
	}

	em := &goEmitter{
		Sink: os.Stdout,
	}

	err = em.EmitCode(&hir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "code emission failed: %+v\n", err)
		os.Exit(1)
		return // unreachable
	}

	err = em.Finalize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "finalization failed: %+v\n", err)
		os.Exit(1)
		return // unreachable
	}
}
