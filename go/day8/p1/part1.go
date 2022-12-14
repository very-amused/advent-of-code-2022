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
type Trees [][]uint

func (t Trees) exteriorVisible() int {
	return (2 * len(t[0])) + (2 * (len(t) - 2))
}

func (t Trees) interiorVisible() (c int) {
	// For each interior col of each interior row, check visibility
	for row := 1; row < len(t)-1; row++ {
		for col := 1; col < len(t[row])-1; col++ {
			if t.isVisible(col, row) {
				c++
			}
		}
	}
	return c
}

func (t Trees) checkLine(line []uint, cell uint) bool {
	for _, n := range line {
		if n >= cell {
			return false
		}
	}
	return true
}

func (t Trees) isVisible(x, y int) bool {
	// lines
	lLine := t[y][:x]
	rLine := t[y][x+1:]
	var tLine []uint
	var bLine []uint
	for i := 0; i < y; i++ {
		tLine = append(tLine, t[i][x])
	}
	for i := len(t) - 1; i > y; i-- {
		bLine = append(bLine, t[i][x])
	}
	cell := t[y][x]
	return t.checkLine(lLine, cell) || t.checkLine(rLine, cell) || t.checkLine(tLine, cell) || t.checkLine(bLine, cell)
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
	return strconv.Itoa(input.exteriorVisible() + input.interiorVisible())
}

func Part1() {
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
