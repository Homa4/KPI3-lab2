package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Homa4/KPI3-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File to write output (optional)")
)

func main() {
	flag.Parse()

	if (*inputExpression != "" && *inputFile != "") || (*inputExpression == "" && *inputFile == "") {
		fmt.Fprintln(os.Stderr, "Error: use either -e or -f, not both")
		flag.Usage()
		os.Exit(1)
	}

	var reader io.Reader
	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	var writer io.Writer = os.Stdout
	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error in calculating result:", err)
		os.Exit(1)
	}
}
