package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var showLines, showBytes, showWords, showMultiBytes bool
	flag.BoolVar(&showBytes, "c", false, "print the byte counts")
	flag.BoolVar(&showLines, "l", false, "print the newline counts")
	flag.BoolVar(&showWords, "w", false, "print the word counts")
	flag.BoolVar(&showMultiBytes, "m", false, "print the character counts")
	flag.Parse()

	// Check if data is being piped in
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		processPipedInput(showLines, showBytes, showWords)
		return
	}

	// Get file name from command line
	fileName := flag.CommandLine.Arg(0)
	if flag.NArg() > 0 { // flag.NArg() returns the number of arguments remaining after flags have been processed
		fileName = flag.Arg(0)
	} else {
		fmt.Println("Provide a file name")
		os.Exit(1)
	}

	// Open file
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}

	// Display counts by default if no flags are provided
	if !showLines && !showBytes && !showWords && !showMultiBytes {
		lines, words, bytes := countLines(file), countWords(file), len(file)
		fmt.Printf("%d %d %d %s\n", lines, words, bytes, fileName)
	} else {
		if showLines {
			fmt.Printf("%d %s\n", countLines(file), fileName)
		}

		if showBytes {
			fmt.Printf("%d %s\n", len(file), fileName)
		}

		if showWords {
			fmt.Printf("%d %s\n", countWords(file), fileName)
		}

		if showMultiBytes {
			fmt.Printf("%d %s\n", countMultiBytes(file), fileName)
		}
	}

	// defer file.Close() // Close file when done
}

func countLines(content []byte) int {
	count := 0
	for _, b := range content {
		if b == '\n' {
			count++
		}
	}
	return count
}

func countWords(content []byte) int {
	words := 0
	fields := strings.Fields(string(content))
	words = len(fields)
	return words
}

func countMultiBytes(content []byte) int {
	return utf8.RuneCount(content)
}

func processPipedInput(showLines, showBytes, showWords bool) {
	fmt.Println("Reading from stdin...")
	scanner := bufio.NewScanner(os.Stdin)
	lines, bytes, words := 0, 0, 0

	for scanner.Scan() {
		lines++

		line := scanner.Text()
		bytes += len(line) + 1 // Include the newline character
		words += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	if !showLines && !showBytes && !showWords {
		fmt.Printf("%d %d %d\n", lines, words, bytes)
	} else {
		if showLines {
			fmt.Printf("%d\n", lines)
		}

		if showBytes {
			fmt.Printf("%d\n", bytes)
		}

		if showWords {
			fmt.Printf("%d\n", words)
		}
	}
}
