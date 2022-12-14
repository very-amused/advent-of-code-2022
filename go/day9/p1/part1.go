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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// #region Structs
type Input []Instruction

type Bridge struct {
	head *Node
	tail *Node
}

func (b *Bridge) move(i Instruction) {
	// Move head
	switch i.direction {
	case 'U':
		b.head.move(b.head.x, b.head.y+i.magnitude)
	case 'D':
		b.head.move(b.head.x, b.head.y-i.magnitude)
	case 'L':
		b.head.move(b.head.x-i.magnitude, b.head.y)
	case 'R':
		b.head.move(b.head.x+i.magnitude, b.head.y)
	}
	{
		xDistance := abs(b.head.x - b.tail.x)
		yDistance := abs(b.head.y - b.tail.y)
		// Tail follows
		// Only xDistance or yDistance can be > 1, never both
		if xDistance > 0 && yDistance > 1 { // Join col
			b.tail.x = b.head.x
		} else if yDistance > 0 && xDistance > 1 { // Join row
			b.tail.y = b.head.y
		}
	}
	// Horizontal moves
	for (b.head.x - b.tail.x) > 1 {
		b.tail.move(b.tail.x+1, b.tail.y)
	}
	for (b.tail.x - b.head.x) > 1 {
		b.tail.move(b.tail.x-1, b.tail.y)
	}
	// Vertical moves
	for (b.head.y - b.tail.y) > 1 {
		b.tail.move(b.tail.x, b.tail.y+1)
	}
	for (b.tail.y - b.head.y) > 1 {
		b.tail.move(b.tail.x, b.tail.y-1)
	}
}

type Instruction struct {
	direction rune
	magnitude int
}

type Node struct {
	x int
	y int

	history map[[2]int]bool // x,y coords that have been visited, does NOT include current position
}

func (n *Node) move(x, y int) {
	n.x = x
	n.y = y
	n.visited() // add position to history
}

func (n *Node) visited() {
	n.history[[2]int{n.x, n.y}] = true
}

// #endregion

// Parse
func parse() (instructions Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parse lines

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		var i Instruction
		must(fmt.Sscanf(l, "%c %d", &i.direction, &i.magnitude))
		instructions = append(instructions, i)
	}
	return instructions
}

// Solve
func solve(ops Input) (solution string) {
	// Initial bridge state
	bridge := Bridge{
		head: &Node{history: make(map[[2]int]bool)},
		tail: &Node{history: make(map[[2]int]bool)}}
	// Mark initial positions as visited
	bridge.head.visited()
	bridge.tail.visited()
	for _, op := range ops {
		bridge.move(op)
	}
	return strconv.Itoa(len(bridge.tail.history))
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
