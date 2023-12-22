
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NumSteps  = 64
	Direction = 4
)

type Pos struct {
	x, y int
}

var (
	dx             = [Direction]int{0, 1, 0, -1}
	dy             = [Direction]int{-1, 0, 1, 0}
	startPos       Pos
	gardens        [][]bool
	width, height  int
	possiblePlots  = make(map[Pos]int)
)

func readInput(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if width == 0 {
			width = len(line)
		}
		height++
		gardenRow := make([]bool, width)
		for i, ch := range line {
			if ch == '.' || ch == 'S' {
				gardenRow[i] = true
				if ch == 'S' {
					startPos = Pos{x: i, y: height - 1}
				}
			}
		}
		gardens = append(gardens, gardenRow)
	}

	return scanner.Err()
}

func stepCounter(current Pos, steps int) {
	if steps == NumSteps {
		possiblePlots[current]++
		return
	}
	for i := 0; i < Direction; i++ {
		next := Pos{x: current.x + dx[i], y: current.y + dy[i]}
		if next.x >= 0 && next.x < width && next.y >= 0 && next.y < height && gardens[next.y][next.x] {
			stepCounter(next, steps+1)
		}
	}
}

func main() {
	err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading the input file:", err)
		return
	}

	stepCounter(startPos, 0)
	fmt.Println(len(possiblePlots))
}
