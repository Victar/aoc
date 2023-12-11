
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pipes := make(map[[2]int]rune)
	start := [2]int{}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, ch := range line {
			if ch == '.' || ch == ' ' {
				continue
			}
			pipes[[2]int{x, y}] = ch
			if ch == 'S' {
				start = [2]int{x, y}
			}
		}
	}

	maxDist := 0

	var dfs func(node, parent [2]int, distance int)
	dfs = func(node, parent [2]int, distance int) {
		if distance > maxDist {
			maxDist = distance
		}
		for _, d := range [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			nextNode := [2]int{node[0] + d[0], node[1] + d[1]}
			if nextNode == parent {
				continue
			}
			nextPipe, ok := pipes[nextNode]
			if !ok {
				continue
			}
			if connects(node, nextNode, nextPipe) {
				dfs(nextNode, node, distance+1)
			}
		}
	}
	dfs(start, [2]int{}, 0)

	fmt.Println(maxDist)
}

func connects(from, to [2]int, nextPipe rune) bool {
	dx, dy := to[0]-from[0], to[1]-from[1]

	switch {
	case dx == 1 && (nextPipe == '-' || nextPipe == 'J' || nextPipe == '7' || nextPipe == 'S'):
		return true
	case dx == -1 && (nextPipe == '-' || nextPipe == 'L' || nextPipe == 'F' || nextPipe == 'S'):
		return true
	case dy == 1 && (nextPipe == '|' || nextPipe == 'L' || nextPipe == 'J' || nextPipe == 'S'):
		return true
	case dy == -1 && (nextPipe == '|' || nextPipe == 'F' || nextPipe == '7' || nextPipe == 'S'):
		return true
	}
	return false
}
