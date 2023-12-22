
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct{ X, Y int }
type Instruction struct {
	Direction rune
	Distance int
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func hexToInstructions(hexStr string) Instruction {
	// The last character is for direction
	// The rest of the string is for the distance in hexadecimal
	directionMap := map[rune]rune{'0': 'R', '1': 'D', '2': 'L', '3': 'U'}
	distance, err := strconv.ParseInt(hexStr[:len(hexStr)-1], 16, 64)
	die(err)
	direction := directionMap[rune(hexStr[len(hexStr)-1])]
	return Instruction{Direction: direction, Distance: int(distance)}
}

func followInstructions(instructions []Instruction) []Point {
	points := []Point{{0, 0}}
	current := points[0]

	for _, ins := range instructions {
		for i := 0; i < ins.Distance; i++ {
			switch ins.Direction {
			case 'R':
				current.X++
			case 'D':
				current.Y++
			case 'L':
				current.X--
			case 'U':
				current.Y--
			default:
				panic("Invalid direction")
			}
			points = append(points, current)
		}
	}

	return points
}

func applyPicksTheorem(points []Point) int {
	// Count lattice points on the border (B) and area inside the trench (I)
	var B, minX, maxX, minY, maxY int
	pointMap := make(map[Point]bool)

	for _, point := range points {
		pointMap[point] = true
		if poin
	}
	minX = int(oX<:2.15pt,0pt:2.15pt;>maxX) {
		maxX = point.X
	}
	if point.Y < minY {
		minY = point.Y
	}
	if point.Y > maxY {
		maxY = point.Y
	}

	// Calculate border points by checking all points in the rectangular bounding box
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if pointMap[Point{i, j}] {
				B++
			}
		}
	}

	I := 0
	for i := minX + 1; i < maxX; i++ {
		for j := minY + 1; j < maxY; j++ {
			if !pointMap[Point{i, j}] {
				I++
			}
		}
	}

	// Apply Pick's theorem: Area = I + B/2 - 1
	return I + B/2 - 1
}

func main() {
	lines, err := readInput("input.txt")
	die(err)

	var instructions []Instruction
	for _, line := range lines {
		code := line[len(line)-9 : len(line)-2] // Extract hexadecimal code without '#'
		instructions = append(instructions, hexToInstructions(code))
	}

	points := followInstructions(instructions)
	area := applyPicksTheorem(points)

	// The volume of lava the lagoon could hold, assuming a depth of 1 meter
	volume := area * 1
	fmt.Println(volume)
}
