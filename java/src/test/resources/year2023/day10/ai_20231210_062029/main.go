
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	up = iota
	right
	down
	left
)

var (
	dx = map[int]int{up: -1, right: 0, down: 1, left: 0}
	dy = map[int]int{up: 0, right: 1, down: 0, left: -1}
)

func nextDirection(grid [][]rune, direction, x, y int) (int, bool) {
	switch grid[x][y] {
	case 'S', '|', '-':
		return direction, true
	case '7':
		if direction == right {
			return up, true
		}
		if direction == down {
			return left, true
		}
	case 'J':
		if direction == left {
			return up, true
		}
		if direction == down {
			return right, true
		}
	case 'L':
		if direction == up {
			return right, true
		}
		if direction == left {
			return down, true
		}
	case 'F':
		if direction == up {
			return left, true
		}
		if direction == right {
			return down, true
		}
	}
	return 0, false
}

func longestPath(grid [][]rune) int {
	startX, startY := -1, -1
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 'S' {
				startX, startY = x, y
				break
			}
		}
		if startX != -1 {
			break
		}
	}

	maxDist := 0
	for d := up; d <= left; d++ {
		x, y := startX+dx[d], startY+dy[d]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] != '.' && grid[x][y] != 'S' {
			visited := map[int]bool{startX*len(grid[0]) + startY: true}
			dist := 0
			direction := d
			for {
				dist++
				visited[x*len(grid[0])+y] = true
				if nextDir, ok := nextDirection(grid, direction, x, y); ok {
					direction = nextDir
					x += dx[direction]
					y += dy[direction]
					if visited[x*len(grid[0])+y] {
						break
					}
				} else {
					break
				}
			}
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	return maxDist
}

func readFile(fileName string) (grid [][]rune, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid, scanner.Err()
}

func main() {
	grid, err := readFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input.txt: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(longestPath(grid))
}
