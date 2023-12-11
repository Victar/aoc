
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var directions = []Point{
	{0, 1},  // South
	{1, 0},  // East
	{0, -1}, // North
	{-1, 0}, // West
}

func main() {
	grid, start := readInput("input.txt")
	visited := make(map[Point]struct{})
	longest := followLoop(grid, start, &visited)
	fmt.Println(longest)
}

func followLoop(grid map[Point]byte, start Point, visited *map[Point]struct{}) int {
	maxDistance := 0
	var dfs func(p Point, distance int)
	dfs = func(p Point, distance int) {
		if _, seen := (*visited)[p]; seen {
			return
		}
		if distance > maxDistance {
			maxDistance = distance
		}
		(*visited)[p] = struct{}{}
		for _, dir := range directions {
			np := Point{p.x + dir.x, p.y + dir.y}
			if next, ok := grid[np]; ok && canConnect(grid[p], next, dir) {
				dfs(np, distance+1)
			}
		}
	}
	dfs(start, 0)
	return maxDistance
}

func canConnect(from, to byte, dir Point) bool {
	switch from {
	case '-', 'S', '7', 'F':
		return dir.x != 0 // Connects horizontally
	case '|', 'L', 'J':
		return dir.y != 0 // Connects vertically
	}
	switch to {
	case '-', 'S', '7', 'F':
		return dir.x != 0
	case '|', 'L', 'J':
		return dir.y != 0
	}
	return false
}

func readInput(filename string) (grid map[Point]byte, start Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid = make(map[Point]byte)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if char != '.' {
				point := Point{x, y}
				grid[point] = byte(char)
				if char == 'S' {
					start = point
				}
			}
		}
		y++
	}
	return grid, start
}
