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
	r, c int
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
	return fmt.Sprintf("(%d,%d)", p.r, p.c)
}

func NewPoint(r, c int) Point {
	return Point{
		r: r,
		c: c,
	}
}

func (p Point) AddDirection(direction Direction) Point {
	if dir, exists := Directions[direction]; exists {
		return Point{r: p.r + dir.r, c: p.c + dir.c}
	}
	return p
}

func (p Point) R() int {
	return p.r
}

func (p Point) C() int {
	return p.c
}
func (p Point) AddPointInBorder(point Point, maxR, maxC int) Point {
	return Point{
		r: (p.r + point.r + maxR) % maxR,
		c: (p.c + point.c + maxC) % maxC,
	}
}

func (p Point) AddPoint(point Point) Point {
	return Point{
		r: p.r + point.r,
		c: p.c + point.c,
	}
}

func (p Point) MinusPoint(point Point) Point {
	return Point{
		r: p.r - point.r,
		c: p.c - point.c,
	}
}

func (p Point) EqualPoint(point Point) bool {
	return p.r == point.r && p.c == point.c
}

func (p Point) TimesPoint(times int) Point {
	return Point{r: p.r * times, c: p.c * times}
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
func (g *Grid) PrintDebugWithDotsSymbol(visited map[Point]bool, symbol rune) {
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
		println()
		for j, char := range line {
			if visited[Point{i, j}] {
				print(string(symbol))
			} else {
				print(string(char))
			}
		}
	}
	fmt.Println()

}

func (g *Grid) IsValidPoint(p Point) bool {
	return g.IsValid(p.r, p.c)
}

func (g *Grid) IsValid(r, c int) bool {
	return r >= 0 && c >= 0 && r < len(g.Grid) && c < len(g.Grid[0])
}

func (g *Grid) At(r, c int) rune {
	return g.Grid[r][c]
}

func (g *Grid) AtPoint(p Point) rune {
	return g.At(p.r, p.c)
}

func (g *Grid) Neighbors(p Point) []Point {
	var neighbors []Point
	for _, dir := range DIRECTIONS_ALL {
		neighbor := Point{r: p.r + dir.r, c: p.c + dir.c}
		if g.IsValidPoint(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}
