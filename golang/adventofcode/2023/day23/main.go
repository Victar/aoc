package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
)

var DAY = "23"

func main() {
	//runSilver()
	runGold()
}

type Point struct {
	c int
	r int
}
type State struct {
	point Point
	dir   Point
}

var LEFT = Point{-1, 0}
var RIGHT = Point{1, 0}
var UP = Point{0, -1}
var DOWN = Point{0, 1}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]rune{}
	for _, line := range lines {
		l := []rune{}
		for _, c := range line {
			l = append(l, c)
		}
		grid = append(grid, l)
	}
	start := Point{1, 0}
	end := Point{len(grid[0]) - 2, len(grid) - 1}
	fmt.Println(countDFS(grid, start, end))
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]rune{}
	for _, line := range lines {
		l := []rune{}
		for _, c := range line {
			l = append(l, c)
		}
		grid = append(grid, l)
	}
	start := Point{1, 0}
	end := Point{len(grid[0]) - 2, len(grid) - 1}
	visited := make(map[Point]bool)
	fmt.Println(countDFSGold(grid, start, end, DOWN, 0, visited))
}

var globMaxGold = 0

func countDFSGold(grid [][]rune, current, end, currentDir Point, step int, visited map[Point]bool) int {

	if current == end {
		return step
	}
	rows := len(grid)
	cols := len(grid[0])
	curMax := 0
	for _, dir := range getDirs(currentDir) {
		newPoint := Point{current.c + dir.c, current.r + dir.r}
		if newPoint.c >= 0 && newPoint.c < cols && newPoint.r >= 0 && newPoint.r < rows && grid[newPoint.r][newPoint.c] != '#' {
			if !visited[newPoint] {
				visited[newPoint] = true
				dfsMax := countDFSGold(grid, newPoint, end, dir, step+1, visited)
				if dfsMax > curMax {
					curMax = dfsMax
					if newPoint == end && globMaxGold < curMax {
						globMaxGold = curMax
						fmt.Println("new max", curMax)
					}
				}
				delete(visited, newPoint)
			}
		}
	}
	return curMax
}

func countDFS(grid [][]rune, start Point, end Point) int {
	maxRoute := 0
	visited := make(map[Point]int)
	rows := len(grid)
	cols := len(grid[0])
	start2 := State{start, DOWN}
	stack := []State{start2}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curDistance, has := visited[cur.point]
		if !has {
			fmt.Println("not found for", cur)
		}
		if cur.point == end && curDistance > maxRoute {
			fmt.Println("new max", maxRoute, curDistance, cur)
			maxRoute = curDistance
		}
		for _, dir := range getDirs(cur.dir) {
			newPoint := Point{cur.point.c + dir.c, cur.point.r + dir.r}
			nextState := State{newPoint, dir}
			if newPoint.c >= 0 && newPoint.c < cols && newPoint.r >= 0 && newPoint.r < rows && grid[newPoint.r][newPoint.c] != '#' && allowPath(grid, newPoint, dir) {
				if val, found := visited[newPoint]; !found || val < curDistance+1 {
					visited[newPoint] = curDistance + 1
					stack = append(stack, nextState)
				}
			}
		}
	}
	return maxRoute
}

func allowPath(grid [][]rune, current, dir Point) bool {
	c := grid[current.r][current.c]
	if c == '.' {
		return true
	}
	if c == 'v' {
		return dir == DOWN
	}
	if c == '^' {
		return dir == UP
	}
	if c == '>' {
		return dir == RIGHT
	}
	if c == '<' {
		return dir == LEFT
	}
	fmt.Println("unexpected case", string(c))
	return false
}

func getDirs(dir Point) []Point {
	if dir == LEFT || dir == RIGHT {
		return []Point{dir, UP, DOWN}
	} else {
		return []Point{dir, LEFT, RIGHT}
	}
}

func printGridP(grid [][]rune, visited map[Point]int) {
	rows := len(grid)
	cols := len(grid[0])
	for r := 0; r < rows; r++ {
		line := ""
		for c := 0; c < cols; c++ {
			cur := " " + string(grid[r][c]) + "  "
			p := Point{c, r}
			val, has := visited[p]
			if has {
				cur = " " + strconv.Itoa(val) + " "
			}
			for len(cur) < 4 {
				cur += " "
			}
			line += cur
		}
		fmt.Println(line)
	}
}
