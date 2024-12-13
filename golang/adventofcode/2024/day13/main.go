package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "13"

func main() {
	runAny(false)
	runAny(true)
}

func runAny(isGold bool) {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var minTokens int
	for i := 0; i < len(lines); i = i + 4 {
		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]
		aX, aY := parseButton(line1)
		bX, bY := parseButton(line2)
		prizeX, prizeY := parsePrize(line3)
		if isGold {
			prizeX, prizeY = prizeX+10000000000000, prizeY+10000000000000
		}
		x, y := solveUsingCramersRule(aX, aY, bX, bY, prizeX, prizeY)
		if x >= 0 && y >= 0 {
			minTokens += 3*x + y
		}
	}
	fmt.Println(minTokens)
}

// https://en.wikipedia.org/wiki/Cramer%27s_rule
func solveUsingCramersRule(aX, aY, bX, bY, prizeX, prizeY int) (int, int) {
	det := aX*bY - aY*bX
	detX := prizeX*bY - prizeY*bX
	detY := prizeY*aX - prizeX*aY
	if det != 0 && detX%det == 0 && detY%det == 0 {
		x := detX / det
		y := detY / det
		return x, y
	}
	return -1, -1
}

func parsePrize(line string) (int, int) {
	parts := strings.Fields(line)
	x, _ := strconv.Atoi(strings.TrimSuffix(parts[1], ",")[2:])
	y, _ := strconv.Atoi(parts[2][2:])
	return x, y
}
func parseButton(line string) (int, int) {
	parts := strings.Fields(line)
	x, _ := strconv.Atoi(strings.TrimSuffix(parts[2], ",")[2:])
	y, _ := strconv.Atoi(parts[3][2:])
	return x, y
}
