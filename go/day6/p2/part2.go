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

// Parse and solve
func solve() (solution string) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	const markerLen = 14
	processed := 0
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		chars := strings.Split(l, "")
	outer:
		for i := 0; (i + markerLen - 1) < len(l); i++ {
			// Verify group of chars is unique
			c := make(map[string]bool)
			for j := 0; j < markerLen; j++ {
				if c[chars[i+j]] {
					continue outer
				}
				c[chars[i+j]] = true
			}
			processed = i + markerLen
			break
		}
	}
	return strconv.Itoa(processed)
}

func Part2() {
	// Solve
	start := time.Now()
	solution := solve()

	// Report solve time and solution
	duration := time.Now().Sub(start)
	fmt.Printf("Solved in \x1b[34m%s\x1b[0m\n", duration)
	fmt.Printf("Solution: \x1b[32m%s\x1b[0m\n", solution)
}
