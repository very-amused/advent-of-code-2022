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
type Input []*Instruction

type CPU struct {
	instructions []*Instruction
	Registers    map[string]int
}

func (c *CPU) cycle() (done bool) {
	op := c.instructions[0]
	op.n--
	if op.n == 0 {
		op.f(c)
		c.instructions = c.instructions[1:]
	}
	return len(c.instructions) == 0
}

type Instruction struct {
	n int          // # of cycles instruction lasts
	f func(c *CPU) //	Function run after n cycles have passed since adding the instruction
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
		if l == "noop" {
			noop := &Instruction{
				n: 1,
				f: func(c *CPU) {}}
			ops = append(ops, noop)
			continue
		}
		var x int
		if n, _ := fmt.Sscanf(l, "addx %d", &x); n == 1 {
			addxOp := &Instruction{
				n: 2,
				f: func(c *CPU) {
					c.Registers["X"] += x
				}}
			ops = append(ops, addxOp)
		}
	}
	return ops
}

// Solve
func solve(ops Input) (solution string) {
	cpu := CPU{
		instructions: ops,
		Registers:    make(map[string]int)}
	cpu.Registers["X"] = 1
	signalStrengths := 0
	for cycles := 1; true; cycles++ {
		if cycles%40 == 20 {
			signalStrengths += cpu.Registers["X"] * cycles
		}
		if done := cpu.cycle(); done {
			break
		}
	}

	return strconv.Itoa(signalStrengths)
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
