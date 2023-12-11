
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	North = iota
	East
	South
	West
)

var (
	dx = []int{0, 1, 0, -1}
	dy = []int{-1, 0, 1, 0}
)

type Pipe struct {
	x, y  int
	pipe  rune
	paths []bool
}

func readInput(filename string) ([][]Pipe, Pipe) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]Pipe
	var start Pipe
	y := 0
	for scanner.Scan() {
		var line []Pipe
		text := scanner.Text()
		for x, char := range text {
			pipe := Pipe{x: x, y: y, pipe: char, paths: make([]bool, 4)}
			switch char {
			case 'S':
				start = pipe
			case '|':
				pipe.paths[North], pipe.paths[South] = true, true
			case '-':
				pipe.paths[East], pipe.paths[West] = true, true
			case 'L':
				pipe.paths[North], pipe.paths[East] = true, true
			case 'J':
				pipe.paths[North], pipe.paths[West] = true, true
			case '7':
				pipe.paths[South], pipe.paths[West] = true, true
			case 'F':
				pipe.paths[South], pipe.paths[East] = true, true
			}
			line = append(line, pipe)
		}
		grid = append(grid, line)
		y++
	}

	return grid, start
}

func validMove(grid [][]Pipe, x, y, direction int) bool {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return false
	}
	return grid[y][x].paths[direction]
}

func backtrack(grid [][]Pipe, current Pipe, prevDir int, steps *int) {
	maxSteps := *steps
	for _, dir := range []int{North, East, South, West} {
		if dir == (prevDir+2)%4 || !current.paths[dir] {
			continue
		}
		xx, yy := current.x+dx[dir], current.y+dy[dir]
		if validMove(grid, xx, yy, dir) {
			nextPipe := grid[yy][xx]
			if nextPipe.paths[(dir+2)%4] {
				*steps = maxSteps + 1
				backtrack(grid, nextPipe, dir, steps)
			}
		}
	}
}

func main() {
	grid, start := readInput("input.txt")
	steps := 0
	backtrack(grid, start, -1, &steps)
	fmt.Println(steps)
}
