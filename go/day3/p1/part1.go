package p1

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
type Input []Rucksack
type Rucksack string

func (s Rucksack) commonItem() rune {
	// Split compartment in 2
	c1 := s[:len(s)/2]
	c2 := s[len(s)/2:]
	// Find common item
	itemMap := make(map[rune]bool)
	for _, char := range c1 {
		itemMap[char] = true
	}
	for _, item := range c2 {
		if itemMap[item] {
			return item
		}
	}
	panic("Failed to find common item: " + s)
}

func (s Rucksack) itemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1 // 1-26
	} else if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27 // 27-52
	}
	panic("Invalid item char: " + string(item))
}

// #endregion

// Parse
func parse() (input Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		input = append(input, Rucksack(l))
	}
	return input
}

// Solve
func solve(input Input) (solution string) {
	c := 0
	for _, sack := range input {
		c += sack.itemPriority(sack.commonItem())
	}
	return strconv.Itoa(c)
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
