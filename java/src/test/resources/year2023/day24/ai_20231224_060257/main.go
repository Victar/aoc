
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
	px, py   int64 // Initial position
	vx, vy   int64 // Velocity
}

func main() {
	hailstones := readHailstones("input.txt")
	testAreaMin, testAreaMax := int64(200000000000000), int64(400000000000000)
	intersections := countIntersections(hailstones, testAreaMin, testAreaMax)
	fmt.Println(intersections)
}

// readHailstones reads the input file and returns a slice of Hailstone structs.
func readHailstones(filename string) []Hailstone {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var hailstones []Hailstone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " @ ")
		position := strings.Split(parts[0], ", ")
		velocity := strings.Split(parts[1], ", ")
		px, _ := strconv.ParseInt(position[0], 10, 64)
		py, _ := strconv.ParseInt(position[1], 10, 64)
		vx, _ := strconv.ParseInt(velocity[0], 10, 64)
		vy, _ := strconv.ParseInt(velocity[1], 10, 64)

		hailstones = append(hailstones, Hailstone{px, py, vx, vy})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return hailstones
}

// countIntersections counts the number of hailstone path intersections within the test area.
func countIntersections(hailstones []Hailstone, testAreaMin, testAreaMax int64) int {
	intersectionCount := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1, h2 := hailstones[i], hailstones[j]
			if intersect(h1, h2, testAreaMin, testAreaMax) {
				intersectionCount++
			}
		}
	}
	return intersectionCount
}

// intersect checks if the paths of two hailstones intersect within the given test area.
func intersect(h1, h2 Hailstone, min, max int64) bool {
	// Calculate intersection point using algebraic equations
	dvx, dvy := h1.vx-h2.vx, h1.vy-h2.vy
	dpx, dpy := h1.px-h2.px, h1.py-h2.py
	if dvx == 0 && dvy == 0 {
		return false // Parallel paths, no intersection
	}
	// Intersection at time t means:
	// h1.px + h1.vx*t = h2.px + h2.vx*t
	// h1.py + h1.vy*t = h2.py + h2.vy*t
	// We solve for t to find the intersection point and if it's in the area
	t := float64(-dpx) / float64(dvx)
	intersectX := float64(h1.px) + float64(h1.vx)*t
	intersectY := float64(h1.py) + float64(h1.vy)*t
	return t >= 0 && int64(intersectX) >= min && int64(intersectX) <= max && int64(intersectY) >= min && int64(intersectY) <= max
}

