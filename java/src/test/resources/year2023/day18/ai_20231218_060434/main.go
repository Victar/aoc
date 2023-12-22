
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, strings.Fields(line)...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	directions := map[string]Point{
		"U": {0, -1},
		"D": {0, 1},
		"L": {-1, 0},
		"R": {1, 0},
	}

	// Initialise the map with a starting point
	lagoon := map[Point]bool{{0, 0}: true}
	pos := Point{0, 0}

	for i := 0; i < len(instructions); i += 3 {
		dir := string(instructions[i][0])
		steps, _ := strconv.Atoi(instructions[i][1:])
		dX, dY := directions[dir].x, directions[dir].y

		for step := 0; step < steps; step++ {
			pos.x += dX
			pos.y += dY
			lagoon[pos] = true
		}
	}

	// Find borders to calculate the enclosed area
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for p := range lagoon {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// Fill the trench
	var count int
	for y := minY - 1; y <= maxY+1; y++ {
		inTrench := false
		for x := minX - 1; x <= maxX+1; x++ {
			if lagoon[Point{x, y}] {
				inTrench = !inTrench
				continue
			}
			if inTrench {
				lagoon[Point{x, y}] = true
				count++
			}
		}
	}

	fmt.Println(count)
}
