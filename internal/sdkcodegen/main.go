//+build sdkcodegen

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// TODO: error handling
	filename := os.Args[1]
	var destFilename string
	if len(os.Args) == 3 {
		destFilename = os.Args[2]
	} else {
		// blindly append `.go` so the result looks like `foo.md.go`
		destFilename = filename + ".go"
	}

	emitToStdout := destFilename == "-"

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

	var sink io.Writer
	if emitToStdout {
		sink = os.Stdout
	} else {
		file, err := os.Create(destFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open '%s' for writing failed: %+v\n", destFilename, err)
			os.Exit(1)
			return // unreachable
		}
		bufWriter := bufio.NewWriter(file)
		sink = bufWriter
		defer func() {
			bufWriter.Flush()
			file.Close()
		}()
	}
	em := &goEmitter{
		Sink: sink,
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
