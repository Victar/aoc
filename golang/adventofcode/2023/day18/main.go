package main

import (
	"adventofcode/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var DAY = "18"

// Point represents coordinates x, y in a 2D grid
type Point struct {
	x, y int
}

// Bounds holds the minimum and maximum values of x and y
type Bounds struct {
	minX, maxX, minY, maxY int
}

// Size returns the width and height of the area within the bounds.
func (b Bounds) Size() (int, int) {
	return b.maxX - b.minX + 1, b.maxY - b.minY + 1
}

func main() {
	runSilver()
	runGold()
}
func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	points := make([]Point, 0, len(lines))
	points = append(points, Point{0, 0})
	start := Point{0, 0}
	points = append(points)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		nums := parts[2][2:7]
		dir := parts[2][7:8]
		n, _ := strconv.ParseInt(nums, 16, 0)
		switch dir {
		case "0":
			start.x += int(n)
		case "1":
			start.y += int(n)
		case "2":
			start.x -= int(n)
		case "3":
			start.y -= int(n)
		}
		points = append(points, start)
	}
	var ans = shoelaceManhattanFormula(points) + 1
	fmt.Println(ans)
}

func shoelaceManhattanFormula(points []Point) int {
	n := len(points)
	if n < 3 {
		return 0
	}
	area := 0
	for i := 0; i < n; i++ {
		point, pointNext := points[i], points[(i+1)%n]
		shoelace := (point.x * pointNext.y) - (pointNext.x * point.y)
		manhattan := abs(point.x-pointNext.x) + abs(point.y-pointNext.y)
		area += shoelace + manhattan
	}
	area = area / 2
	if area < 0 {
		area = -area
	}
	return area
}

func abs(c int) int {
	if c < 0 {
		return -c
	}
	return c
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var bounds Bounds
	bounds.minX, bounds.minY = 0, 0
	bounds.maxX, bounds.maxY = 0, 0
	pos := Point{0, 0}
	areaBorder := make(map[Point]bool)
	areaBorder[pos] = true

	for _, line := range lines {
		r := regexp.MustCompile(`([UDLR]) (\d+)`)
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			dir := match[1]
			distance, _ := strconv.Atoi(match[2])
			dx, dy := 0, 0
			switch dir {
			case "U":
				dy = -1
			case "D":
				dy = 1
			case "L":
				dx = -1
			case "R":
				dx = 1
			}
			for i := 0; i < distance; i++ {
				pos.x += dx
				pos.y += dy
				areaBorder[pos] = true
				if pos.x < bounds.minX {
					bounds.minX = pos.x
				}
				if pos.x > bounds.maxX {
					bounds.maxX = pos.x
				}
				if pos.y < bounds.minY {
					bounds.minY = pos.y
				}
				if pos.y > bounds.maxY {
					bounds.maxY = pos.y
				}
			}
		}
	}

	areaFull := isInside(bounds, areaBorder)
	fmt.Println(len(areaFull))

}

func isInside(bounds Bounds, border map[Point]bool) map[Point]bool {
	fullArea := make(map[Point]bool)
	globalVisited := make(map[Point]bool)
	dirs := []Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for y := bounds.minY; y <= bounds.maxY; y++ {
		for x := bounds.minX; x <= bounds.maxX; x++ {
			start := Point{x, y}
			if border[start] {
				fullArea[start] = true
			} else if !globalVisited[start] {
				queue := []Point{start}
				visited := make(map[Point]bool)
				inArea := true
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					visited[cur] = true
					if cur.x == bounds.minX || cur.x == bounds.maxX || cur.y == bounds.minY || cur.y == bounds.maxY {
						inArea = false
						break
					}
					for _, dir := range dirs {
						newPoint := Point{cur.x + dir.x, cur.y + dir.y}
						if newPoint.x >= bounds.minX && newPoint.x <= bounds.maxX && newPoint.y >= bounds.minY && newPoint.y <= bounds.maxY && !border[newPoint] && !visited[newPoint] {
							queue = append(queue, newPoint)
							visited[newPoint] = true
						}
					}
				}
				if inArea {
					fullArea[start] = true
					for k, _ := range visited {
						fullArea[k] = true
					}
				}
				for k, _ := range visited {
					globalVisited[k] = true
				}
			}
		}
	}
	return fullArea
}

func drawArea(bounds Bounds, visited map[Point]bool) {
	for y := bounds.minY; y <= bounds.maxY; y++ {
		line := strconv.Itoa(y) + " "
		for x := bounds.minX; x <= bounds.maxX; x++ {
			if visited[Point{x, y}] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}
