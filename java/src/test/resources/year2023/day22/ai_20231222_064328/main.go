
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Brick represents the structure of a single brick.
type Brick struct {
	x1, y1, z1 int
	x2, y2, z2 int
}

func main() {
	bricks := readInput("input.txt")
	safeToDisintegrate := countSafeDisintegrations(bricks)
	fmt.Println(safeToDisintegrate)
}

// readInput reads and parses the input file.
func readInput(filename string) []Brick {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var bricks []Brick
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		brick := parseBrick(line)
		bricks = append(bricks, brick)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return bricks
}

// parseBrick parses a line representing a single brick.
func parseBrick(line string) Brick {
	coords := strings.Split(line, "~")
	start := strings.Split(coords[0], ",")
	end := strings.Split(coords[1], ",")
	brick := Brick{
		x1: strToInt(start[0]),
		y1: strToInt(start[1]),
		z1: strToInt(start[2]),
		x2: strToInt(end[0]),
		y2: strToInt(end[1]),
		z2: strToInt(end[2]),
	}
	return brick
}

// strToInt converts a string to an integer.
func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// countSafeDisintegrations counts the number of bricks that can be safely disintegrated.
func countSafeDisintegrations(bricks []Brick) int {
	safeCount := 0
	for _, b := range bricks {
		if isSafeToDisintegrate(b, bricks) {
			safeCount++
		}
	}
	return safeCount
}

// isSafeToDisintegrate determines if a brick can be safely disintegrated.
func isSafeToDisintegrate(b Brick, bricks []Brick) bool {
	for _, other := range bricks {
		if b != other && isSupportedBy(b, other) {
			if wouldFallWithout(b, other, bricks) {
				return false
			}
		}
	}
	return true
}

// isSupportedBy checks if brick a is supported by brick b.
func isSupportedBy(a, b Brick) bool {
	return a.z1 == b.z1+1 &&
		overlap(a.x1, a.x2, b.x1, b.x2) &&
		overlap(a.y1, a.y2, b.y1, b.y2)
}

// overlap checks if there's an overlap in the given range.
func overlap(a1, a2, b1, b2 int) bool {
	return a1 <= b2 && a2 >= b1
}

// wouldFallWithout checks if the brick would fall without the support brick.
func wouldFallWithout(b, support Brick, bricks []Brick) bool {
	for _, other := range bricks {
		if support != other && b != other && isSupportedBy(b, other) {
			return false
		}
	}
	return true
}
