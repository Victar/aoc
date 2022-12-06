package main

import (
	"adventofcode/util"
)

func main() {
	runSilver()
	runGold()
}

func runAny(count int) {
	lines, err := util.ReadFile("year2022/day6/input.txt")
	if err != nil {
		panic(err)
	}
	var input = lines[0]
	for i := 0; i < len(input); i++ {
		var un = make(map[byte]byte)
		for j := 0; j < count; j++ {
			un[input[i+j]] = input[i+j]
		}
		if len(un) == count {
			println(i + count)
			break
		}
	}
}

func runSilver() {
	runAny(4)
}

func runGold() {
	runAny(14)
}
