
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
	px, py, vx, vy int
}

// isIntersectionInRange checks if the hailstones intersect within the given range.
func isIntersectionInRange(a, b Hailstone, minX, maxX, minY, maxY int) bool {
	// Time at which the two hailstones intersect can be found by solving two equations:
	// a.px + t * a.vx = b.px + t * b.vx
	// a.py + t * a.vy = b.py + t * b.vy

	// Reformulate as t = (b.px - a.px) / (a.vx - b.vx) and similarly for y
	// The intersection is valid only if both t values for x and y are equal.

	deltaVx := a.vx - b.vx
	deltaVy := a.vy - b.vy
	deltaPx := b.px - a.px
	deltaPy := b.py - a.py

	// If velocities are the same in any direction, paths are parallel and do not intersect unless they are already at the same point.
	if deltaVx == 0 || deltaVy == 0 {
		return false
	}

	// Time at which the paths intersect in both x and y dimensions.
	tX := float64(deltaPx) / float64(deltaVx)
	tY := float64(deltaPy) / float64(deltaVy)

	if tX != tY || tX < 0 {
		// The paths do not intersect or intersect in the past.
		return false
	}

	// Check if the intersection point is within the given range.
	intersectionX := a.px + int(tX)*a.vx
	intersectionY := a.py + int(tY)*a.vy

	return minX <= intersectionX && intersectionX <= maxX && minY <= intersectionY && intersectionY <= maxY
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: unable to open input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hailstones := make([]Hailstone, 0)

	// Minimum and maximum values for the X and Y coordinates.
	minX, maxX := 200000000000000, 400000000000000
	minY, maxY := minX, maxX

	// Read the input and parse the hailstone data
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			continue
		}

		pos := strings.Split(fields[0], ",")
		vel := strings.Split(fields[1], "@")

		// Parse positions
		px, errPx := strconv.Atoi(strings.TrimSpace(pos[0]))
		py, errPy := strconv.Atoi(strings.TrimSpace(pos[1]))
		if errPx != nil || errPy != nil {
			fmt.Println("Error: invalid position data in input file")
			os.Exit(1)
		}

		// Parse velocities
		vx, errVx := strconv.Atoi(strings.TrimSpace(vel[0]))
		vy, errVy := strconv.Atoi(strings.TrimSpace(vel[1]))
		if errVx != nil || errVy != nil {
			fmt.Println("Error: invalid velocity data in input file")
			os.Exit(1)
		}

		hailstones = append(hailstones, Hailstone{px: px, py: py, vx: vx, vy: vy})
	}

	// Check for errors during scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: unable to read input file:", err)
		os.Exit(1)
	}

	// Now, loop over the pairs of hailstones and count the intersections that occur within the range.
	count := 0
	for i, a := range hailstones {
		for j, b := range hailstones {
			if i >= j {
				continue // Avoid double counting and self-intersection
			}
			if isIntersectionInRange(a, b, minX, maxX, minY, maxY) {
				count++
			}
		}
	}

	// Print the result
	fmt.Println(count)
}
