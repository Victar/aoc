
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X, Y int
}

var directions = map[rune]Point{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Use a regular expression to parse each line.
	re := regexp.MustCompile(`([UDLR]) (\d+)`)
	scanner := bufio.NewScanner(file)
	grid := make(map[Point]bool)

	x, y := 0, 0
	// Parse the input and fill in the grid with trenches.
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			dir := match[1][0]
			steps, _ := strconv.Atoi(match[2])
			dx, dy := directions[rune(dir)].X, directions[rune(dir)].Y
			// Move in the specified direction and add points to the grid.
			for s := 0; s < steps; s++ {
				x += dx
				y += dy
				grid[Point{x, y}] = true
			}
		}
	}

	// Calculate the volume of lava the lagoon can hold.
	lagoonVolume := calculateVolume(grid)
	fmt.Println(lagoonVolume)
}

func calculateVolume(grid map[Point]bool) int {
	// Initiate a variable to hold the volume of lava the lagoon can hold.
	volume := 0

	// We need to find the minimum and maximum x and y to bound the lagoon for calculations.
	minX, minY := 0, 0
	maxX, maxY := 0, 0
	for p := range grid {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// Check all points within the bounds and count those that are part of the lagoon.
	for y := minY + 1; y < maxY; y++ {
		for x := minX + 1; x < maxX; x++ {
			if grid[Point{x - 1, y}] && grid[Point{x + 1, y}] && grid[Point{x, y - 1}] && grid[Point{x, y + 1}] {
				volume++
			}
		}
	}

	return volume
}
