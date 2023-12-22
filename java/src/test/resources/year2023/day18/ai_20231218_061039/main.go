
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
	X, Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lagoon := make(map[Point]bool)
	directions := []string{"R", "U", "L", "D"}
	dx := map[string]int{"R": 1, "U": 0, "L": -1, "D": 0}
	dy := map[string]int{"R": 0, "U": -1, "L": 0, "D": 1}
	x, y := 0, 0
	lagoon[Point{x, y}] = true

	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, " ")
		for i := 0; i < len(commands); i += 3 {
			dir := commands[i]
			steps, _ := strconv.Atoi(commands[i+1])
			for step := 0; step < steps; step++ {
				x += dx[dir]
				y += dy[dir]
				lagoon[Point{x, y}] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Filling the lagoon
	filledLagoon := make(map[Point]bool)
	for p := range lagoon {
		for dir := range dx {
			np := Point{p.X + dx[dir], p.Y + dy[dir]}
			if !lagoon[np] {
				fill(lagoon, filledLagoon, np)
			}
		}
	}
	capacity := len(filledLagoon)
	fmt.Println(capacity)
}

func fill(lagoon, filledLagoon map[Point]bool, p Point) {
	if filledLagoon[p] {
		return
	}
	if lagoon[p] {
		filledLagoon[p] = true
		return
	}
	dirs := []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	for _, dir := range dirs {
		np := Point{p.X + dir.X, p.Y + dir.Y}
		fill(lagoon, filledLagoon, np)
	}
}
