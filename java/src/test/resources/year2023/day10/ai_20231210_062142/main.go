package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	maze, startPos := loadMaze("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day10/input.txt")
	if maze == nil {
		fmt.Println("Error: Maze could not be loaded.")
		return
	}

	farthestDistance := findFarthestDistance(maze, startPos)
	fmt.Println(farthestDistance)
}

type Pos struct{ x, y int }

func loadMaze(filename string) ([][]rune, Pos) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var maze [][]rune
	startPos := Pos{-1, -1}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		x := -1
		if startPos.y == -1 {
			x = findStartPos(row)
		}
		if x != -1 {
			startPos = Pos{x, y}
		}
		maze = append(maze, row)
		y++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return maze, startPos
}

func findStartPos(row []rune) int {
	for i, c := range row {
		if c == 'S' {
			return i
		}
	}
	return -1
}

var directions = []Pos{
	{0, 1},  // down
	{1, 0},  // right
	{0, -1}, // up
	{-1, 0}, // left
}

func findFarthestDistance(maze [][]rune, startPos Pos) int {
	farthest := 0
	visited := make(map[Pos]int)
	var search func(pos Pos, distance int)
	search = func(pos Pos, distance int) {
		if d, seen := visited[pos]; seen && d <= distance {
			return
		}
		visited[pos] = distance
		if distance > farthest {
			farthest = distance
		}
		for _, dir := range directions {
			nextPos := Pos{pos.x + dir.x, pos.y + dir.y}
			if isValid(nextPos, maze, dir) {
				search(nextPos, distance+1)
			}
		}
	}
	search(startPos, 0)
	return farthest
}

func isValid(pos Pos, maze [][]rune, dir Pos) bool {
	if pos.y < 0 || pos.y >= len(maze) || pos.x < 0 || pos.x >= len(maze[pos.y]) {
		return false
	}

	switch maze[pos.y][pos.x] {
	case '-':
		return dir.y == 0 // horizontal pipe, vertical movement not allowed
	case '|':
		return dir.x == 0 // vertical pipe, horizontal movement not allowed
	case '7':
		return dir == Pos{-1, 0} || dir == Pos{0, -1} // 7 bend
	case 'J':
		return dir == Pos{1, 0} || dir == Pos{0, -1} // J bend
	case 'L':
		return dir == Pos{0, 1} || dir == Pos{-1, 0} // L bend
	case 'F':
		return dir == Pos{0, 1} || dir == Pos{1, 0} // F bend
	case '.':
		return false // ground, no pipe
	case 'S':
		return true // start, can go any direction
	default:
		return true // pipe part of the loop
	}
}

//
//func main() {
//	maze, startPos := loadMaze("input.txt")
//	if maze == nil {
//		fmt.Println("Error: Maze could not be loaded.")
//		return
//	}
//
//	farthestDistance := findFarthestDistance(maze, startPos)
//	fmt.Println(farthestDistance)
//}
