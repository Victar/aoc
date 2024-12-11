package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "11"

type Stone struct {
	number, blink int
}

var DP = make(map[Stone]int)

func main() {
	runBoth()
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	input := lines[0]
	inputStr := strings.TrimSpace(string(input))
	stonesStr := strings.Split(inputStr, " ")
	stones := make([]int, len(stonesStr))
	for i, str := range stonesStr {
		stones[i], _ = strconv.Atoi(str)
	}
	fmt.Println(countAllStones(stones, 25))
	fmt.Println(countAllStones(stones, 75))
}

func countAllStones(numbers []int, blink int) int {
	sum := 0
	for _, number := range numbers {
		sum += countStones(Stone{number, blink})
	}
	return sum
}

func countStones(stone Stone) int {
	if c, exists := DP[stone]; exists {
		return c
	}
	var count int
	if stone.blink == 0 {
		count = 1
	} else if stone.number == 0 {
		count = countStones(Stone{1, stone.blink - 1})
	} else if len(strconv.Itoa(stone.number))%2 == 0 {
		left, right := splitDigits(stone.number)
		count = countStones(Stone{left, stone.blink - 1}) + countStones(Stone{right, stone.blink - 1})
	} else {
		count = countStones(Stone{stone.number * 2024, stone.blink - 1})
	}
	DP[stone] = count
	return count
}

func splitDigits(num int) (int, int) {
	digitsStr := strconv.Itoa(num)
	mid := len(digitsStr) / 2
	left, _ := strconv.Atoi(digitsStr[:mid])
	right, _ := strconv.Atoi(digitsStr[mid:])
	return left, right
}
