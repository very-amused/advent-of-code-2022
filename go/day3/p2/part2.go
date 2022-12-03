package p2

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
type Input []Group
type Group [3]string

func (g Group) badgeItem() rune {
	// Find common item
	itemMap := make(map[rune]int)
	for _, sack := range g {
		groupMap := make(map[rune]bool) // Items the group has
		for _, char := range sack {
			groupMap[char] = true
		}
		for char, t := range groupMap {
			if t {
				itemMap[char]++
				if itemMap[char] == 3 {
					return char
				}
			}
		}
	}
	panic("Failed to find common item: " + fmt.Sprint(g))
}

func (g Group) itemPriority(item rune) int {
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

	i := 0
	group := Group{} // Group buffer
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		group[i] = l
		i++
		if i%3 == 0 { // Split into groups
			input = append(input, group)
			i = 0
		}
	}
	return input
}

// Solve
func solve(groups Input) (solution string) {
	c := 0
	for _, g := range groups {
		c += g.itemPriority(g.badgeItem())
	}
	return strconv.Itoa(c)
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
