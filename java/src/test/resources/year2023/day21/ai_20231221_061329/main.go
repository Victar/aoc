
package main

import (
	"bufio"
	"fmt"
	"os"
)

const steps = 64

// Point represents a location in 2D space.
type Point struct {
	X, Y int
}

// InBounds checks if the given point is within the grid boundary.
func InBounds(p Point, grid [][]byte) bool {
	return p.Y >= 0 && p.Y < len(grid) && p.X >= 0 && p.X < len(grid[p.Y])
}

// Step in four possible directions: up, down, left, right.
var directions = []Point{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]byte, 0)
	var start Point

	// Read the input file into a grid and find the starting point.
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []byte(line)
		grid = append(grid, row)
		for x, char := range row {
			if char == 'S' {
				start = Point{X: x, Y: y}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Using dynamic programming to count reachable points.
	dp := make([]map[Point]int, steps+1)
	for i := range dp {
		dp[i] = make(map[Point]int)
	}
	dp[0][start] = 1

	for s := 1; s <= steps; s++ {
		for p := range dp[s-1] {
			for _, d := range directions {
				next := Point{X: p.X + d.X, Y: p.Y + d.Y}
				if InBounds(next, grid) && grid[next.Y][next.X] == '.' {
					dp[s][next] += dp[s-1][p]
				}
			}
		}
	}

	// Count distinct garden plots that can be reached in exactly 'steps' steps.
	var count int
	for p := range dp[steps] {
		if grid[p.Y][p.X] == '.' {
			count++
		}
	}

	fmt.Println(count)
}
