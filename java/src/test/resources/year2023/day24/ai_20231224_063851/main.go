
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
	pos, vel Vector
}

func main() {
	hailstones, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if len(hailstones) == 0 {
		fmt.Println("No hailstones provided in the input.")
		return
	}

	pos, ok := calculateRockPositionAndVelocity(hailstones)
	if !ok {
		fmt.Println("Unable to find a position and velocity to hit all hailstones.")
		return
	}
	fmt.Println(pos.x + pos.y + pos.z)
}

func parseInput(filename string) ([]Hailstone, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hailstones []Hailstone
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " @ ")
		if len(fields) != 2 {
			continue // Invalid line format
		}
		posParts := strings.Split(fields[0], ", ")
		velParts := strings.Split(fields[1], ", ")
		if len(posParts) != 3 || len(velParts) != 3 {
			continue // Invalid number of coordinates
		}

		pos, err := parseVector(posParts)
		if err != nil {
			return nil, err
		}
		vel, err := parseVector(velParts)
		if err != nil {
			return nil, err
		}

		hailstones = append(hailstones, Hailstone{pos, vel})
	}
	return hailstones, scanner.Err()
}

func parseVector(parts []string) (Vector, error) {
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return Vector{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Vector{}, err
	}
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		return Vector{}, err
	}
	return Vector{x, y, z}, nil
}

func calculateRockPositionAndVelocity(hailstones []Hailstone) (Vector, bool) {
	type collision struct {
		pos Vector
		t   int
	}

	// Assuming the rock needs to hit all hailstones
	collisions := make(map[int]collision)
	for i, h := range hailstones {
		t := h.pos.x / h.vel.x
		if t < 0 {
			return Vector{}, false // Hailstone has already passed
		}
		xPos := h.pos.x + t*h.vel.x
		yPos := h.pos.y + t*h.vel.y
		zPos := h.pos.z + t*h.vel.z
		coll, ok := collisions[t]

		if ok {
			if coll.pos.x != xPos || coll.pos.y != yPos || coll.pos.z != zPos {
				return Vector{}, false // Hailstones collide at different locations
			}
		} else {
			collisions[t] = collision{Vector{xPos, yPos, zPos}, t}
		}

		// Update the initial rock position to match the first hailstone
		if i == 0 {
			collisions[t] = collision{h.pos, t}
		}
	}

	// Find the common collision point
	var commonCollision collision
	for _, c := range collisions {
		commonCollision = c
		break
	}

	// The initial rock position will match the first collision point
	return commonCollision.pos, true
}

