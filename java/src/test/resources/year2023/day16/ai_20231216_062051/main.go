
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	right = iota
	up
	left
	down
)

type Point struct {
	x, y int
}

func main() {
	layout, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var energizedTiles = make(map[Point]bool)

	currentDirection := right
	currentPosition := Point{0, 0}

	// Process the light beam
	for 0 <= currentPosition.y && currentPosition.y < len(layout) && 0 <= currentPosition.x && currentPosition.x < len(layout[0]) {
		elem := layout[currentPosition.y][currentPosition.x]
		energizedTiles[currentPosition] = true

		switch elem {
		case '.':
			// Continue in the same direction
			currentPosition = moveInDirection(currentPosition, currentDirection)
		case '/':
			// Reflect the beam 90 degrees depending on the mirror angle
			switch currentDirection {
			case right:
				currentDirection = up
			case up:
				currentDirection = right
			case left:
				currentDirection = down
			case down:
				currentDirection = left
			}
			currentPosition = moveInDirection(currentPosition, currentDirection)
		case '\\':
			// Reflect the beam 90 degrees depending on the mirror angle
			switch currentDirection {
			case right:
				currentDirection = down
			case up:
				currentDirection = left
			case left:
				currentDirection = up
			case down:
				currentDirection = right
			}
			currentPosition = moveInDirection(currentPosition, currentDirection)
		case '|', '-':
			// If it's the flat side of a splitter, the beam splits
			if (elem == '|' && (currentDirection == right || currentDirection == left)) || (elem == '-' && (currentDirection == up || currentDirection == down)) {
				splitBeam(currentPosition, currentDirection, layout, energizedTiles)
				break
			}
			// If it's the pointy end of a splitter, pass through
			currentPosition = moveInDirection(currentPosition, currentDirection)
		}
	}

	fmt.Println(len(energizedTiles))
}

func moveInDirection(p Point, direction int) Point {
	switch direction {
	case right:
		p.x++
	case up:
		p.y--
	case left:
		p.x--
	case down:
		p.y++
	}
	return p
}

func splitBeam(position Point, currentDirection int, layout [][]rune, energizedTiles map[Point]bool) {
	// The beam splits into two beams: one vertical and one horizontal
	var newDirections []int
	if currentDirection == right || currentDirection == left {
		newDirections = []int{up, down}
	} else {
		newDirections = []int{right, left}
	}

	for _, dir := range newDirections {
		pos := position
		for layout[pos.y][pos.x] != '#' && 0 <= pos.y && pos.y < len(layout) && 0 <= pos.x && pos.x < len(layout[0]) {
			energizedTiles[pos] = true
			pos = moveInDirection(pos, dir)
			if pos.x < 0 || pos.x >= len(layout[0]) || pos.y < 0 || pos.y >= len(layout) {
				break
			}

			elem := layout[pos.y][pos.x]
			if elem == '#' || elem == '|' || elem == '-' {
				break
			}
		}
	}
}

func readInput(filepath string) ([][]rune, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var layout [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		layout = append(layout, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return layout, nil
}
