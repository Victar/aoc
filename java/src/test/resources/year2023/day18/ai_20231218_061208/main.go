package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Position represents coordinates x, y in a 2D grid
type Position struct {
	x, y int
}

// Bounds holds the minimum and maximum values of x and y
type Bounds struct {
	minX, maxX, minY, maxY int
}

// Size returns the width and height of the area within the bounds.
func (b Bounds) Size() (int, int) {
	return b.maxX - b.minX + 1, b.maxY - b.minY + 1
}

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day18/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// The lagoon's boundaries
	var bounds Bounds
	bounds.minX, bounds.minY = 0, 0
	bounds.maxX, bounds.maxY = 0, 0

	// The starting location at the center of the trench
	pos := Position{0, 0}

	// Keep track of visited Positions and update boundaries
	visited := make(map[Position]bool)
	visited[pos] = true

	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`([UDLR]) (\d+)`)
		matches := r.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			dir := match[1]
			distance, _ := strconv.Atoi(match[2])

			dx, dy := 0, 0
			switch dir {
			case "U":
				dy = -1
			case "D":
				dy = 1
			case "L":
				dx = -1
			case "R":
				dx = 1
			}

			// Move and update positions and bounds
			for i := 0; i < distance; i++ {
				pos.x += dx
				pos.y += dy

				// Mark position as visited and update boundaries
				visited[pos] = true
				if pos.x < bounds.minX {
					bounds.minX = pos.x
				}
				if pos.x > bounds.maxX {
					bounds.maxX = pos.x
				}
				if pos.y < bounds.minY {
					bounds.minY = pos.y
				}
				if pos.y > bounds.maxY {
					bounds.maxY = pos.y
				}
			}
		}
	}

	// Determine the area enclosed by the trench
	area := 0

	width, height := bounds.Size()
	for y := bounds.minY; y <= bounds.maxY; y++ {
		for x := bounds.minX; x <= bounds.maxX; x++ {
			if visited[Position{x, y}] {
				area++
			}
		}
	}

	// Subtract visited area (trench path) from total grid area to get the interior area
	interiorArea := (width * height) - area

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Given the interior is dug out to be one meter deep, the volume is equal to the interior area
	fmt.Println(interiorArea, bounds, len(visited))
}
