package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
)

var DAY = "22"

const (
	iterations = 2000
	modulo     = 16777216
)

var scores = map[[4]int]int{}

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	result := 0
	for _, line := range lines {
		curInt, _ := strconv.Atoi(line)
		for i := 0; i < iterations; i++ {
			curInt = nextSecret(curInt)
		}
		result += curInt
	}
	fmt.Println(result)
}

func runGold() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		curInt, _ := strconv.Atoi(line)
		prices := generatePrices(curInt)
		curScores := getScores(prices)
		for curKey, curScore := range curScores {
			scores[curKey] += curScore
		}
	}
	maxScore := 0
	for _, v := range scores {
		if v > maxScore {
			maxScore = v
		}
	}
	fmt.Println(maxScore)
}

func getScores(prices []int) map[[4]int]int {
	curScores := make(map[[4]int]int)
	for i := 4; i < len(prices); i++ {
		keyCur := [4]int{prices[i-3] - prices[i-4], prices[i-2] - prices[i-3], prices[i-1] - prices[i-2], prices[i] - prices[i-1]}
		if _, ok := curScores[keyCur]; !ok {
			curScores[keyCur] = prices[i]
		}
	}
	return curScores
}

func generatePrices(secret int) []int {
	prices := make([]int, iterations)
	for i := 0; i < iterations; i++ {
		prices[i] = secret % 10
		secret = nextSecret(secret)
	}
	return prices
}

func nextSecret(secret int) int {
	secret = (secret ^ (secret * 64)) % modulo
	secret = (secret ^ (secret / 32)) % modulo
	secret = (secret ^ (secret * 2048)) % modulo
	return secret
}
