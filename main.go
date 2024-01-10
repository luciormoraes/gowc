package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var showLines bool
	flag.BoolVar(&showLines, "l", false, "Show lines")
	flag.Parse()

	// Get file name from command line
	fileName := flag.CommandLine.Arg(0)
	if flag.NArg() > 0 { // flag.NArg() returns the number of arguments remaining after flags have been processed
		fileName = flag.Arg(0)
	} else {
		fmt.Println("Provide a file name")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}
	defer file.Close() // Close file when done

	// Scan file
	scanner := bufio.NewScanner(file)
	lines, words, characters := 0, 0, 0

	for scanner.Scan() {
		lines++

		line := scanner.Text()
		characters += len(line) + 1 // Include the newline character

		// Count words using strings.Fields for accurate word count
		words += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Printf("%d %d %d %s\n", lines, words, characters, fileName)

	if showLines {
		fmt.Printf("%d %s\n", lines, fileName)
	}
}
