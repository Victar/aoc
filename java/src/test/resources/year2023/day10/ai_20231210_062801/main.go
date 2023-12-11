
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
	grid := readInput("input.txt")

	// Find the start position
	var start Point
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				start = Point{X: x, Y: y}
				break
			}
		}
		if start != (Point{}) {
			break
		}
	}

	// Find all tiles in the loop through Depth-First Search (DFS)
	visited := make(map[Point]bool)
	visited[start] = true
	dfs(grid, start, visited)

	// Find the tile in the loop that is farthest from the starting position
	maxDistance := 0
	for point := range visited {
		distance := measureDistance(grid, visited, start, point)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Printf("%d\n", maxDistance)
}

func dfs(grid [][]rune, current Point, visited map[Point]bool) {
	directions := []Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	for _, d := range directions {
		next := Point{X: current.X + d.X, Y: current.Y + d.Y}
		if next.X < 0 || next.Y < 0 || next.Y >= len(grid) || next.X >= len(grid[next.Y]) || visited[next] || grid[next.Y][next.X] == '.' {
			continue
		}
		// Check if next is connected to current
		if isConnected(grid[next.Y][next.X], d) {
			visited[next] = true
			dfs(grid, next, visited)
		}
	}
}

func isConnected(r rune, direction Point) bool {
	switch r {
	case '|':
		return direction.Y != 0
	case '-':
		return direction.X != 0
	case 'L':
		return direction.Y != -1 && direction.X != 1
	case 'J':
		return direction.Y != -1 && direction.X != -1
	case '7':
		return direction.Y != 1 && direction.X != -1
	case 'F':
		return direction.Y != 1 && direction.X != 1
	case 'S':
		return true
	}
	return false
}

func measureDistance(grid [][]rune, visited map[Point]bool, start, end Point) int {
	queue := []Point{start}
	distance := make(map[Point]int)
	distance[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			return distance[current]
		}

		directions := []Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
		for _, d := range directions {
			next := Point{X: current.X + d.X, Y: current.Y + d.Y}
			if _, ok := distance[next]; !ok && visited[next] {
				distance[next] = distance[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return 0
}

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %s", err)
	}
	return grid
}
