package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
)

var DAY = "17"

func main() {
	runAny(1, 3)
	runAny(4, 10)
}

type Point struct {
	c int
	r int
}

var LEFT = Point{-1, 0}
var RIGHT = Point{1, 0}
var UP = Point{0, -1}
var DOWN = Point{0, 1}

type State struct {
	point  Point
	dir    Point
	length int
}

func (s *State) Equals(other *State) bool {
	return s.point.Equals(other.point) &&
		s.dir.Equals(other.dir) &&
		s.length == other.length
}

func (p Point) Equals(other Point) bool {
	return p.c == other.c && p.r == other.r
}

func runAny(min int, max int) {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]int{}
	for _, line := range lines {
		l := []int{}
		for _, c := range line {
			i, _ := strconv.Atoi(string(c))
			l = append(l, i)
		}
		grid = append(grid, l)
	}
	//fmt.Println(grid)
	countMin(grid, Point{0, 0}, Point{len(grid[0]) - 1, len(grid) - 1}, min, max)
}

func getDirs(dir Point) []Point {
	if dir == LEFT || dir == RIGHT {
		return []Point{dir, UP, DOWN}
	} else {
		return []Point{dir, LEFT, RIGHT}
	}
}
func countMin(grid [][]int, start Point, end Point, min int, max int) int {
	minHeat := 10000000
	visited := make(map[State]int)
	rows := len(grid)
	cols := len(grid[0])
	start1 := State{start, RIGHT, 1}
	start2 := State{start, DOWN, 1}
	queue := []State{start1, start2}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curHeat, has := visited[cur]
		if !has {
			fmt.Println("not found for", cur)
		}
		if cur.point == end && curHeat < minHeat {
			fmt.Println("new min", minHeat, curHeat, cur)
			minHeat = curHeat
		}

		for _, dir := range getDirs(cur.dir) {
			newPoint := Point{cur.point.c + dir.c, cur.point.r + dir.r}
			length := 1
			if dir == cur.dir {
				length = cur.length + 1
			}
			nextState := State{newPoint, dir, length}
			if newPoint.c >= 0 && newPoint.c < cols && newPoint.r >= 0 && newPoint.r < rows && length <= max && (dir == cur.dir || cur.length >= min) {
				nextHeat := curHeat + grid[newPoint.r][newPoint.c]
				nextHeatExists, nextVisited := visited[nextState]
				if nextVisited {
					if nextHeatExists > nextHeat {
						visited[nextState] = nextHeat
						queue = append(queue, nextState)
					}
				} else {
					visited[nextState] = nextHeat
					queue = append(queue, nextState)
				}
			}
		}
	}
	fmt.Println(minHeat)
	return minHeat
}
