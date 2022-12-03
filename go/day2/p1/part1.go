package p1

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
type Input []Round
type Round [2]rune

func (r Round) play() (result int) {
	var (
		rock     = [2]rune{'A', 'X'}
		paper    = [2]rune{'B', 'Y'}
		scissors = [2]rune{'C', 'Z'}
	)

	var drawMatch, winMatch rune // Opposing plays that result in draw/win
	switch r[1] {                // Your move
	case rock[1]: // Rock
		result += 1
		drawMatch = rock[0]
		winMatch = scissors[0]
	case paper[1]: // Paper
		result += 2
		drawMatch = paper[0]
		winMatch = rock[0]
	case scissors[1]: // Scissors
		result += 3
		drawMatch = scissors[0]
		winMatch = paper[0]
	}
	if r[0] == drawMatch {
		result += 3
	} else if r[0] == winMatch {
		result += 6
	}
	return result
}

// #endregion

// Parse
func parse() (rounds Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parse lines

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		parts := strings.Split(l, " ")
		if len(parts) != 2 {
			panic("Malformed line: " + l)
		}
		round := Round{rune(parts[0][0]), rune(parts[1][0])}
		rounds = append(rounds, round)
	}
	return rounds
}

// Solve
func solve(rounds Input) (solution string) {
	count := 0
	for _, round := range rounds {
		count += round.play()
	}
	return strconv.Itoa(count)
}

func Part1() {
	// Parse
	rounds := parse()

	// Solve
	start := time.Now()
	solution := solve(rounds)

	// Report solve time and solution
	duration := time.Now().Sub(start)
	fmt.Printf("Solved in \x1b[34m%s\x1b[0m\n", duration)
	fmt.Printf("Solution: \x1b[32m%s\x1b[0m\n", solution)
}
