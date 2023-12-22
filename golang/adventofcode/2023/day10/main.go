package main

import (
	"adventofcode/util"
	"fmt"
	"os"
	"strings"
)

var DAY = "10"

func main() {
	runBoth()
}

type Point struct {
	c, r int
}

var dirs = map[rune][]Point{
	'|': {{0, -1}, {0, 1}},
	'-': {{-1, 0}, {1, 0}},
	'F': {{1, 0}, {0, 1}},
	'7': {{-1, 0}, {0, 1}},
	'L': {{1, 0}, {0, -1}},
	'J': {{0, -1}, {-1, 0}},
	'S': {{0, 1}}, //, {0, 1}},
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

func runBoth() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	start := findStart(grid)
	if start.c == -1 || start.r == -1 {
		fmt.Println("Start position (S) not found in the input")
	}

	maxDistance := 0
	visited := make(map[Point]bool)
	queue := []Point{}
	cur := start
	prev := start
	cycle := false
	for !cycle {
		queue = append(queue, cur)
		visited[cur] = true
		dirOptions, ok := dirs[grid[cur.r][cur.c]]
		if !ok {
			fmt.Println("not found: ", string(grid[cur.r][cur.c]), cur)
			break
		}
		for _, opt := range dirOptions {
			next := Point{cur.c + opt.c, cur.r + opt.r}
			if next != prev {
				prev = cur
				cur = next
				if visited[cur] {
					cycle = true
				}
				break
			}
		}
	}

	maxDistance = len(queue) / 2
	fmt.Println(maxDistance)
	shoelace := shoelaceFormula(queue)
	pick := shoelace - len(queue)/2 + 1
	fmt.Println(pick)
}

func shoelaceFormula(vertices []Point) int {
	n := len(vertices)
	if n < 3 {
		return 0
	}
	area := 0
	for i := 0; i < n-1; i++ {
		area += (vertices[i].c * vertices[i+1].r) - (vertices[i+1].c * vertices[i].r)
	}
	area += (vertices[n-1].c * vertices[0].r) - (vertices[0].c * vertices[n-1].r)
	area = area / 2
	if area < 0 {
		area = -area
	}
	return area
}

func printDistance(grid [][]rune, distances map[Point]int, insideArea map[Point]bool) {
	fmt.Println(strings.Repeat("=", 40))
	rows := len(grid)
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		str := ""
		for c := 0; c < cols; c++ {
			point := Point{c, r}
			cur := " " + string(grid[r][c]) + " "
			_, has := distances[point]
			if has {
				cur = "0"
			}
			if insideArea[point] {
				cur = "(" + string(grid[r][c]) + ")"
			}
			if has && insideArea[point] {
				cur = "Q"
				os.Exit(1)
			}
			str += cur
		}
		fmt.Println(str)
	}
}
