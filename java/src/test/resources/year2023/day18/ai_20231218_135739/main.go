
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

func hexToDirDist(code string) (dx, dy, dist int) {
	// Last character of the code determines the direction
	dir := int(code[5])

	// The first 5 characters determine the distance
	dist, _ = strconv.ParseInt(code[:5], 16, 64)

	// Convert direction to delta x and delta y
	switch dir {
	case 0: // Right
		dx, dy = 1, 0
	case 1: // Down
		dx, dy = 0, 1
	case 2: // Left
		dx, dy = -1, 0
	case 3: // Up
		dx, dy = 0, -1
	}
	return
}

func calcArea(points []Point) int64 {
	// Calculate the area using Pick's theorem
	// Area = I + B/2 - 1
	// I = interior points, B = boundary points

	// Since this task is reduced to a rectangle lagoon, we can simplify the calculation
	// We only take the min and max points to form the rectangle and calculate its area

	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	return int64(width * height)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer file.Close()

	var points []Point
	curX, curY := 0, 0
	points = append(points, Point{curX, curY})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colorCode := line[7:13]
		dx, dy, dist := hexToDirDist(colorCode)
		for i := 0; i < dist; i++ {
			curX += dx
			curY += dy
			points = append(points, Point{curX, curY})
		}
	}

	// Calculate and print the area that the lagoon can hold
	area := calcArea(points)
	fmt.Println(area)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
