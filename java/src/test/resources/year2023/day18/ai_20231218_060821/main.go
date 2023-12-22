
// main.go
package main

import (
	"bufio"
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make(map[Point]bool)
	var x, y int
	maxX, maxY, minX, minY := 0, 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		instructions := strings.Split(line, " ")
		for i := 0; i < len(instructions); i += 3 {
			direction := instructions[i]
			length, _ := strconv.Atoi(strings.TrimSpace(instructions[i+1]))

			for j := 0; j < length; j++ {
				grid[Point{x, y}] = true
				switch direction {
				case "R":
					x++
				case "L":
					x--
				case "U":
					y++
				case "D":
					y--
				}

				// Record the maximum and minimum extents of the trench
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	capacity := 0
	// Calculate the number of cubic meters the lagoon could hold by filling in the interior
	for i := minX + 1; i < maxX; i++ {
		for j := minY + 1; j < maxY; j++ {
			if grid[Point{i, j}] {
				// This is an edge; exclude from filling
				continue
			}

			isInside := true
			for x := i; x >= minX; x-- {
				if !grid[Point{x, j}] {
					isInside = false
					break
				}
			}
			if isInside {
				for x := i; x <= maxX; x++ {
					if !grid[Point{x, j}] {
						isInside = false
						break
					}
				}
			}
			if isInside {
				capacity++
			}
		}
	}

	fmt.Println(capacity)
}
