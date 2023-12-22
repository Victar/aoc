
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	points := []Point{{0, 0}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := parseHexAndDirection(line)
		if parsed.steps == 0 {
			continue
		}
		last := points[len(points)-1]
		var next Point
		switch parsed.direction {
		case "R":
			next = Point{last.x + parsed.steps, last.y}
		case "D":
			next = Point{last.x, last.y + parsed.steps}
		case "L":
			next = Point{last.x - parsed.steps, last.y}
		case "U":
			next = Point{last.x, last.y - parsed.steps}
		}
		points = append(points, next)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	area := calculateLavaArea(points)
	fmt.Println(area)
}

func parseHexAndDirection(input string) (parsed struct{ steps int64; direction string }) {
	colorCode := input[len(input)-9 : len(input)-2]
	lengthHex, directionHex := colorCode[:5], colorCode[5:6]
	length, _ := strconv.ParseInt(lengthHex, 16, 64)
	directionNum, _ := strconv.ParseInt(directionHex, 16, 64)
	directions := map[int64]string{
		0: "R",
		1: "D",
		2: "L",
		3: "U",
	}
	parsed.direction = directions[directionNum]
	parsed.steps = length
	return
}

func calculateLavaArea(points []Point) int64 {
	// Convert the points to edges and calculate boundary points using Pick's theorem
	edges := 0
	area := int64(0)
	for i := 1; i < len(points); i++ {
		dx, dy := points[i].x-points[i-1].x, points[i].y-points[i-1].y
		edges += int(gcd(abs(dx), abs(dy)))
	}

	// Close the loop if it is not already closed
	if points[0] != points[len(points)-1] {
		dx, dy := points[0].x-points[len(points)-1].x, points[0].y-points[len(points)-1].y
		edges += int(gcd(abs(dx), abs(dy)))
		points = append(points, points[0])
	}

	// Using Pick's theorem and calculate the area
	internalPoints := int64((edges/2)-len(points)+1)
	area = internalPoints

	return area
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
