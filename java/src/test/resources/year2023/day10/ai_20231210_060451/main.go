
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	grid, start := readInput("input.txt")

	// We start with the initial point 'S' and perform a DFS traversing all connected pipes to find the loop.
	visited := make(map[Point]bool)
	var path []Point
	dfs(start, grid, visited, &path)

	// We identify the farthest point from the start by comparing visited points in the path.
	maxDistance := 0
	for p := range visited {
		distance := len(findPathTo(start, p, path))
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	fmt.Println(maxDistance)
}

// readInput reads the grid from the input file and locates the start point 'S'.
func readInput(filename string) (map[Point]rune, Point) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer file.Close()

	grid := make(map[Point]rune)
	var start Point
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			grid[Point{x, y}] = char
			if char == 'S' {
				start = Point{x, y}
			}
		}
		y++	
	}
	return grid, start
}

// dfs performs a Depth-First Search from the current point to find all points in the loop.
func dfs(p Point, grid map[Point]rune, visited map[Point]bool, path *[]Point) {
	// If already visited, return.
	if visited[p] {
		return
	}
	visited[p] = true
	*path = append(*path, p)

	// Check all four neighboring points for possible connected pipes.
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// Skip diagonals and the current point
			if dx != 0 && dy != 0 || (dx == 0 && dy == 0) {
				continue
			}

			next := Point{p.X + dx, p.Y + dy}
			if isConnected(grid[p], grid[next], dx, dy) {
				dfs(next, grid, visited, path)
			}
		}
	}
}

// isConnected checks if two pipes at points p1 and p2 are connected, given the offset from p1 to p2.
func isConnected(p1, p2 rune, dx, dy int) bool {
	switch {
	case dx == 1: // to the right
		return (p1 == '-' || p1 == 'J' || p1 == '7' || p1 == 'S') && (p2 == '-' || p2 == 'L' || p2 == 'F' || p2 == 'S')
	case dx == -1: // to the left
		return (p1 == '-' || p1 == 'L' || p1 == 'F' || p1 == 'S') && (p2 == '-' || p2 == 'J' || p2 == '7' || p2 == 'S')
	case dy == 1: // downwards
		return (p1 == '|' || p1 == 'L' || p1 == 'J' || p1 == 'S') && (p2 == '|' || p2 == 'F' || p2 == '7' || p2 == 'S')
	case dy == -1: // upwards
		return (p1 == '|' || p1 == 'F' || p1 == '7' || p1 == 'S') && (p2 == '|' || p2 == 'L' || p2 == 'J' || p2 == 'S')
	}
	return false
}

// findPathTo finds a path between two points in the loop, counting the steps taken.
func findPathTo(start, end Point, path []Point) []Point {
	// Create a map from points to their path indexes for quick lookup.
	indexMap := make(map[Point]int, len(path))
	for i, p := range path {
		indexMap[p] = i
	}

	startIndex, startFound := indexMap[start]
	endIndex, endFound := indexMap[end]
	if !startFound || !endFound {
		return nil // One of the points wasn't in the path.
	}

	// Build a slice of points from start to end.
	var result []Point
	// Determine path direction for the shortest path on the loop.
	if startIndex < endIndex || endIndex == 0 {
		result = append(result, path[startIndex:endIndex+1]...)
	} else {
		result = append(result, path[startIndex:]...)
		result = append(result, path[:endIndex+1]...)
	}
	return result
}
