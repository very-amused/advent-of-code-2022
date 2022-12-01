package p1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
type Input []int

// #endregion

// Parse
func parse() (elves Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parsing state vars go here (if any)
	c := 0
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			elves = append(elves, c)
			c = 0
			continue
		}

		cals := must(strconv.Atoi(l))
		c += cals
	}
	return elves
}

// Solve
func solve(elves Input) (solution string) {
	// top elf value
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] < elves[j]
	})
	return strconv.Itoa(elves[len(elves)-1])
}

func Part1() {
	// Parse
	elves := parse()

	// Solve
	start := time.Now()
	solution := solve(elves)

	// Report solve time and solution
	duration := time.Now().Sub(start)
	fmt.Printf("Solved in \x1b[34m%s\x1b[0m\n", duration)
	fmt.Printf("Solution: \x1b[32m%s\x1b[0m\n", solution)
}
