
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Directions maps to coordinate changes for N, E, S, W respectively
var Directions = []struct{ dx, dy int }{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

// ReadInput reads the input from "input.txt" and returns a slice of strings
func ReadInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// isValidPipe tells if the pipe can connect in the given direction
func isValidPipe(pipe rune, direction int) bool {
	switch pipe {
	case '|':
		return direction == 0 || direction == 2 // N or S
	case '-':
		return direction == 1 || direction == 3 // E or W
	case 'L':
		return direction == 0 || direction == 1 // N or E
	case 'J':
		return direction == 0 || direction == 3 // N or W
	case '7':
		return direction == 2 || direction == 3 // S or W
	case 'F':
		return direction == 1 || direction == 2 // E or S
	case 'S':
		return true // S can be any valid direction
	}
	return false
}

// dfs uses Depth First Search to explore the maze and return the maximum steps
func dfs(maze []string, seen [][]bool, startX, startY, prevDirection, steps int) int {
	if seen[startY][startX] {
		return steps - 1 // Subtract 1 to ignore counting the start position twice
	}
	seen[startY][startX] = true

	maxSteps := steps
	for i, dir := range Directions {
		nextX, nextY := startX+dir.dx, startY+dir.dy
		if nextY >= 0 && nextY < len(maze) && nextX >= 0 && nextX < len(maze[nextY]) {
			nextPipe := rune(maze[nextY][nextX])
			if isValidPipe(nextPipe, i) && ((prevDirection+2)%4 != i || rune(maze[startY][startX]) == 'S') {
				result := dfs(maze, seen, nextX, nextY, i, steps+1)
				if result > maxSteps {
					maxSteps = result
				}
			}
		}
	}
	return maxSteps
}

func main() {
	maze, err := ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var startX, startY int
	found := false
	for y, row := range maze {
		for x, tile := range row {
			if tile == 'S' {
				startX, startY = x, y
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[i]))
	}

	maxSteps := dfs(maze, seen, startX, startY, -1, 0)
	fmt.Println(maxSteps)
}
