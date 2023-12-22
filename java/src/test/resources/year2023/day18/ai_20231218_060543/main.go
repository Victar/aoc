package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	// Open the input file
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day18/sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Initialize the area using a map to represent the dug trenches
	area := make(map[Point]bool)
	scanner := bufio.NewScanner(file)

	// Set the starting point
	var current Point

	// Process each line from the input
	for scanner.Scan() {
		// Parse the instructions
		instructions := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(instructions); i += 3 {
			direction := instructions[i]
			steps, _ := strconv.Atoi(instructions[i+1])

			for j := 0; j < steps; j++ {
				// Mark the current position as dug
				area[current] = true

				// Move to the next position based on the direction
				switch direction {
				case "U":
					current.y--
				case "D":
					current.y++
				case "L":
					current.x--
				case "R":
					current.x++
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Compute bounds of the dug area
	var minX, maxX, minY, maxY int
	for point := range area {
		if point.x < minX {
			minX = point.x
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	// Iterate through all points within the bounding box
	// and determine if they are inside the lagoon using the ray casting algorithm.
	volume := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			point := Point{x, y}
			if isInside(point, area) {
				volume++
			}
		}
	}

	fmt.Println(volume, len(area))
}

// Function to determine if a point is inside the lagoon using ray casting algorithm.
func isInside(point Point, area map[Point]bool) bool {
	rayEndPoint := Point{point.x + 1, point.y}
	count := 0
	for rayEndPoint.x <= len(area) {
		rayEndPoint.x++
		if area[rayEndPoint] {
			count++
		}
	}
	// If the ray intersects odd number of times, the point is inside
	return count%2 != 0
}
