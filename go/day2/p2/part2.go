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
type Input []Round
type Round [2]rune

func (r Round) play() (result int) {
	var (
		rock     rune = 'A'
		paper    rune = 'B'
		scissors rune = 'C'

		// Desired outcomes
		loss rune = 'X'
		draw rune = 'Y'
		win  rune = 'Z'
	)

	var lPlay, dPlay, wPlay rune // Potential plays to get the desired outcome
	switch r[0] {                // Opponent's move
	case rock: // Rock
		lPlay = scissors
		dPlay = rock
		wPlay = paper
	case paper: // Paper
		lPlay = rock
		dPlay = paper
		wPlay = scissors
	case scissors: // Scissors
		lPlay = paper
		dPlay = scissors
		wPlay = rock
	}
	switch r[1] { // Desired outcome
	case loss:
		result += r.playBonus(lPlay)
	case draw:
		result += 3
		result += r.playBonus(dPlay)
	case win:
		result += 6
		result += r.playBonus(wPlay)
	}
	return result
}

func (r Round) playBonus(play rune) int {
	switch play {
	case 'A':
		return 1
	case 'B':
		return 2
	case 'C':
		return 3
	}
	panic("Invalid play:" + string(play))
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

func Part2() {
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
