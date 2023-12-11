
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	startRune      = 'S'
	horizontalRune = '-'
	verticalRune   = '|'
	groundRune     = '.'
)

var (
	directions = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

type Point struct {
	x int
	y int
}

func main() {
	grid, startX, startY, err := readInput("input.txt")

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	dist := bfs(grid, Point{x: startX, y: startY})
	fmt.Println(dist)
}

func bfs(grid [][]rune, start Point) int {
	// Initialize visited grid with false values
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	// Create queue and add starting position
	queue := []Point{start}
	visited[start.y][start.x] = true

	// Initialize variables for BFS loop
	var maxDist int
	dist := map[Point]int{start: 0}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:] // Dequeue

		// Update max distance
		if dist[point] > maxDist {
			maxDist = dist[point]
		}

		// Process neighbors
		for _, dir := range directions {
			nextPoint := Point{point.x + dir.x, point.y + dir.y}
			if isAccessible(grid, point, nextPoint) && !visited[nextPoint.y][nextPoint.x] {
				queue = append(queue, nextPoint)
				visited[nextPoint.y][nextPoint.x] = true
				dist[nextPoint] = dist[point] + 1
			}
		}
	}

	return maxDist
}

func isAccessible(grid [][]rune, from, to Point) bool {
	if to.y < 0 || to.y >= len(grid) || to.x < 0 || to.x >= len(grid[to.y]) || grid[to.y][to.x] == groundRune {
		return false
	}
	fromRune, toRune := grid[from.y][from.x], grid[to.y][to.x]

	// Define rules for checking if a move is accessible based on the pipe characters
	dx, dy := to.x-from.x, to.y-from.y
	switch {
	case dx == 1 && (fromRune == horizontalRune || fromRune == startRune) && (toRune == horizontalRune || "L7F".contains(string(toRune))):
		return true
	case dx == -1 && (fromRune == horizontalRune || fromRune == startRune) && (toRune == horizontalRune || "JLF".contains(string(toRune))):
		return true
	case dy == 1 && (fromRune == verticalRune || fromRune == startRune) && (toRune == verticalRune || "FJ7".contains(string(toRune))):
		return true
	case dy == -1 && (fromRune == verticalRune || fromRune == startRune) && (toRune == verticalRune || "LJ7".contains(string(toRune))):
		return true
	default:
		return false
	}
}

func readInput(filename string) ([][]rune, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	var startX, startY int

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
		for x, cell := range row {
			if cell == startRune {
				startX = x
				startY = y
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	return grid, startX, startY, nil
}

func (s string) contains(substr string) bool {
	for _, c := range substr {
		if !containsRune(s, c) {
			return false
		}
	}
	return true
}

func containsRune(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}
