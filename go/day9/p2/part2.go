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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// #region Structs
type Input []Instruction

type Bridge struct {
	nodes [10]*Node
}

func makeBridge() (b Bridge) {
	for i := range b.nodes {
		b.nodes[i] = makeNode()
	}
	return b
}

func (b *Bridge) move(op Instruction) {
	for i := 0; i < op.magnitude; i++ {
		b.moveHead(op.direction)
	}
}

func (b *Bridge) moveHead(dir rune) {
	// Move head via op
	head := b.nodes[0]
	switch dir {
	case 'U':
		head.move(head.x, head.y+1)
	case 'D':
		head.move(head.x, head.y-1)
	case 'L':
		head.move(head.x-1, head.y)
	case 'R':
		head.move(head.x+1, head.y)
	}
	// Move tails
	for i := 1; i < len(b.nodes); i++ {
		head := b.nodes[i-1]
		tail := b.nodes[i]
		{
			dx := func() int { return tail.x - head.x }
			dy := func() int { return tail.y - head.y }
			// Diagonal processing
			yDiag := func(y int) int {
				if dy() > 0 {
					y--
				} else if dy() < 0 {
					y++
				}
				return y
			}
			xDiag := func(x int) int {
				if dx() > 0 {
					x--
				} else if dx() < 0 {
					x++
				}
				return x
			}
			// Movement
			for abs(dx()) > 1 || abs(dy()) > 1 {
				x := tail.x
				y := tail.y
				if dx() > 1 {
					x--
					y = yDiag(y)
				} else if dx() < -1 {
					x++
					y = yDiag(y)
				} else if dy() > 1 {
					y--
					x = xDiag(x)
				} else if dy() < -1 {
					y++
					x = xDiag(x)
				}
				tail.move(x, y)
			}
		}
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

func makeNode() (n *Node) {
	n = new(Node)
	n.history = make(map[[2]int]bool)
	n.visited()
	return n
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
func parse() (ops Input) {
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
		ops = append(ops, i)
	}
	return ops
}

// Solve
func solve(ops Input) (solution string) {
	// Initial bridge state
	bridge := makeBridge()
	for i, op := range ops {
		bridge.move(op)
		fmt.Printf("After %d ops\n", i+1)
		fmt.Printf("4: %d,%d\n", bridge.nodes[4].x, bridge.nodes[4].y)
		fmt.Printf("5: %d,%d\n", bridge.nodes[5].x, bridge.nodes[5].y)
	}
	return strconv.Itoa(len(bridge.nodes[len(bridge.nodes)-1].history))
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
