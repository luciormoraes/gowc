package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// To read from stdin
	if len(os.Args) < 2 {
		fmt.Println("please pass a file name")
		os.Exit(1)
	}

	fileName := os.Args[1]

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}
	defer file.Close() // Close file when done

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
}
