
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dug := make(map[Point]bool)
	x, y := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		instructions := strings.Split(line, " ")

		for i := 0; i < len(instructions); i += 3 {
			direction := instructions[i]
			length, _ := strconv.Atoi(strings.TrimSpace(instructions[i+1]))

			for j := 0; j < length; j++ {
				switch direction {
				case "U":
					y--
				case "D":
					y++
				case "L":
					x--
				case "R":
					x++
				}

				// Mark the block as dug.
				dug[Point{x, y}] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate the size of the lagoon.
	minX, maxX, minY, maxY := calculateBounds(dug)
	lagoonSize := 0
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if _, ok := dug[Point{i, j}]; ok {
				lagoonSize++
			} else {
				// Check if this block is inside the lagoon by seeing if it's surrounded by dug blocks.
				if isInsideLagoon(dug, i, j, minX, maxX, minY, maxY) {
					lagoonSize++
				}
			}
		}
	}

	fmt.Println(lagoonSize)
}

// calculateBounds finds the minimum and maximum x and y coordinates for the trench.
func calculateBounds(dug map[Point]bool) (minX, maxX, minY, maxY int) {
	first := true
	for p := range dug {
		if first {
			minX, maxX, minY, maxY = p.X, p.X, p.Y, p.Y
			first = false
		} else {
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
	}
	return
}

// isInsideLagoon checks if a block is inside the lagoon defined by the dug map.
func isInsideLagoon(dug map[Point]bool, x, y, minX, maxX, minY, maxY int) bool {
	if x == minX || x == maxX || y == minY || y == maxY {
		return false // On the boundary.
	}
	for i := x; i >= minX; i-- {
		if !dug[Point{i, y}] {
			return false
		}
	}
	for i := x; i <= maxX; i++ {
		if !dug[Point{i, y}] {
			return false
		}
	}
	for j := y; j >= minY; j-- {
		if !dug[Point{x, j}] {
			return false
		}
	}
	for j := y; j <= maxY; j++ {
		if !dug[Point{x, j}] {
			return false
		}
	}
	return true
}
