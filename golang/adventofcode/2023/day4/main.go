package main

import (
	"adventofcode/util"
	"slices"
	"strings"
)

var DAY = "4"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	ans := 0
	for _, line := range lines {
		cardArr := strings.Split(line, ":")
		numbersArr := strings.Split(cardArr[1], "|")
		winNums := strings.Fields(strings.TrimSpace(numbersArr[0]))
		myNums := strings.Fields(strings.TrimSpace(numbersArr[1]))
		ans += getCardScore(myNums, winNums, false)
	}
	println(ans)
}

func getCardScore(myNums []string, winNums []string, isGold bool) int {
	ans := 0
	for _, myNumber := range myNums {
		if slices.Contains(winNums, myNumber) {
			if ans == 0 {
				ans = 1
			} else {
				if isGold {
					ans++
				} else {
					ans = 2 * ans
				}
			}
		}
	}
	return ans
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	ans := 0
	ansArr := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		ansArr[i] = 1
	}
	for cur, line := range lines {
		cardArr := strings.Split(line, ":")
		numbersArr := strings.Split(cardArr[1], "|")
		winNums := strings.Fields(strings.TrimSpace(numbersArr[0]))
		myNums := strings.Fields(strings.TrimSpace(numbersArr[1]))
		score := getCardScore(myNums, winNums, true)
		curCount := ansArr[cur]
		for i := cur + 1; i <= cur+score; i++ {
			ansArr[i] = ansArr[i] + curCount
		}
	}
	for i := 0; i < len(lines); i++ {
		ans += ansArr[i]
	}
	println(ans)
}
