package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

var DAY = "2"

func main() {
	runSilver()
	runGold()
}

var ballCount = map[string]int{
	"red": 12, "green": 13, "blue": 14,
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	ans := 0
	for _, line := range lines {
		gamePossible := true
		parts := strings.Split(line, ":")
		gameId, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		games := strings.Split(parts[1], ";")
		for _, game := range games {
			balls := strings.Split(game, ",")
			for _, ball := range balls {
				ballParts := strings.Split(strings.TrimSpace(ball), " ")
				count, _ := strconv.Atoi(ballParts[0])
				color := ballParts[1]
				limit, ok := ballCount[color]
				if !(ok && limit >= count) {
					gamePossible = false
					break
				}
			}
		}
		if gamePossible {
			ans += gameId
		}
	}
	println(ans)
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	ans := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		games := strings.Split(parts[1], ";")
		ballCount = map[string]int{
			"red": 1, "green": 1, "blue": 1,
		}
		for _, game := range games {
			balls := strings.Split(game, ",")
			for _, ball := range balls {
				ballParts := strings.Split(strings.TrimSpace(ball), " ")
				count, _ := strconv.Atoi(ballParts[0])
				color := ballParts[1]
				limit, ok := ballCount[color]
				if !ok || limit < count {
					ballCount[color] = count
				}
			}
		}
		ans += ballCount["red"] * ballCount["green"] * ballCount["blue"]
	}
	println(ans)
}
