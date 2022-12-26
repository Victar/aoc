package main

import (
	"adventofcode/util"
)

var DAY = "X"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2022/day/" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}

func runGold() {
	lines, err := util.ReadFile("year2022/day/" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
