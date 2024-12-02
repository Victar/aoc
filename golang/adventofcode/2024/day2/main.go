package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

var DAY = "2"

func main() {
	runAny(false)
	runAny(true)
}

func runAny(isGold bool) {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	count := 0
	for _, line := range lines {
		numbersStr := strings.Fields(line)
		numbers := make([]int, len(numbersStr))
		for i, n := range numbersStr {
			numbers[i], _ = strconv.Atoi(n)
		}
		if len(numbers) > 0 && isSafeExtended(numbers, isGold) {
			count++
		}
	}
	println(count)
}

func isSafeExtended(report []int, useDampener bool) bool {
	isSafeInitial := isSafe(report)
	if !useDampener {
		return isSafeInitial
	}
	for i := range report {
		modifiedReport := append([]int(nil), report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		if isSafe(modifiedReport) {
			return true
		}
	}
	return false
}

func isSafe(report []int) bool {
	increasing, decreasing := true, true
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if diff > 0 {
			increasing = false
		} else {
			decreasing = false
		}
	}
	return increasing || decreasing
}
