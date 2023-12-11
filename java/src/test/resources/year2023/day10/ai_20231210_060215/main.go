
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var (
	deltas = map[rune][]Point{
		'F': {{-1, 0}, {0, 1}},
		'7': {{1, 0}, {0, -1}},
		'L': {{-1, 0}, {0, -1}},
		'J': {{1, 0}, {0, 1}},
		'|': {{-1, 0}, {1, 0}},
		'-': {{0, -1}, {0, 1}},
	}
)

func main() {
	path, startPos := readInput("input.txt")

	// Find the loop and get the farthest distance
	distance := findFarthestDistance(path, startPos)
	fmt.Println(distance)
}

func readInput(fileName string) (map[Point]rune, Point) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	path := make(map[Point]rune)
	var startPos Point
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if char != '.' {
				path[Point{x, y}] = char
			}
			if char == 'S' {
				startPos = Point{x, y}
			}
		}
		y++
	}
	return path, startPos
}

func findFarthestDistance(path map[Point]rune, startPos Point) int {
	visited := make(map[Point]bool)
	maxDistance := 0
	var backtrack func(p Point, distance int)
	backtrack = func(p Point, distance int) {
		if visited[p] {
			return
		}
		visited[p] = true

		if distance > maxDistance {
			maxDistance = distance
		}

		tile, found := path[p]
		if !found {
			return
		}
		for _, delta := range deltas[tile] {
			nextP := Point{p.x + delta.x, p.y + delta.y}
			_, foundNext := path[nextP]
			if foundNext {
				backtrack(nextP, distance+1)
			}
		}
	}

	backtrack(startPos, 0)
	return maxDistance
}
