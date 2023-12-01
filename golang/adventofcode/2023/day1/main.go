package main

import (
	"adventofcode/util"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	runSilver()
	runGold()
}

var numberWords = map[string]string{
	"one": "1", "two": "2", "three": "3",
	"four": "4", "five": "5", "six": "6",
	"seven": "7", "eight": "8", "nine": "9",
}

func runSilver() {
	ar, err := util.ReadFile("year2023/day1/input.txt")
	if err != nil {
		panic(err)
	}
	var sum = 0
	for _, line := range ar {
		firstDigit, lastDigit := findDigits(line)
		if firstDigit == -1 || lastDigit == -1 {
			continue // Skip lines with no digits.
		}
		calibrationValue, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		if err != nil {
			log.Fatal(err)
		}
		sum += calibrationValue
	}
	println(sum)
}

func runGold() {
	ar, err := util.ReadFile("year2023/day1/input.txt")
	if err != nil {
		panic(err)
	}
	var sum = 0
	for _, line := range ar {
		line = numberToDigit(line)
		firstDigit, lastDigit := findDigits(line)
		if firstDigit == -1 || lastDigit == -1 {
			continue // Skip lines with no digits.
		}
		calibrationValue, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		if err != nil {
			log.Fatal(err)
		}
		sum += calibrationValue
	}
	println(sum)
}

func numberToDigit(line string) string {
	repIndex := 100000
	repDigit := ""

	for text, digit := range numberWords {
		curIndex := strings.Index(line, text)
		if curIndex > -1 && repIndex > curIndex {
			repIndex = curIndex
			repDigit = digit
		}
	}
	if repDigit != "" {
		line = line[0:repIndex] + repDigit + line[repIndex+1:] //strings.Replace(line, repText, repDigit, 1)
		return numberToDigit(line)
	}
	return line
}

// findDigits returns the first and last digits in a string.
func findDigits(s string) (firstDigit int, lastDigit int) {
	firstDigit = -1
	lastDigit = -1

	for _, r := range s {
		if unicode.IsDigit(r) {
			if firstDigit == -1 {
				firstDigit = int(r - '0')
			}
			lastDigit = int(r - '0')
		}
	}

	return firstDigit, lastDigit
}
