package util

import (
	"fmt"
	"strconv"
	"strings"
)

type Direction string

const (
	LEFT       Direction = "LEFT"
	LEFT_UP    Direction = "LEFT_UP"
	LEFT_DOWN  Direction = "LEFT_DOWN"
	RIGHT      Direction = "RIGHT"
	RIGHT_UP   Direction = "RIGHT_UP"
	RIGHT_DOWN Direction = "RIGHT_DOWN"
	UP         Direction = "UP"
	DOWN       Direction = "DOWN"
)

var Directions = map[Direction]Point{
	LEFT:       {0, -1},
	LEFT_UP:    {-1, -1},
	LEFT_DOWN:  {1, -1},
	RIGHT:      {0, 1},
	RIGHT_UP:   {-1, 1},
	RIGHT_DOWN: {1, 1},
	UP:         {-1, 0},
	DOWN:       {1, 0},
}

var DIRECTIONS_ALL = []Point{
	{-1, 0}, {-1, -1}, {-1, 1}, {1, 0}, {1, -1}, {1, 1}, {0, -1}, {0, 1},
}

type Point struct {
	R, C int
}

type Grid struct {
	Grid [][]rune
}

func NewGridEmpty() *Grid {
	var grid [][]rune
	return &Grid{
		Grid: grid,
	}
}

func (g *Grid) BFS(start Point, end Point) (bool, []Point) {
	visited := make(map[Point]bool)
	visitedParent := make(map[Point]*Point)
	queue := []Point{start}
	visited[start] = true
	var directions = []Point{Directions[RIGHT], Directions[DOWN], Directions[LEFT], Directions[UP]}
	for len(queue) > 0 {
		curPoint := queue[0]
		queue = queue[1:]
		if curPoint == end {
			var path []Point
			for p := &curPoint; p != nil; p = visitedParent[*p] {
				path = append([]Point{*p}, path...)
			}
			return true, path
		}
		for _, dir := range directions {
			nextPoint := curPoint.AddPoint(dir)
			if g.IsValidPoint(nextPoint) && !visited[nextPoint] && g.AtPoint(nextPoint) == '.' {
				visited[nextPoint] = true
				visitedParent[nextPoint] = &curPoint
				queue = append(queue, nextPoint)
			}
		}
	}
	return false, nil
}

func (g *Grid) Copy() *Grid {
	if g == nil {
		return nil
	}

	newGrid := make([][]rune, len(g.Grid))
	for i := range g.Grid {
		newGrid[i] = make([]rune, len(g.Grid[i]))
		copy(newGrid[i], g.Grid[i])
	}
	return &Grid{
		Grid: newGrid,
	}
}

func (g *Grid) RowColSize() (int, int) {
	if g == nil || len(g.Grid) == 0 {
		return 0, 0
	}
	return len(g.Grid), len(g.Grid[0])
}

func (g *Grid) CollectPoints(char rune) map[Point]bool {
	result := map[Point]bool{}
	for r, row := range g.Grid {
		for c, cell := range row {
			if cell == char {
				result[NewPoint(r, c)] = true
			}
		}
	}
	return result
}

func (g *Grid) SetRune(r, c int, char rune) {
	if g.IsValid(r, c) {
		g.Grid[r][c] = char
	} else {
		panic(fmt.Sprintf("(%d,%d)", r, c))
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.R, p.C)
}

func NewPoint(r, c int) Point {
	return Point{
		R: r,
		C: c,
	}
}

func (p Point) AddDirection(direction Direction) Point {
	if dir, exists := Directions[direction]; exists {
		return Point{R: p.R + dir.R, C: p.C + dir.C}
	}
	return p
}

func (p Point) AddPointInBorder(point Point, maxR, maxC int) Point {
	return Point{
		R: (p.R + point.R + maxR) % maxR,
		C: (p.C + point.C + maxC) % maxC,
	}
}

func (p Point) AddPoint(point Point) Point {
	return Point{
		R: p.R + point.R,
		C: p.C + point.C,
	}
}

func (p Point) MinusPoint(point Point) Point {
	return Point{
		R: p.R - point.R,
		C: p.C - point.C,
	}
}

func (p Point) EqualPoint(point Point) bool {
	return p.R == point.R && p.C == point.C
}

func (p Point) TimesPoint(times int) Point {
	return Point{R: p.R * times, C: p.C * times}
}

func NewGrid(lines []string) *Grid {
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return &Grid{
		Grid: grid,
	}
}

func (g *Grid) AddRow(line string) {
	g.Grid = append(g.Grid, []rune(line))

}

func (g *Grid) Print() {
	for _, line := range g.Grid {
		println(string(line))
	}
}

func (g *Grid) PrintDebug() {
	if len(g.Grid) == 0 {
		fmt.Println("Grid is empty")
		return
	}
	fmt.Printf("Grid row %d columns %d \n", len(g.Grid), len(g.Grid[0]))
	columns := len(g.Grid[0])
	if columns > 10 {
		firstLine := ""
		for i := 0; i < 10; i++ {
			firstLine += strings.Repeat(strconv.Itoa(i), 10)
		}
		println(strings.Repeat(firstLine, columns/100+1)[:columns])
	}
	println(strings.Repeat("0123456789", columns/10+1)[:columns])

	for i, line := range g.Grid {
		println(string(line), i)
	}
	fmt.Println()

}

func (g *Grid) PrintDebugWithDots(visited map[Point]bool) {
	g.PrintDebugWithDotsSymbol(visited, 'X')
}

func (g *Grid) PrintDebugVisitedOnly(visited map[Point]bool, symbol rune) {
	g.printDebugAny(visited, symbol, true)
}

func (g *Grid) PrintDebugWithDotsSymbol(visited map[Point]bool, symbol rune) {
	g.printDebugAny(visited, symbol, false)
}
func (g *Grid) printDebugAny(visited map[Point]bool, symbol rune, visitedOnly bool) {
	if len(g.Grid) == 0 {
		fmt.Println("Grid is empty")
		return
	}
	fmt.Printf("Grid row %d columns %d \n", len(g.Grid), len(g.Grid[0]))
	columns := len(g.Grid[0])
	if columns > 10 {
		firstLine := ""
		for i := 0; i < 10; i++ {
			firstLine += strings.Repeat(strconv.Itoa(i), 10)
		}
		println(strings.Repeat(firstLine, columns/100+1)[:columns])
	}
	println(strings.Repeat("0123456789", columns/10+1)[:columns])
	sb := strings.Builder{}
	for i, line := range g.Grid {
		for j, char := range line {
			if visited[Point{i, j}] {
				sb.WriteByte(byte(symbol))
			} else {
				if visitedOnly {
					sb.WriteByte(' ')
				} else {
					sb.WriteByte(byte(char))
				}
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func (g *Grid) IsValidPoint(p Point) bool {
	return g.IsValid(p.R, p.C)
}

func ConvertSliceToMap(points []Point) map[Point]bool {
	pointMap := make(map[Point]bool)
	for _, point := range points {
		pointMap[point] = true
	}
	return pointMap
}

func (g *Grid) IsValid(r, c int) bool {
	return r >= 0 && c >= 0 && r < len(g.Grid) && c < len(g.Grid[0])
}

func (g *Grid) At(r, c int) rune {
	return g.Grid[r][c]
}

func (g *Grid) AtPoint(p Point) rune {
	return g.At(p.R, p.C)
}

func (g *Grid) Neighbors(p Point) []Point {
	var neighbors []Point
	for _, dir := range DIRECTIONS_ALL {
		neighbor := Point{R: p.R + dir.R, C: p.C + dir.C}
		if g.IsValidPoint(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}
