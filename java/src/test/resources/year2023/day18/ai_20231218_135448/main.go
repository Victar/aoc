
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Point struct {
	x, y int64
}

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := []Point{{0, 0}} // start at origin
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		hexCode := line[9:15]
		d, distance := parseInstruction(hexCode)
		addPoints(d, distance, &points)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Compute the area using Pick's theorem: A = i + b/2 - 1
	// Where A is the area, i the number of interior lattice points,
	// and b the number of boundary lattice points.
	interior, boundary := countPoints(points)
	area := interior + boundary/2 - 1

	// Print the total volcanic cubic meters:
	fmt.Println(area)
}

func parseInstruction(hexCode string) (Direction, int64) {
	dirCode, _ := strconv.ParseInt(string(hexCode[len(hexCode)-1]), 16, 64)
	distance, _ := strconv.ParseInt(hexCode[:len(hexCode)-1], 16, 64)
	return Direction(dirCode), distance
}

func addPoints(dir Direction, distance int64, points *[]Point) {
	last := (*points)[len(*points)-1]
	var dx, dy int64
	switch dir {
	case Right:
		dx = 1
	case Down:
		dy = -1
	case Left:
		dx = -1
	case Up:
		dy = 1
	}
	for i := int64(0); i < distance; i++ {
		newPoint := Point{last.x + dx, last.y + dy}
		*points = append(*points, newPoint)
		last = newPoint
	}
}

func countPoints(points []Point) (int64, int64) {
	var interior, boundary int64
	// Use a map to count unique points on the boundary
	pointMap := make(map[Point]int)

	for i := 0; i < len(points)-1; i++ {
		pointMap[points[i]]++
		xDiff := points[i+1].x - points[i].x
		yDiff := points[i+1].y - points[i].y

		// Count boundary points. If xDiff or yDiff is 0, it's a boundary point
		if xDiff == 0 || yDiff == 0 {
			boundary += abs(xDiff) + abs(yDiff) - 1 // subtract one to avoid double counting corners
		} else {
			// If it's diagonal, we're adding interior points; count them
			// Each diagonal contributes to additional interior points
			// Calculate by looking for GCD of delta x and y
			interior += abs(gcd(xDiff, yDiff)) - 1 // again, subtract one to avoid double counting corners
		}
	}

	for _, count := range pointMap {
		if count == 1 { // if the point is only visited once, it's a boundary point
			boundary++
		}
	}

	return interior, boundary
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
