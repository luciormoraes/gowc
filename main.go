package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var showLines, showBytes, showWords bool
	flag.BoolVar(&showBytes, "c", false, "print the byte counts")
	flag.BoolVar(&showLines, "l", false, "print the newline counts")
	flag.BoolVar(&showWords, "w", false, "print the word counts")
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

	// fmt.Printf("%d %d %d %s\n", lines, words, characters, fileName)

	// Display counts by default if no flags are provided
	if !showLines && !showBytes && !showWords {
		fmt.Printf("%d %d %d %s\n", lines, words, characters, fileName)
	} else {
		if showLines {
			fmt.Printf("%d %s\n", lines, fileName)
		}

		if showBytes {
			fmt.Printf("%d %s\n", characters, fileName)
		}

		if showWords {
			fmt.Printf("%d %s\n", words, fileName)
		}
	}
}
