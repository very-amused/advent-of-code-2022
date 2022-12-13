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
type Dir struct {
	parent   *Dir
	children map[string]*Dir
	files    []uint64
}

func (d *Dir) size() (s uint64) {
	for _, file := range d.files {
		s += file
	}
	for _, dir := range d.children {
		s += dir.size()
	}
	return s
}

func (d *Dir) flatten() (children []*Dir) {
	for _, c := range d.children {
		children = append(children, c)
		children = append(children, c.flatten()...)
	}
	return children
}

func (d *Dir) cd(path string) (n *Dir) {
	if path == ".." {
		if d.parent != nil {
			return d.parent
		}
		return d // cd .. from root doesn't move
	}
	// If the directory has not previously been explored, create it
	if d.children[path] == nil {
		d.mkdir(path)
	}
	return d.children[path]
}

func (d *Dir) mkdir(name string) {
	if d.children == nil {
		d.children = make(map[string]*Dir)
	}
	d.children[name] = &Dir{
		parent:   d,
		children: make(map[string]*Dir)}
}

// #endregion

// Parse
func parse() (root *Dir) {
	// Open scanner to read input line by line
	scanner := bufio.NewScanner(must(os.Open(inputFile)))

	// Parse lines
	scanner.Scan()
	if l := scanner.Text(); l != "$ cd /" {
		panic("Input doesn't start with cd to root dir: " + l)
	}
	root = &Dir{children: make(map[string]*Dir)}
	wd := root
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		{
			var cdPath string
			if n, _ := fmt.Sscanf(l, "$ cd %s", &cdPath); n == 1 {
				wd = wd.cd(cdPath)
				continue
			}
		}
		{
			var filename string
			var size uint64
			if n, _ := fmt.Sscanf(l, "%d %s", &size, &filename); n == 2 {
				wd.files = append(wd.files, size)
				continue
			}
		}
		{
			var dirname string
			if n, _ := fmt.Sscanf(l, "dir %s", &dirname); n == 1 {
				wd.mkdir(dirname)
				continue
			}
		}
	}
	return root
}

// Solve
func solve(root *Dir) (solution string) {
	allDirs := root.flatten()
	c := uint64(0)
	for _, dir := range allDirs {
		if s := dir.size(); s <= 100000 {
			c += s
		}
	}
	return strconv.FormatUint(c, 10)
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
