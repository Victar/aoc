
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	X, Y int64
}

type Hailstone struct {
	Position, Velocity Vector
}

func main() {
	hailstones := readInput("input.txt")
	count := countIntersections(hailstones, 200000000000000, 400000000000000)
	fmt.Println(count)
}

func readInput(filename string) []Hailstone {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var hailstones []Hailstone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		position := parseVector(fields[0])
		velocity := parseVector(fields[3])

		hailstones = append(hailstones, Hailstone{Position: position, Velocity: velocity})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return hailstones
}

func parseVector(field string) Vector {
	numbers := strings.Split(field, ",")
	x, errX := strconv.ParseInt(numbers[0], 10, 64)
	y, errY := strconv.ParseInt(numbers[1], 10, 64)
	if errX != nil || errY != nil {
		panic("Error parsing vector components")
	}
	return Vector{X: x, Y: y}
}

func countIntersections(hailstones []Hailstone, min, max int64) int {
	count := 0
	for i, stoneA := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			if willIntersect(stoneA, hailstones[j], min, max) {
				count++
			}
		}
	}
	return count
}

func willIntersect(a, b Hailstone, min, max int64) bool {
	dx := a.Position.X - b.Position.X
	dy := a.Position.Y - b.Position.Y
	vx := a.Velocity.X - b.Velocity.X
	vy := a.Velocity.Y - b.Velocity.Y

	if vx == 0 && vy == 0 {
		return false // parallel and never intersect
	}

	t, u := solveLinearEquation(vx, -dx, vy, -dy)
	return t >= 0 && u >= 0 && withinBounds(a.Position.X+vx*t, min, max) && withinBounds(a.Position.Y+vy*t, min, max)
}

func solveLinearEquation(a, c, b, d int64) (int64, int64) {
	det := a*b - b*a
	if det == 0 {
		return 0, 0 // No solution or infinite solutions
	}
	return d*b - b*c, a*d - d*a
}

func withinBounds(value, min, max int64) bool {
	return value >= min && value <= max
}
