package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int64
}

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day18/ai_20231218_135834/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []string
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		instructions = append(instructions, strings.Split(line, " ")...)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var points []Point
	current := Point{0, 0}
	points = append(points, current)

	for _, instruction := range instructions {
		hexColor := strings.TrimSpace(instruction[8:15])
		distance, direction := parseInstruction(hexColor)
		switch direction {
		case 'R':
			current.X += distance
		case 'U':
			current.Y += distance
		case 'L':
			current.X -= distance
		case 'D':
			current.Y -= distance
		}
		points = append(points, current)
	}

	area := calculateLavaArea(points)
	fmt.Println(area)
}

func parseInstruction(hexColor string) (int64, rune) {
	distance, _ := strconv.ParseInt(hexColor[:5], 16, 64)
	directionCode, _ := strconv.ParseInt(string(hexColor[5]), 16, 0)
	var direction rune
	switch directionCode {
	case 0:
		direction = 'R'
	case 1:
		direction = 'D'
	case 2:
		direction = 'L'
	case 3:
		direction = 'U'
	}
	return distance, direction
}

// Pick's theorem: A = I + B/2 - 1, where A is the area, I is the number of integer lattice points inside the polygon,
// B is the number of integer lattice points on the border, and -1 is a constant adjustment.
func calculateLavaArea(points []Point) *big.Int {
	var boundaryPoints int64

	// Calculate the boundary points by gcd algorithm
	for i := 0; i < len(points); i++ {
		nextIndex := (i + 1) % len(points)
		boundaryPoints += gcd(abs(points[nextIndex].X-points[i].X), abs(points[nextIndex].Y-points[i].Y))
	}

	// Calculate the interior points using Shoelace formula
	var area *big.Int
	area = big.NewInt(0)
	for i := 0; i < len(points); i++ {
		nextIndex := (i + 1) % len(points)
		area.Add(area, big.NewInt((points[i].X*points[nextIndex].Y)-(points[i].Y*points[nextIndex].X)))
	}
	area.Div(area, big.NewInt(2))
	area.Abs(area)

	// Convert to int64 for the calculation of boundaryPoints/2 - 1
	lavaArea := new(big.Int).SetInt64((boundaryPoints - 2) / 2)
	lavaArea.Add(area, lavaArea)

	return lavaArea
}

// gcd calculates the greatest common divisor using Euclidean algorithm
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// abs calculates the absolute value of a 64-bit integer
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
