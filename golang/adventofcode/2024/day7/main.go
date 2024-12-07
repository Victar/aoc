package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

var DAY = "7"

func main() {
	runAny(false)
	runAny(true)
}

func runAny(isGold bool) {

	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}

	res := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(parts[0])
		numberStrings := strings.Split(parts[1], " ")
		numbers := make([]int, len(numberStrings))
		for i, ns := range numberStrings {
			num, _ := strconv.Atoi(ns)
			numbers[i] = num
		}
		if canBeTrue(numbers, testValue, 0, isGold) {
			res = res + testValue
		}
	}
	println(res)
}

func canBeTrue(numbers []int, targetValue int, curValue int, isGold bool) bool {

	if len(numbers) == 0 {
		return curValue == targetValue
	}

	curValueAdd := curValue + numbers[0]
	curValueMult := curValue * numbers[0]
	curValueCombine, _ := strconv.Atoi(strconv.Itoa(curValue) + strconv.Itoa(numbers[0]))

	return canBeTrue(numbers[1:], targetValue, curValueAdd, isGold) ||
		canBeTrue(numbers[1:], targetValue, curValueMult, isGold) ||
		(canBeTrue(numbers[1:], targetValue, curValueCombine, isGold) && isGold)

}
