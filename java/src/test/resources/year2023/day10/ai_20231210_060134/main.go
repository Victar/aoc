
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Directions as dx, dy arrays for N, E, S, W respectively
var dirMap = map[rune][2]int{
	'N': {0, -1},
	'E': {1, 0},
	'S': {0, 1},
	'W': {-1, 0},
}

var pipeConnections = map[rune]string{
	'|': "NS",
	'-': "WE",
	'7': "SW",
	'L': "NE",
	'J': "NW",
	'F': "SE",
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pipes [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines := []rune(scanner.Text())
		pipes = append(pipes, lines)
	}

	return pipes, scanner.Err()
}

func findStart(pipes [][]rune) (int, int) {
	for y, row := range pipes {
		for x, tile := range row {
			if tile == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

// Check if the next pipe can be connected from direction d
func canConnectFrom(pipe rune, d rune) bool {
	if directions, ok := pipeConnections[pipe]; ok {
		for _, direction := range directions {
			if direction == d {
				return true
			}
		}
	}
	return false
}

func oppositeDirection(d rune) rune {
	switch d {
	case 'N':
		return 'S'
	case 'S':
		return 'N'
	case 'W':
		return 'E'
	case 'E':
		return 'W'
	}
	return ' '
}

func findFarthestPoint(pipes [][]rune, startX, startY int) int {
	maxDistance := 0
	visited := make(map[[2]int]bool)
	var search func(x, y int, distance int, comingFrom rune)
	search = func(x, y int, distance int, comingFrom rune) {
		if visited[[2]int{x, y}] {
			// We looped back or hit a visited pipe.
			return
		}
		visited[[2]int{x, y}] = true
		if distance > maxDistance {
			maxDistance = distance
		}
		for direction, diff := range dirMap {
			nx, ny := x+diff[0], y+diff[1]
			if ny >= 0 && ny < len(pipes) && nx >= 0 && nx < len(pipes[ny]) {
				nextPipe := pipes[ny][nx]
				// Check for connection from the opposite direction
				if canConnectFrom(nextPipe, oppositeDirection(direction)) && direction != comingFrom {
					search(nx, ny, distance+1, direction)
				}
			}
		}
	}
	search(startX, startY, 0, ' ')
	return maxDistance
}

func main() {
	pipes, err := readInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	startX, startY := findStart(pipes)
	if startX == -1 || startY == -1 {
		fmt.Println("Error: Starting position 'S' not found in the input")
		return
	}
	farthest := findFarthestPoint(pipes, startX, startY)
	fmt.Println(farthest) // Print the answer only
}
