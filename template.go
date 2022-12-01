package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const inputFile = "sample.txt" // Change to input.txt for final solution

func must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

// #region Structs
type Input any

// #endregion

// Parse input
func parse() (input Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parsing state vars go here (if any)

	for scanner.Scan() {
		//l := scanner.Text()
	}

	return input
}

// Solve problem
func solve(input Input) (solution string) {
	return "" // Placeholder
}

func main() {
	// Parse
	input := parse()

	// Solve
	start := time.Now()
	solution := solve(input)

	// Report solve time and solution
	duration := time.Now().Sub(start)
	fmt.Printf("Solved in \x1b[34m%s\x1b[0m\n", duration)
	fmt.Printf("Solution: \x1b[32m%s\x1b[0m\n", solution)
}
