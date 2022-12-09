package main

import (
	"adventofcode/util"
	"fmt"
	"image"
	"math"
)

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	runAny(1)
}

func runGold() {
	runAny(9)
}

func runAny(size int) {
	lines, err := util.ReadFile("year2022/day9/input.txt")
	if err != nil {
		panic(err)
	}
	moves := map[string]image.Point{"U": {0, -1}, "D": {0, 1}, "L": {-1, 0}, "R": {1, 0}}
	tail := make([]image.Point, size+1)
	var visited = make(map[image.Point]image.Point)
	for _, s := range lines {
		var move string
		var length int
		_, err := fmt.Sscanf(s, "%s %d", &move, &length)
		if err != nil {
			panic(err)
		}

		for j := 0; j < length; j++ {
			tail[0] = tail[0].Add(moves[move])
			for i := 1; i < len(tail); i++ {
				tail[i] = getNewTail(tail[i-1], tail[i])
			}
			visited[tail[size]] = tail[size]
		}
	}
	fmt.Println(len(visited))
}
func getNewTail(head image.Point, t image.Point) image.Point {
	var tail = image.Point{X: t.X, Y: t.Y}
	distX := head.X - tail.X
	distY := head.Y - tail.Y
	if math.Abs(float64(distX)) > 1 || math.Abs(float64(distY)) > 1 {
		if math.Abs(float64(distX)) != 0 {
			if distX > 0 {
				tail.X = tail.X + 1
			} else {
				tail.X = tail.X - 1
			}
		}
		if math.Abs(float64(distY)) != 0 {
			if distY > 0 {
				tail.Y = tail.Y + 1
			} else {
				tail.Y = tail.Y - 1
			}
		}
	}
	return tail
}
