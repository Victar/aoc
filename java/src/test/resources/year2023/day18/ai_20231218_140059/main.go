
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	// Read entire file line by line and collect each color code
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	var x, y, minX, minY, maxX, maxY, boundaryPoints int64
	var interiorPoints float64
	coordinates := make(map[[2]int64]struct{})
	coordinates[[2]int64{0, 0}] = struct{}{}

	for _, instruction := range instructions {
		instructionParts := instruction[9 : len(instruction)-1] // Get the color code
		steps, dir := decodeInstruction(instructionParts)
		for i := int64(0); i < steps; i++ {
			switch dir {
			case 0: // R
				x++
			case 1: // D
				y++
			case 2: // L
				x--
			case 3: // U
				y--
			}
			// Calculate the bounding box and keep track of boundary points
			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
			coordinates[[2]int64{x, y}] = struct{}{}
			boundaryPoints++
		}
	}

	// Calculate the interior points using Pick's theorem
	width := maxX - minX + 1
	height := maxY - minY + 1
	// Account for the grid points on the boundary
	interiorPoints = float64(width*height) - float64(boundaryPoints)/2 - 1

	lavaVolume := int64(interiorPoints) + boundaryPoints
	fmt.Println(lavaVolume)
}

// Helper function to decode instruction from hexadecimal color code
func decodeInstruction(code string) (int64, int64) {
	stepsHex := code[:5] // First 5 characters are steps
	steps, err := strconv.ParseInt(stepsHex, 16, 64) // parsed as base 16 (hex)
	if err != nil {
		panic(err)
	}

	dirHex := code[5:]                           // Last character is direction
	dir, err := strconv.ParseInt(dirHex, 16, 64) // parsed as base 16 (hex)
	if err != nil {
		panic(err)
	}
	dir &= 3 // Only interested in the last 2 bits for direction

	return steps, dir
}
