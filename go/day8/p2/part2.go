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
type Trees [][]uint

func (t Trees) scenicScores() (scores []int) {
	for row := range t {
		for col := range t[row] {
			scores = append(scores, t.scenicScore(row, col))
		}
	}
	return scores
}

func (t Trees) viewDistance(line []uint, cell uint) (distance int) {
	for i, n := range line {
		if n >= cell {
			return i + 1
		}
	}
	return len(line)
}

func (t Trees) scenicScore(y, x int) int {
	// lines
	lLine := make([]uint, len(t[y][:x]))
	copy(lLine, t[y][:x])
	for i, j := 0, len(lLine)-1; i < j; i, j = i+1, j-1 {
		lLine[i], lLine[j] = lLine[j], lLine[i]
	}
	rLine := t[y][x+1:]
	var tLine []uint
	var bLine []uint
	for i := y - 1; i >= 0; i-- {
		tLine = append(tLine, t[i][x])
	}
	for i := y + 1; i < len(t); i++ {
		bLine = append(bLine, t[i][x])
	}
	cell := t[y][x]
	return t.viewDistance(lLine, cell) * t.viewDistance(rLine, cell) * t.viewDistance(tLine, cell) * t.viewDistance(bLine, cell)
}

// #endregion

// Parse
func parse() (input Trees) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parse lines

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		digits := strings.Split(l, "")
		input = append(input, make([]uint, len(digits)))
		for i, n := range digits {
			input[len(input)-1][i] = uint(must(strconv.Atoi(n)))
		}
	}
	return input
}

// Solve
func solve(input Trees) (solution string) {
	scores := input.scenicScores()
	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return strconv.Itoa(max)
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
