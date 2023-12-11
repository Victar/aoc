
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var directions = [...]Point{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

// Helper function to see if a point is within the grid boundaries.
func inBounds(grid [][]string, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x])
}

// Helper function to get connecting points for a given pipe type.
func validMoves(pipe string) []Point {
	switch pipe {
	case "|":
		return []Point{directions[0], directions[1]}
	case "-":
		return []Point{directions[2], directions[3]}
	case "L":
		return []Point{directions[0], directions[3]}
	case "J":
		return []Point{directions[0], directions[2]}
	case "7":
		return []Point{directions[1], directions[2]}
	case "F":
		return []Point{directions[1], directions[3]}
	case ".", "S":
		// This acts as a wildcard.
		return directions[:]
	default:
		return []Point{}
	}
}

// Helper function to read a grid from a file.
func readGridFromFile(filename string) ([][]string, Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, Point{}, err
	}
	defer file.Close()

	var grid [][]string
	var startX, startY int
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)

		if x := strings.Index(line, "S"); x != -1 {
			startX, startY = x, y
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		return nil, Point{}, err
	}
	return grid, Point{startX, startY}, nil
}

// Helper function to determine if two points are in the same direction.
func sameDirection(p1, p2 Point) bool {
	return (p1.x == p2.x && p1.y == p2.y) || (p1.x == -p2.x && p1.y == -p2.y)
}

// Function to solve the maze.
func solveMaze(grid [][]string, start Point) int {
	var maxDistance int
	visited := make(map[Point]bool)
	var dfs func(p Point, from Point, distance int)
	dfs = func(p Point, from Point, distance int) {
		if visited[p] {
			return
		}
		visited[p] = true
		if distance > maxDistance {
			maxDistance = distance
		}

		pipe := "S"
		if grid[p.x][p.y] != "S" {
			pipe = grid[p.x][p.y]
		}
		validDirs := validMoves(pipe)
		for _, dir := range validDirs {
			next := Point{p.x + dir.x, p.y + dir.y}
			if !inBounds(grid, next.x, next.y) || sameDirection(dir, from) {
				continue
			}
			if grid[next.x][next.y] != "." && !visited[next] {
				dfs(next, Point{-dir.x, -dir.y}, distance+1)
			}
		}
	}
	dfs(start, Point{0, 0}, 0)

	return maxDistance
}

func main() {
	grid, start, err := readGridFromFile("input.txt")
	if err != nil {
		log.Fatalf("Unable to read from file: %v", err)
	}

	maxDistance := solveMaze(grid, start)
	fmt.Print(maxDistance)
}
