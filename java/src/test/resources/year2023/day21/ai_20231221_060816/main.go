
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	empty    = '.'
	rock     = '#'
	start    = 'S'
	steps    = 64
	filename = "input.txt"
)

type Point struct {
	X, Y int
}

var directions = []Point{
	{X: 1, Y: 0},  // Right
	{X: -1, Y: 0}, // Left
	{X: 0, Y: 1},  // Down
	{X: 0, Y: -1}, // Up
}

func loadDataFromFile(filepath string) ([]string, Point, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, Point{}, err
	}
	defer file.Close()

	var data []string
	startPos := Point{}
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		data = append(data, line)
		if x := findStart(line); x != -1 {
			startPos = Point{X: x, Y: y}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, Point{}, err
	}

	return data, startPos, nil
}

func findStart(line string) int {
	for i, c := range line {
		if c == start {
			return i
		}
	}
	return -1
}

func validStep(grid []string, x, y int) bool {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return false
	}
	return grid[y][x] == empty
}

func bfs(grid []string, startPos Point) map[Point]bool {
	queue := []Point{startPos}
	visited := map[Point]bool{startPos: true}
	reached := make(map[Point]bool)

	for len(queue) > 0 {
		currentSize := len(queue)
		for s := 0; s < currentSize; s++ {
			point := queue[0]
			queue = queue[1:]

			if _, found := reached[point]; !found && len(visited)-1 == steps { // -1 to exclude start
				reached[point] = true
			}

			for _, dir := range directions {
				nextX, nextY := point.X+dir.X, point.Y+dir.Y
				nextPoint := Point{X: nextX, Y: nextY}
				if validStep(grid, nextX, nextY) && !visited[nextPoint] {
					visited[nextPoint] = true
					queue = append(queue, nextPoint)
				}
			}
		}

		if len(visited)-1 == steps { // Early stop after the required steps
			break
		}
	}

	return reached
}

func main() {
	grid, startPos, err := loadDataFromFile(filename)
	if err != nil {
		panic(err)
	}

	reached := bfs(grid, startPos)
	fmt.Println(len(reached))
}
