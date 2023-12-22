
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A Point represents a coordinate on a plane
type Point struct {
	x, y int
}

// Parses hexadecimal color codes into the correct instructions
func parseInstructions(fileName string) ([]Point, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)
	current := Point{0, 0}

	for scanner.Scan() {
		line := scanner.Text()
		colorCode := line[len(line)-7:] // Extract the color code
		distanceStr := colorCode[:5]   // The first 5 characters are the distance
		directionStr := colorCode[5:]  // The last character is the direction

		// Convert the distance from hexadecimal to decimal
		distance, err := strconv.ParseInt(distanceStr, 16, 64)
		if err != nil {
			return nil, err
		}

		// Decode the direction
		direction, err := strconv.ParseInt(directionStr, 16, 64)
		if err != nil {
			return nil, err
		}

		points = append(points, current) // Store the starting point

		switch direction {
		case 0: // Right
			current.x += int(distance)
		case 1: // Down
			current.y += int(distance)
		case 2: // Left
			current.x -= int(distance)
		case 3: // Up
			current.y -= int(distance)
		default:
			return nil, fmt.Errorf("unknown direction: %v", direction)
		}

		points = append(points, current) // Store the ending point
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

// Calculate the area using Pick's Theorem: A = i + b/2 - 1
func calculateArea(points []Point) int64 {
	var boundaryPoints, interiorPoints int64

	for i := 1; i < len(points); i++ {
		dx := abs(points[i].x - points[i-1].x)
		dy := abs(points[i].y - points[i-1].y)

		boundaryPoints += gcd(dx, dy)

		if points[i].x != points[i-1].x {
			// Vertical movement
			interiorPoints += (dx - 1) * int64(min(points[i].y, points[i-1].y))
		}
	}

	// calculate the area using Pick's Theorem
	return interiorPoints + boundaryPoints/2 + 1
}

// Helper functions
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Main function
func main() {
	points, err := parseInstructions("input.txt")
	if err != nil {
		fmt.Println("Error parsing instructions:", err)
		return
	}

	area := calculateArea(points)
	fmt.Println(area)
}
