package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	// DefaultN is the default side of board
	DefaultN = 8
	// DefaultOutputFile is the default output file, in which "" is os.Stdout
	DefaultOutputFile = ""
)

// Version holds the version of the current software based on compilation

var (
	Version string = "unknown"
	n       int
	v       bool
	out     *os.File
)

func init() {
	var outputFile string
	flag.IntVar(&n, "n", DefaultN, "Length of side of chess board")
	flag.StringVar(&outputFile, "o", DefaultOutputFile, "Output file")
	flag.BoolVar(&v, "v", false, "Display version")
	flag.Usage = Usage
	flag.Parse()

	if v {
		printVersion()
		os.Exit(0)
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
	fmt.Fprintf(out, "Hello there")
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
