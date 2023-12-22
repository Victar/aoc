
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
}

type Brick struct {
	Start, End Point
}

// ParseBrick parses a single brick description
func ParseBrick(s string) (Brick, error) {
	points := strings.Split(s, "~")
	if len(points) != 2 {
		return Brick{}, fmt.Errorf("invalid input: %s", s)
	}

	startStr := strings.Split(points[0], ",")
	endStr := strings.Split(points[1], ",")

	if len(startStr) != 3 || len(endStr) != 3 {
		return Brick{}, fmt.Errorf("invalid coordinates: %s", s)
	}

	start, err := parsePoint(startStr)
	if err != nil {
		return Brick{}, err
	}

	end, err := parsePoint(endStr)
	if err != nil {
		return Brick{}, err
	}

	return Brick{Start: start, End: end}, nil
}

// parsePoint converts string array to a Point
func parsePoint(pointStr []string) (Point, error) {
	x, err := strconv.Atoi(pointStr[0])
	if err != nil {
		return Point{}, fmt.Errorf("invalid x coordinate: %s", pointStr[0])
	}

	y, err := strconv.Atoi(pointStr[1])
	if err != nil {
		return Point{}, fmt.Errorf("invalid y coordinate: %s", pointStr[1])
	}

	z, err := strconv.Atoi(pointStr[2])
	if err != nil {
		return Point{}, fmt.Errorf("invalid z coordinate: %s", pointStr[2])
	}

	return Point{X: x, Y: y, Z: z}, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bricks []Brick
	for scanner.Scan() {
		line := scanner.Text()
		brick, err := ParseBrick(line)
		if err != nil {
			fmt.Println("Error parsing brick:", err)
			return
		}
		bricks = append(bricks, brick)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Now you have a slice of bricks `bricks`
	// Implement the logic to calculate the correct answer
	
	// Stub for answer calculation
	safeDisintegrateCount := calculateSafeDisintegration(bricks)
	fmt.Println(safeDisintegrateCount)
}

// TODO: Implement this function
func calculateSafeDisintegration(bricks []Brick) int {
	// Implement the logic to identify which bricks can be safely disintegrated

	// This is a placeholder return
	return 0
}
