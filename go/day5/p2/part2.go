package p2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
type Input []Stack
type Stack []rune

// #endregion

// Parse

// Parse row of stack spaces (crates or empty)
func parseRow(stacks Input, row string) {
	spaceRegex := regexp.MustCompile(fmt.Sprintf("\\[[A-Z]\\]|%s", strings.Repeat(" ", 4)))
	spaces := spaceRegex.FindAllString(row, -1)
	for i, space := range spaces {
		if !strings.HasPrefix(space, "[") {
			continue // Ignore empty space
		}
		stacks[i] = append(stacks[i], rune(space[1])) // Append stack content
	}
}

func parse() (stacks Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Get number of initial stacks from first line
	const stackLen = 3 // # of chars composing a stack, stacks are space separated, so all but 1 per row take up 4 characters in actuality

	// Parse lines

	stacks = nil
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		// Allocate crates from initial line
		if stacks == nil {
			nStacks := (len(l) + 1) / 4
			stacks = make([]Stack, nStacks)
		}

		// Construct initial stacks
		if strings.HasPrefix(l, strings.Repeat(" ", 3)) || strings.HasPrefix(l, "[") {
			parseRow(stacks, l)
			continue
		}

		// Done parsing stacks, reverse from insertion order
		if strings.HasPrefix(l, " 1") {
			for i := range stacks {
				for j, k := 0, len(stacks[i])-1; j < k; j, k = j+1, k-1 {
					stacks[i][j], stacks[i][k] = stacks[i][k], stacks[i][j]
				}
			}
			continue
		}

		// Process move instructions
		if strings.HasPrefix(l, "move") {
			var n int
			var from, to uint
			fmt.Sscanf(l, "move %d from %d to %d", &n, &from, &to)
			// Use 0 based indexes
			from--
			to--
			// Crane move point
			movePoint := len(stacks[from]) - n
			stacks[to] = append(stacks[to], stacks[from][movePoint:]...)
			stacks[from] = stacks[from][:movePoint]
		}
	}
	return stacks
}

// Solve
func solve(stacks Input) (solution string) {
	// Get top crates for each stack
	var s strings.Builder
	for _, stack := range stacks {
		s.WriteRune(stack[len(stack)-1])
	}

	return s.String()
}

func Part2() {
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
