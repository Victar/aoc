package main

import (
	"adventofcode/util"
	"image"
)

func main() {
	runSilver()
	runGold()

}
func runSilver() {
	ar, err := util.ReadFile("year2015/day3/input.txt")
	if err != nil {
		panic(err)
	}
	visited := make(map[image.Point]bool)
	var curX, curY = 0, 0
	for _, c := range ar[0] {
		//>v<>^
		if c == '>' {
			curX++
		}
		if c == '<' {
			curX--
		}
		if c == 'v' {
			curY++
		}
		if c == '^' {
			curY--
		}
		visited[image.Point{curX, curY}] = true
	}
	println(len(visited))
}
func runGold() {
	ar, err := util.ReadFile("year2015/day3/input.txt")
	if err != nil {
		panic(err)
	}
	visited := make(map[image.Point]bool)
	var curX, curY, curXR, curYR = 0, 0, 0, 0
	for i, c := range ar[0] {
		if i%2 == 0 {
			if c == '>' {
				curX++
			}
			if c == '<' {
				curX--
			}
			if c == 'v' {
				curY++
			}
			if c == '^' {
				curY--
			}
			visited[image.Point{curX, curY}] = true
		} else {
			if c == '>' {
				curXR++
			}
			if c == '<' {
				curXR--
			}
			if c == 'v' {
				curYR++
			}
			if c == '^' {
				curYR--
			}
			visited[image.Point{curXR, curYR}] = true
		}
	}
	println(len(visited))
}
