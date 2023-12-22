
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	right = iota
	down
	left
	up
)

type Instruction struct {
	direction int
	distance  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inst, ok := parseHexInstruction(line)
		if !ok {
			continue
		}
		instructions = append(instructions, inst)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	lavaVolume := calculateLavaVolume(instructions)
	fmt.Println(lavaVolume)
}

func parseHexInstruction(input string) (Instruction, bool) {
	var hexStr string
	_, err := fmt.Sscanf(input, "%*s %*s (#%s)", &hexStr)
	if err != nil {
		return Instruction{}, false
	}

	if len(hexStr) != 6 {
		return Instruction{}, false
	}

	distStr := hexStr[:5]
	dirStr := hexStr[5]

	dist, err := strconv.ParseInt(distStr, 16, 64)
	if err != nil {
		return Instruction{}, false
	}

	dir, err := strconv.ParseInt(string(dirStr), 16, 64)
	if err != nil {
		return Instruction{}, false
	}

	return Instruction{direction: int(dir), distance: int(dist)}, true
}

func calculateLavaVolume(instructions []Instruction) int64 {
	var x, y, area, boundaryPoints, cornerPoints int64

	visit := map[int64]map[int64]bool{}

	for _, inst := range instructions {
		for i := 0; i < inst.distance; i++ {
			if visit[y] == nil {
				visit[y] = map[int64]bool{}
			}

			if !visit[y][x] {
				visit[y][x] = true
				boundaryPoints++
				if len(visit[y-1]) > 0 && visit[y-1][x] {
					boundaryPoints--
				}
				if visit[y][x-1] {
					boundaryPoints--
				}
			}

			switch inst.direction {
			case right:
				x++
			case down:
				y++
			case left:
				x--
			case up:
				y--
			}

			if visit[y] == nil {
				visit[y] = map[int64]bool{}
			}
			if !visit[y][x] {
				visit[y][x] = true
				boundaryPoints++
				if len(visit[y-1]) > 0 && visit[y-1][x] {
					boundaryPoints--
				}
				if visit[y][x-1] {
					boundaryPoints--
				}
				if visit[y][x+1] {
					boundaryPoints--
				}
				if visit[y+1][x] {
					boundaryPoints--
				}
			}
		}
		cornerPoints++
	}

	area = (boundaryPoints - cornerPoints) / 2
	return area
}

