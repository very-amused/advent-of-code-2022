package p1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
type Input []*Monkey

type Monkey struct {
	items     []int
	operation func(item *int)
	test      func(item int) bool

	// T/F throw destination indexes
	dest1 int
	dest2 int

	inspectCount int
}

func (m *Monkey) inspect(index int) {
	// Monkey inspects item
	m.operation(&m.items[index])
	// Relief that inspection didn't damage item
	m.items[index] /= 3
	// Increment inspect count
	m.inspectCount++
}

// Decide where to throw an item based on worry level
func (m *Monkey) findDest(index int) (dest int) {
	if m.test(m.items[index]) {
		return m.dest1
	} else {
		return m.dest2
	}
}

// #endregion

// Parse
func parseMonkey(s *bufio.Scanner) (monkey *Monkey) {
	monkey = new(Monkey)
	var body string
	getBody := func(s *bufio.Scanner) string {
		s.Scan()
		l := s.Text()
		return strings.Split(l, ": ")[1]
	}
	// Parse starting items
	{
		body = getBody(s)
		startingItems := strings.Split(body, ", ")
		for _, item := range startingItems {
			monkey.items = append(monkey.items, must(strconv.Atoi(item)))
		}
	}
	// Operation
	{
		body = getBody(s)
		op := strings.Split(strings.Split(body, "= ")[1], " ")[1:]
		var opValue int
		if op[1] != "old" {
			opValue = must(strconv.Atoi(op[1]))
		}
		switch op[0][0] {
		case '*':
			if op[1] == "old" {
				monkey.operation = func(item *int) {
					*item *= *item
				}
			} else {
				monkey.operation = func(item *int) {
					*item *= opValue
				}
			}
		case '+':
			if op[1] == "old" {
				monkey.operation = func(item *int) {
					*item += *item
				}
			} else {
				monkey.operation = func(item *int) {
					*item += opValue
				}
			}
		default:
			panic("Unknown operator: " + op[0])
		}
	}
	// Test
	{
		body = getBody(s)
		var divisor int
		must(fmt.Sscanf(body, "divisible by %d", &divisor))
		monkey.test = func(item int) bool {
			return item%divisor == 0
		}
	}
	// Possible destinations
	{
		getDest := func() int {
			body = getBody(s)
			parts := strings.Split(body, " ")
			return must(strconv.Atoi(parts[len(parts)-1]))
		}
		monkey.dest1 = getDest()
		monkey.dest2 = getDest()
	}
	return monkey
}

func parse() (input Input) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parse lines
	for scanner.Scan() {
		l := scanner.Text()
		if !strings.HasPrefix(l, "Monkey") {
			continue
		}
		monkey := parseMonkey(scanner)
		input = append(input, monkey)
	}
	return input
}

// Solve
func solve(input Input) (solution string) {
	const rounds = 20
	for r := 0; r < rounds; r++ {
		fmt.Println("Round", r)
		for _, monkey := range input {
			for range monkey.items {
				monkey.inspect(0)
				dest := monkey.findDest(0)
				// Throw element to end of items list
				input[dest].items = append(input[dest].items, monkey.items[0])
				monkey.items = append(monkey.items[:0], monkey.items[1:]...) // Remove from throwing monkey
			}
		}
	}
	// Sort monkeys by inspect count
	sort.Slice(input, func(i, j int) bool {
		return input[i].inspectCount < input[j].inspectCount
	})
	// Multiply 2 highest counts
	monkeyBusiness := input[len(input)-2].inspectCount * input[len(input)-1].inspectCount
	return strconv.Itoa(monkeyBusiness)
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
