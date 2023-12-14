package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "14"

func main() {
	runSilver()
	//runGold()
}

func runSilver() {
	rowCount := 0
	rows, err := util.ReadInput("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(rowCount, rows)
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
