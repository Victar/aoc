package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vec2 struct {
	X, Y float64
}

type Hailstone struct {
	Position, Velocity Vec2
}

func main() {
	// Open the input file
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day24/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the input file
	hailstones := make([]Hailstone, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " @ ")
		posStr := strings.Split(parts[0], ",")
		velStr := strings.Split(parts[1], ",")
		px, _ := strconv.ParseFloat(strings.TrimSpace(posStr[0]), 64)
		py, _ := strconv.ParseFloat(strings.TrimSpace(posStr[1]), 64)
		vx, _ := strconv.ParseFloat(strings.TrimSpace(velStr[0]), 64)
		vy, _ := strconv.ParseFloat(strings.TrimSpace(velStr[1]), 64)
		hailstones = append(hailstones, Hailstone{Position: Vec2{X: px, Y: py}, Velocity: Vec2{X: vx, Y: vy}})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Define the test area boundaries
	//minTestX, maxTestX := .0, 27.0
	//minTestY, maxTestY := 7.0, 27.0

	minTestX, maxTestX := 200000000000000.0, 400000000000000.0
	minTestY, maxTestY := 200000000000000.0, 400000000000000.0

	// Find number of intersections within the test area
	count := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if isIntersecting(hailstones[i], hailstones[j], minTestX, maxTestX, minTestY, maxTestY) {
				count++
			}
		}
	}

	fmt.Println(count)
}

// isIntersecting checks whether the paths of two hailstones intersect within a test area.
func isIntersecting(a, b Hailstone, minX, maxX, minY, maxY float64) bool {
	// Using the formula to check intersection:
	// a.Position + t * a.Velocity = b.Position + s * b.Velocity
	// Solve for t and s, and see if they generate a valid intersection point
	determinant := a.Velocity.X*b.Velocity.Y - a.Velocity.Y*b.Velocity.X
	if determinant == 0 {
		// The paths are parallel and will never intersect
		return false
	}

	dx := b.Position.X - a.Position.X
	dy := b.Position.Y - a.Position.Y

	tNumerator := dx*b.Velocity.Y - dy*b.Velocity.X
	sNumerator := dx*a.Velocity.Y - dy*a.Velocity.X

	if determinant < 0 {
		determinant = -determinant
		tNumerator = -tNumerator
		sNumerator = -sNumerator
	}

	// Make sure the intersection is not happening in reverse time
	if tNumerator < 0 || sNumerator < 0 {
		return false
	}

	// Find intersection point
	t := tNumerator / determinant

	intersectX := a.Position.X + t*a.Velocity.X
	intersectY := a.Position.Y + t*a.Velocity.Y
	//fmt.Println(intersectX, intersectY, t, a, b)
	// Check if intersection is within the test area
	return intersectX >= minX && intersectX <= maxX && intersectY >= minY && intersectY <= maxY
}
