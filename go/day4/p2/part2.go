package p2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
type Input []Pair
type Pair [2]Assignment

func (p Pair) overlaps() bool {
	return p[0].overlaps(p[1]) || p[1].overlaps(p[0])
}

type Assignment struct {
	start int
	end   int
}

func (a Assignment) overlaps(a2 Assignment) bool {
	return (a.start >= a2.start && a.start <= a2.end) ||
		(a.end <= a2.end && a.end >= a2.start)
}
func assignmentFrom(r string) (a Assignment) {
	parts := strings.Split(r, "-")
	if len(parts) != 2 {
		panic("Unexpected number of parts in assignment range: " + r)
	}
	a.start = must(strconv.Atoi(parts[0]))
	a.end = must(strconv.Atoi(parts[1]))
	return a
}

// #endregion

// Parse
func parse() (pairs Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parse lines

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		parts := strings.Split(l, ",")
		pairs = append(pairs, Pair{assignmentFrom(parts[0]), assignmentFrom(parts[1])})
	}
	return pairs
}

// Solve
func solve(pairs Input) (solution string) {
	c := 0
	for _, p := range pairs {
		if p.overlaps() {
			c++
		}
	}
	return strconv.Itoa(c)
}

func Part2() {
	// Parse
	pairs := parse()

	// Solve
	start := time.Now()
	solution := solve(pairs)

	// Report solve time and solution
	duration := time.Now().Sub(start)
	fmt.Printf("Solved in \x1b[34m%s\x1b[0m\n", duration)
	fmt.Printf("Solution: \x1b[32m%s\x1b[0m\n", solution)
}
