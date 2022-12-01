package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const inputFile = "sample.txt" // Change to input.txt for final solution
const ignoreEmptyLines = true  // Don't parse blank lines

func must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

// #region Structs

// #endregion

// Parse input
func parse() {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parsing state vars go here (if any)

	for scanner.Scan() {
		l := scanner.Text()
		if ignoreEmptyLines && len(l) == 0 {
			continue
		}

	}
}

// Solve problem
func solve() (solution string) {
	return "" // Placeholder
}

func main() {
	// Parse
	fmt.Printf("Parsing input (%s)\n", inputFile)
	parse()

	// Solve
	fmt.Println("Solving")
	start := time.Now()
	solution := solve()

	// Report solve time and solution
	duration := time.Now().Sub(start)
	fmt.Println("Solved in", duration)
	fmt.Println("Solution:", solution)
}
