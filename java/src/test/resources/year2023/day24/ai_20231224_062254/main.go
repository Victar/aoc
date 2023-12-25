
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
	px, py, pz int64 // Initial position
	vx, vy, vz int64 // Velocity
}

type Vector struct {
	x, y int64
}

// Function that reads input data from input.txt and returns a slice of Hailstones
func readInput(filename string) ([]Hailstone, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hailstones []Hailstone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " @ ")
		posStr := strings.Split(parts[0], ", ")
		velStr := strings.Split(parts[1], ", ")

		px, _ := strconv.ParseInt(posStr[0], 10, 64)
		py, _ := strconv.ParseInt(posStr[1], 10, 64)
		pz, _ := strconv.ParseInt(posStr[2], 10, 64)
		vx, _ := strconv.ParseInt(velStr[0], 10, 64)
		vy, _ := strconv.ParseInt(velStr[1], 10, 64)
		vz, _ := strconv.ParseInt(velStr[2], 10, 64)

		hailstones = append(hailstones, Hailstone{px, py, pz, vx, vy, vz})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hailstones, nil
}

// This function will check hailstone intersections within the provided bounds
// You will need to implement the logic to solve for intersection points
func countIntersections(hailstones []Hailstone) int {
	var count int
	// You have to implement intersection logic as per the task description
	// This will involve checking if the paths of two hailstones intersect within the specified range
	// ...

	return count
}

func main() {
	hailstones, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
		return
	}

	intersections := countIntersections(hailstones)
	fmt.Println(intersections)
}
