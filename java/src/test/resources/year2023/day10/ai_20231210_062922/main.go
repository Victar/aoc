
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

type Pipe struct {
	connections map[Point]bool
}

func main() {
	grid, start := readInput("input.txt")
	visited := make(map[Point]int)
	visit(start, grid, visited, start, 0)

	maxDistance := 0
	for _, d := range visited {
		if d > maxDistance {
			maxDistance = d
		}
	}

	fmt.Println(maxDistance)
}

func readInput(filename string) (map[Point]Pipe, Point) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make(map[Point]Pipe)
	var start Point
	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, ch := range line {
			point := Point{x, y}
			if ch == 'S' {
				start = point
			}
			pipe := parsePipe(ch)
			if pipe != nil {
				grid[point] = *pipe
			}
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, start
}

func parsePipe(ch rune) *Pipe {
	switch ch {
	case 'S', '-', '|', 'L', 'J', '7', 'F':
		return &Pipe{connections: getConnections(ch)}
	default:
		return nil
	}
}

func getConnections(ch rune) map[Point]bool {
	connections := make(map[Point]bool)

	switch ch {
	case 'S', '-', 'F', '7':
		connections[Point{-1, 0}] = true
		connections[Point{1, 0}] = true
	case '|', 'L', 'J':
		connections[Point{0, -1}] = true
		connections[Point{0, 1}] = true
	}

	switch ch {
	case 'S', '|', 'L', 'F':
		connections[Point{0, -1}] = true
	case 'J', '7':
		connections[Point{0, 1}] = true
	}

	return connections
}

func visit(p Point, grid map[Point]Pipe, visited map[Point]int, prev Point, distance int) {
	if d, ok := visited[p]; ok && d <= distance {
		return
	}
	visited[p] = distance

	pipe, exists := grid[p]
	if !exists {
		return
	}

	for dir := range pipe.connections {
		next := Point{p.x + dir.x, p.y + dir.y}
		if next == prev {
			continue
		}
		if nextPipe, ok := grid[next]; ok {
			nextDir := Point{-dir.x, -dir.y}
			if _, ok := nextPipe.connections[nextDir]; ok {
				visit(next, grid, visited, p, distance+1)
			}
		}
	}
}
