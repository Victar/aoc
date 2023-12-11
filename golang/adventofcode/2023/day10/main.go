package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "10"

func main() {
	runSilver()
	//runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/sample.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]rune{}

	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	fmt.Println(grid)

	fmt.Println(solve(grid))
}

type Point struct {
	x, y int
}

var dirs = map[rune][]Point{
	'|': {{0, -1}, {0, 1}},
	'-': {{-1, 0}, {1, 0}},
	'F': {{1, 0}, {0, -1}},
	'7': {{-1, 0}, {0, -1}},
	'L': {{0, 1}, {1, 0}},
	'J': {{0, 1}, {-1, 0}},
}

func nextDirection(grid [][]rune, from, to Point) []Point {
	dirOptions, ok := dirs[grid[to.y][to.x]]
	if !ok {
		return nil
	}

	for _, opt := range dirOptions {
		next := Point{to.x + opt.x, to.y + opt.y}
		if next != from {
			return []Point{next}
		}
	}

	return nil
}

func findStart(grid [][]rune) Point {
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				return Point{x, y}
			}
		}
	}
	return Point{-1, -1} // not found
}

func solve(grid [][]rune) int {
	start := findStart(grid)
	if start.x == -1 || start.y == -1 {
		fmt.Println("Start position (S) not found in the input")
		return -1
	}

	maxDistance := 0
	distances := make(map[Point]int)
	queue := []Point{start}
	distances[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		nextTiles := nextDirection(grid, start, current)

		for _, next := range nextTiles {
			if _, visited := distances[next]; !visited && grid[next.y][next.x] != '.' {
				distances[next] = distances[current] + 1
				queue = append(queue, next)

				if distances[next] > maxDistance {
					maxDistance = distances[next]
				}
			}
		}
		start = current
	}

	return maxDistance
}

//func main() {
//	file, err := os.Open("input.txt")
//	if err != nil {
//		fmt.Printf("Error opening input file: %v\n", err)
//		return
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	grid := [][]rune{}
//
//	for scanner.Scan() {
//		grid = append(grid, []rune(scanner.Text()))
//	}
//	if err := scanner.Err(); err != nil {
//		fmt.Printf("Error reading input file: %v\n", err)
//		return
//	}
//
//	fmt.Println(solve(grid))
//}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
