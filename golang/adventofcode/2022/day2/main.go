package main

import (
	"adventofcode/util"
)

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2022/day2/input.txt")
	if err != nil {
		panic(err)
	}
	var score = 0
	for _, line := range lines {
		score += NewGame(line).countSilverScore()
	}
	println(score)
}

func runGold() {
	lines, err := util.ReadFile("year2022/day2/input.txt")
	if err != nil {
		panic(err)
	}
	var score = 0
	for _, line := range lines {
		score += NewGame(line).countGoldScore()
	}
	println(score)
}

type Game struct {
	p1 string
	p2 string
}

func NewGame(input string) Game {
	return Game{
		p1: string(input[0]),
		p2: string(input[2]),
	}
}

func (game Game) countSilverScore() int {
	if game.p1 == "A" {
		if game.p2 == "X" {
			return 3 + 1
		}
		if game.p2 == "Y" {
			return 6 + 2
		}
		if game.p2 == "Z" {
			return 0 + 3
		}
	}
	if game.p1 == "B" {
		if game.p2 == "X" {
			return 0 + 1
		}
		if game.p2 == "Y" {
			return 3 + 2
		}
		if game.p2 == "Z" {
			return 6 + 3
		}
	}
	if game.p1 == "C" {
		if game.p2 == "X" {
			return 6 + 1
		}
		if game.p2 == "Y" {
			return 0 + 2
		}
		if game.p2 == "Z" {
			return 3 + 3
		}
	}
	return 0
}

func (game Game) countGoldScore() int {
	if game.p2 == "X" {
		if game.p1 == "A" {
			return 3 + 0
		}
		if game.p1 == "B" {
			return 1 + 0
		}
		if game.p1 == "C" {
			return 2 + 0
		}
	}
	if game.p2 == "Y" {
		if game.p1 == "A" {
			return 1 + 3
		}
		if game.p1 == "B" {
			return 2 + 3
		}
		if game.p1 == "C" {
			return 3 + 3
		}
	}
	if game.p2 == "Z" {
		if game.p1 == "A" {
			return 2 + 6
		}
		if game.p1 == "B" {
			return 3 + 6
		}
		if game.p1 == "C" {
			return 1 + 6
		}
	}
	return 0

}
