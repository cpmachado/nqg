package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DefaultN          = 8
	DefaultOutputFile = ""
)

var Version string = "unknown"

func main() {
	var (
		n          int
		outputFile string
		v          bool
	)
	flag.IntVar(&n, "n", DefaultN, "Length of side of chess board")
	flag.StringVar(&outputFile, "o", DefaultOutputFile, "Output file")
	flag.BoolVar(&v, "v", false, "Display version")
	flag.Usage = Usage
	flag.Parse()

	if v {
		printVersion()
		os.Exit(0)
	}
	fmt.Println(n, outputFile, v)
}

func printVersion() {
	fmt.Printf("%s-%s Copyright (C) 2025 Carlos Pinto Machado\n", os.Args[0], Version)
}

func Usage() {
	prog := os.Args[0]
	fmt.Fprintf(flag.CommandLine.Output(), "%s is a program to solve the N queen puzzle\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", prog)
	flag.PrintDefaults()
}
