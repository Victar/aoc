
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct{ x, y int }

var directions = map[rune][]Point{
	'F': {{x: 1, y: 0}, {x: 0, y: 1}},
	'7': {{x: -1, y: 0}, {x: 0, y: 1}},
	'L': {{x: 0, y: -1}, {x: 1, y: 0}},
	'J': {{x: 0, y: -1}, {x: -1, y: 0}},

	'|': {{x: 0, y: -1}, {x: 0, y: 1}},
	'-': {{x: -1, y: 0}, {x: 1, y: 0}},
}

func readInput(filename string) ([][]rune, Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]rune
	var start Point

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		for x, r := range row {
			if r == 'S' {
				start = Point{x: x, y: y}
			}
		}
		matrix = append(matrix, row)
		y++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return matrix, start
}

func findLoop(matrix [][]rune, start Point) int {
	visited := make(map[Point]bool)
	queue := []Point{start}
	prev := make(map[Point]Point)
	distances := make(map[Point]int)
	maxDistance := 0

	for len(queue) > 0 {
		current := queue[0]
		visited[current] = true
		queue = queue[1:]

		currentChar := matrix[current.y][current.x]
		if currentChar == '.' || currentChar == 'S' {
			currentChar = 'F' // Assuming 'S' is a 90-degree F bend by default since it's not specified
		}

		if options, ok := directions[currentChar]; ok {
			for _, dir := range options {
				next := Point{x: current.x + dir.x, y: current.y + dir.y}
				if next.x < 0 || next.y < 0 || next.y >= len(matrix) || next.x >= len(matrix[next.y]) {
					continue
				}
				nextChar := matrix[next.y][next.x]
				if nextChar != '.' && !visited[next] {
					prev[next] = current
					distances[next] = distances[current] + 1
					if distances[next] > maxDistance {
						maxDistance = distances[next]
					}
					queue = append(queue, next)
				}
			}
		}
	}

	return maxDistance
}

func resolveMaze(maze [][]rune, start Point) {
	farthest := findLoop(maze, start)
	fmt.Println(farthest)
}

func main() {
	maze, startPos := readInput("input.txt")
	resolveMaze(maze, startPos)
}
