package util

//const (
//	LEFT       = Point{-1, 0}
//	LEFT_UP    = Point{-1, -1}
//	LEFT_DOWN  = Point{-1, 1}
//	RIGHT      = Point{1, 0}
//	RIGHT_UP   = Point{1, -1}
//	RIGHT_DOWN = Point{1, 1}
//	UP         = Point{0, -1}
//	DOWN       = Point{0, 1}
//)

var LEFT = Point{-1, 0}
var LEFT_UP = Point{-1, -1}
var LEFT_DOWN = Point{-1, 1}
var RIGHT = Point{1, 0}
var RIGHT_UP = Point{1, -1}
var RIGHT_DOWN = Point{1, 1}
var UP = Point{0, -1}
var DOWN = Point{0, 1}

var DIRECTIONS_ALL = []Point{LEFT, LEFT_UP, LEFT_DOWN, RIGHT, RIGHT_UP, RIGHT_DOWN, UP, DOWN}
var DIRECTIONS_CROS = []Point{LEFT, RIGHT, UP, DOWN}

type Point struct {
	C, R int
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

func NewGrid(lines []string) *Grid {
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return &Grid{
		Grid: grid,
	}
}

func (g *Grid) AddRaw(line string) {
	g.Grid = append(g.Grid, []rune(line))

}

func (g *Grid) Print() {
	for _, line := range g.Grid {
		println(string(line))
	}
}

func (g *Grid) IsValid(r, c int) bool {
	return r >= 0 && c >= 0 && r < len(g.Grid) && c < len(g.Grid[0])
}

func (g *Grid) At(r, c int) rune {
	return g.Grid[r][c]
}
