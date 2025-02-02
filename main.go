package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"go.cpmachado.pt/nqg/snq"
)

const (
	// DefaultN is the default side of board
	DefaultN = 8
	// DefaultOutputFile is the default output file, in which "" is os.Stdout
	DefaultOutputFile = ""
)

var (
	// Version holds the version of the current software based on compilation
	Version string = "1.0.0"
	n       int
	v       bool
	out     *os.File
)

func init() {
	var outputFile string
	flag.IntVar(&n, "n", DefaultN, fmt.Sprintf("Length of side of chess board in [1,%d]", DefaultN))
	flag.StringVar(&outputFile, "o", DefaultOutputFile, "Output file")
	flag.BoolVar(&v, "v", false, "Display version")
	flag.Usage = Usage
	flag.Parse()

	if v {
		printVersion()
		os.Exit(0)
	}
	if n < 1 || n > DefaultN {
		log.Fatalf("invalid n value, it should be in [1,%d]\n", DefaultN)
	}
	if outputFile != "" {
		var err error
		out, err = os.Create(outputFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		out = os.Stdout
	}
}

func main() {
	defer out.Close()
	pos := make([]int, n)
	snq.SNQ(out, pos, 0)
}

// printVersion Prints the version command
func printVersion() {
	fmt.Printf("%s-%s Copyright (C) 2025 Carlos Pinto Machado\n", os.Args[0], Version)
}

// Usage is a function to replace the default provided in the flag package
func Usage() {
	prog := os.Args[0]
	fmt.Fprintf(flag.CommandLine.Output(), "%s is a program to solve the N queen puzzle\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", prog)
	flag.PrintDefaults()
}
