
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Right = 'R'
	Up    = 'U'
	Left  = 'L'
	Down  = 'D'
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	var (
		minX, minY, maxX, maxY int
		x, y                   int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, " ")

		for i := 0; i < len(commands); i += 2 {
			direction := commands[i][0]
			length, _ := strconv.Atoi(commands[i][1:])
			switch direction {
			case Right:
				x += length
				if x > maxX {
					maxX = x
				}
			case Left:
				x -= length
				if x < minX {
					minX = x
				}
			case Up:
				y -= length
				if y < minY {
					minY = y
				}
			case Down:
				y += length
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// Create the grid with the proper size.
	gridWidth := maxX - minX + 1
	gridHeight := maxY - minY + 1
	grid := make([][]bool, gridHeight)
	for i := range grid {
		grid[i] = make([]bool, gridWidth)
	}

	// Redo the path to mark the grid.
	x, y = -minX, -minY // Adjusting the original starting point (0,0) with new grid bounds
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning input file: %v", err)
	}

	file.Seek(0, 0) // Go back to the beginning of the file
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, " ")

		for i := 0; i < len(commands); i += 2 {
			direction := commands[i][0]
			length, _ := strconv.Atoi(commands[i][1:])
			for l := 0; l < length; l++ {
				grid[y][x] = true // Mark the grid cell
				switch direction {
				case Right:
					x++
				case Left:
					x--
				case Up:
					y--
				case Down:
					y++
				}
			}
		}
	}

	// Calculate the area.
	area := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				area++
			}
		}
	}
	// The final answer should be the total area minus the perimeter.
	perimeter := 2*(gridWidth+gridHeight) - 4
	result := area - perimeter
	println(result)
}
