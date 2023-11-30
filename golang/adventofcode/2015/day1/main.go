package main

import (
	"adventofcode/util"
)

func main() {
	runSilver()
	runGold()
}
func runSilver() {
	ar, err := util.ReadFile("year2015/day1/input.txt")
	if err != nil {
		panic(err)
	}
	var floor = 0
	for _, char := range ar[0] {
		if char == ')' {
			floor--
		}
		if char == '(' {
			floor++
		}
	}
	println(floor)
}

func runGold() {
	ar, err := util.ReadFile("year2015/day1/input.txt")
	if err != nil {
		panic(err)
	}
	var floor = 0
	for pos, char := range ar[0] {
		if char == ')' {
			floor--
		}
		if char == '(' {
			floor++
		}
		if floor == -1 {
			println(pos + 1)
			break
		}
	}
}
