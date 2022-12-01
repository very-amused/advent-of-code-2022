package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const inputFile = "input.txt"  // Change to input.txt for final solution
const ignoreEmptyLines = false // Don't parse blank lines

func must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

// #region Structs

// #endregion

// Solve
func solve() (solution string) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parsing state vars go here (if any)
	max := 0
	c := 0
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			c = 0
			continue
		}

		cals := must(strconv.Atoi(l))
		c += cals
		if c > max {
			max = c
		}
	}
	return strconv.Itoa(max)
}

func main() {
	// Solve
	fmt.Println("Solving")
	start := time.Now()
	solution := solve()

	// Report solve time and solution
	duration := time.Now().Sub(start)
	fmt.Println("Solved in", duration)
	fmt.Println("Solution:", solution)
}
