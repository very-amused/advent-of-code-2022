package main

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

// #endregion

// Solve
func solve() (solution string) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parsing state vars go here (if any)
	var elves []int
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
	// sum top 3
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i] < elves[j]
	})
	sum := 0
	for _, c := range elves[len(elves)-3:] {
		sum += c
	}
	return strconv.Itoa(sum)
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
