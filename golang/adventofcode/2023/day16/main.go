package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "16"

type Point struct {
	x, y int
}

type Direction int

const (
	Right Direction = iota
	Left
	Up
	Down
)

type PointMove struct {
	Point     Point
	Direction Direction
}

func main() {
	runSilver()
	runGold()
}

func printGrid(grid []string, visited map[Point]bool) {
	cols := len(grid[0])
	rows := len(grid)
	for x := 0; x < rows; x++ {
		line := ""
		for y := 0; y < cols; y++ {
			if visited[Point{x, y}] {
				line += "#"
			} else {
				line += "."
			}

		}
		fmt.Println(line)
	}
}
func runSilver() {
	grid, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	startPoint := PointMove{Point: Point{0, 0}, Direction: Right}
	energizedTiles := simulateLightBeam(grid, startPoint)
	println(len(energizedTiles))
}

func runGold() {
	grid, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	max := 0
	cols := len(grid[0])
	rows := len(grid)
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for i := 0; i < 4; i++ {
				if (x == 0 || x == rows-1) || (y == 0 || y == cols-1) {
					startPoint := PointMove{Point: Point{x, y}, Direction: Direction(i)}
					energizedTiles := simulateLightBeam(grid, startPoint)
					cur := len(energizedTiles)
					if cur > max {
						max = cur
					}
				}
			}
		}
	}
	println(max)
}

func simulateLightBeam(grid []string, startPoint PointMove) map[Point]bool {
	height := len(grid)
	width := len(grid[0])
	energized := make(map[Point]bool)
	directions := map[Direction][]int{
		Right: {0, 1},
		Left:  {0, -1},
		Up:    {-1, 0},
		Down:  {1, 0},
	}
	mirrorDirection := map[rune]map[Direction]Direction{
		'/': {
			Right: Up,
			Down:  Left,
			Left:  Down,
			Up:    Right,
		},
		'\\': {
			Right: Down,
			Up:    Left,
			Left:  Up,
			Down:  Right,
		},
	}
	visited := make(map[PointMove]bool)
	queueMove := []PointMove{startPoint}
	visited[startPoint] = true

	for len(queueMove) > 0 {
		p := queueMove[0]
		queueMove = queueMove[1:]
		currentDir := p.Direction
		split := false
		for {
			if p.Point.x < 0 || p.Point.x >= height || p.Point.y < 0 || p.Point.y >= width {
				break
			}

			// Mark the current tile as energized
			energized[p.Point] = true
			if split {
				break
			}
			tile := rune(grid[p.Point.x][p.Point.y])
			switch tile {
			case '.':
				dx, dy := directions[currentDir][0], directions[currentDir][1]
				p.Point.x += dx
				p.Point.y += dy
			case '|':
				if currentDir == Left || currentDir == Right {
					down := PointMove{Point: Point{p.Point.x + 1, p.Point.y}, Direction: Down}
					up := PointMove{Point: Point{p.Point.x - 1, p.Point.y}, Direction: Up}
					if !visited[down] {
						queueMove = append(queueMove, down)
						visited[down] = true
					}
					if !visited[up] {
						queueMove = append(queueMove, up)
						visited[up] = true
					}
					split = true
					break
				} else {
					dx, dy := directions[currentDir][0], directions[currentDir][1]
					p.Point.x += dx
					p.Point.y += dy
					break
				}
			case '-':
				if currentDir == Up || currentDir == Down {
					right := PointMove{Point: Point{p.Point.x, p.Point.y + 1}, Direction: Right}
					left := PointMove{Point: Point{p.Point.x, p.Point.y - 1}, Direction: Left}
					if !visited[right] {
						queueMove = append(queueMove, right)
						visited[right] = true
					}
					if !visited[left] {
						queueMove = append(queueMove, left)
						visited[left] = true
					}
					split = true
				} else {
					dx, dy := directions[currentDir][0], directions[currentDir][1]
					p.Point.x += dx
					p.Point.y += dy
					break
				}
			case '/', '\\':
				currentDir = mirrorDirection[tile][currentDir]
				dx, dy := directions[currentDir][0], directions[currentDir][1]
				p.Point.x += dx
				p.Point.y += dy
			default:
				break
			}
		}
	}
	return energized
}
