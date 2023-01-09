package main

import (
	"adventofcode/util"
	"strings"
)

var DAY = "3"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2022/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, line := range lines {
		sum += getItem(line)
	}
	println(sum)
}

func runGold() {
	lines, err := util.ReadFile("year2022/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	ans := 0
	for i := 0; i < len(lines); i = i + 3 {
		ans += getItemGold(lines[i], lines[i+1], lines[i+2])
	}
	println(ans)
}

func getItem(input string) int {
	size := len(input) / 2
	partOne := input[:size]
	partTwo := input[size:]
	for _, cur := range partOne {
		if strings.Contains(partTwo, string(cur)) {
			return charToInt(cur)
		}
	}
	return 0
}

func getItemGold(partOne, partTwo, partThree string) int {
	for _, cur := range partOne {
		if strings.Contains(partTwo, string(cur)) && strings.Contains(partThree, string(cur)) {
			return charToInt(cur)
		}
	}
	return 0
}

func charToInt(c rune) int {
	if int(c) > 90 {
		return int(c) - int('a') + 1
	}
	return int(c) - int('A') + 27
}
