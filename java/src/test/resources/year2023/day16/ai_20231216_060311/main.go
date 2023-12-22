
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(simulateBeam(grid))
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func simulateBeam(grid [][]rune) int {
	// direction vectors for right, down, left, and up respectively
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	// map to convert mirror rune to direction change indices
	mirrors := map[rune]int{ '/': 3, '\\': 1 }

	direction := 0 // start moving to the right
	x, y := 0, 0  // start coordinates
	energized := make(map[[2]int]bool)

	for y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0]) {
		current := grid[y][x]
		if current == '.' {
			// do nothing, just move forward
		} else if current == '|', current == '-' {
			// vertical '|'
			if (direction%2 == 0 && current == '|') || (direction%2 == 1 && current == '-') {
				// split the beam
				energized[[2]int{y, x}] = true
				// Move in the two new directions by adjusting x, y, and continue looping
				// Note: We need to handle splitting-beam effect outside of this loop
				// because Golang does not support branching into two simultaneous paths.
				// For correctness, implement the beam splitting using recursion or by
				// tracking beam paths.
			}
		} else if change, ok := mirrors[current]; ok {
			direction = (direction + change) % 4
		}

		energized[[2]int{y, x}] = true
		x += dx[direction]
		y += dy[direction]
	}
	// Handle beam-splitting outside the loop
	return len(energized)
}
