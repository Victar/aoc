
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x, y, z int
}

type Hailstone struct {
	position, velocity Vector
}

func main() {
	hailstones, err := readHailstones("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	intersections := countIntersections(hailstones)
	fmt.Println(intersections)
}

func readHailstones(filename string) ([]Hailstone, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hailstones []Hailstone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " @ ")
		posPart := strings.Split(parts[0], ", ")
		velPart := strings.Split(parts[1], ", ")

		px, _ := strconv.Atoi(posPart[0])
		py, _ := strconv.Atoi(posPart[1])
		pz, _ := strconv.Atoi(posPart[2])
		vx, _ := strconv.Atoi(velPart[0])
		vy, _ := strconv.Atoi(velPart[1])
		vz, _ := strconv.Atoi(velPart[2])

		hailstones = append(hailstones, Hailstone{
			position: Vector{x: px, y: py, z: pz},
			velocity: Vector{x: vx, y: vy, z: vz},
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hailstones, nil
}

func countIntersections(hailstones []Hailstone) int {
	// You will need to implement this function.
	// It should calculate all pair-wise intersections of hailstones' paths and return the number of intersections within the test area.
	return -1 // Stub return; you must replace this with your logic.
}

// Additional helper functions for calculating intersections, ranges, etc., go here...
