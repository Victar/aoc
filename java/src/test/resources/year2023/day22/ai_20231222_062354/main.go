
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

type Brick struct {
	start, end Point
}

func (b *Brick) Supports(other *Brick) bool {
	if b.start.z != other.end.z+1 {
		// If b is not exactly one level below other, it can't support it
		return false
	}
	// Assume that bricks lay flat. If any point of brick `other` is above `b`
	return (other.start.x >= b.start.x && other.start.x <= b.end.x ||
		other.end.x >= b.start.x && other.end.x <= b.end.x) &&
		(other.start.y >= b.start.y && other.start.y <= b.end.y ||
			other.end.y >= b.start.y && other.end.y <= b.end.y)
}

// Parses a line from input to get a Brick struct
func NewBrick(line string) Brick {
	points := strings.Split(line, "~")
	startStr := strings.Split(points[0], ",")
	endStr := strings.Split(points[1], ",")

	start := Point{
		x: parseInt(startStr[0]),
		y: parseInt(startStr[1]),
		z: parseInt(startStr[2]),
	}
	end := Point{
		x: parseInt(endStr[0]),
		y: parseInt(endStr[1]),
		z: parseInt(endStr[2]),
	}
	return Brick{start: start, end: end}
}

// Parses an integer from a string
func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var bricks []Brick
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		brick := NewBrick(line)
		bricks = append(bricks, brick)
	}

	// Assuming bricks are ordered by their appearance in the stack (from top to bottom)
	safeToRemove := 0
	for i, brick := range bricks {
		cannotRemove := false
		for j := i + 1; j < len(bricks); j++ {
			// "If any other brick would fall as a result of removing this brick, it is not safe to remove"
			if bricks[j].Supports(&brick) {
				cannotRemove = true
				break
			}
		}
		if !cannotRemove {
			safeToRemove++
		}
	}

	fmt.Println(safeToRemove)
}
