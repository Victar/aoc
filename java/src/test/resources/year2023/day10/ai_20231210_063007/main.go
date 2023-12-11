
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	maze := strings.Split(strings.TrimSpace(string(input)), "\n")
	answer := findFarthestPoint(maze)
	fmt.Println(answer)
}

func findFarthestPoint(maze []string) int {
	var start Coord
	pipeMap := make(map[Coord]rune)
	directions := []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	connects := map[rune][]int{
		'|': {0, 2},
		'-': {1, 3},
		'L': {0, 1},
		'J': {0, 3},
		'7': {2, 3},
		'F': {1, 2},
	}

	for y, line := range maze {
		for x, char := range line {
			coord := Coord{x, y}
			if char == '.' {
				continue
			}

			pipeMap[coord] = char

			if char == 'S' {
				start = coord
				pipeMap[coord] = findConnection(maze, coord)
			}
		}
	}

	maxDist := 0
	queue := []CoordDist{{start, 0}}
	visited := make(map[Coord]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.Coord] {
			continue
		}

		visited[current.Coord] = true

		if current.Dist > maxDist {
			maxDist = current.Dist
		}

		for _, dirIdx := range connects[pipeMap[current.Coord]] {
			dir := directions[dirIdx]
			nextCoord := Coord{current.X + dir.X, current.Y + dir.Y}
			if _, ok := pipeMap[nextCoord]; ok && !visited[nextCoord] {
				queue = append(queue, CoordDist{nextCoord, current.Dist + 1})
			}
		}
	}

	return maxDist
}

type Coord struct {
	X, Y int
}

type CoordDist struct {
	Coord
	Dist int
}

func findConnection(maze []string, start Coord) rune {
	directions := []struct {
		Dir   rune
		Delta Coord
	}{
		{'|', {0, -1}}, {'|', {0, 1}},
		{'-', {-1, 0}}, {'-', {1, 0}},
	}
	var connections int
	for _, dir := range directions {
		candidate := Coord{start.X + dir.Delta.X, start.Y + dir.Delta.Y}
		if candidate.Y < 0 || candidate.Y >= len(maze) || candidate.X < 0 || candidate.X >= len(maze[candidate.Y]) {
			continue
		}

		if isPipe(maze[candidate.Y][candidate.X]) {
			connections++
		}
	}

	switch connections {
	case 2:
		return 'S' // S is assumed to act like a pipe piece with exactly two connections.
	default:
		panic("invalid connections at start")
	}
}

func isPipe(ch rune) bool {
	return ch == '|' || ch == '-' || ch == 'L' || ch == 'J' || ch == '7' || ch == 'F' || ch == 'S'
}
