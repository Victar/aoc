
// Save this code in a file named "main.go".

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var wordToDigit = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// findDigits finds and returns the first and last "digit" in the line.
func findDigits(line string) (string, string, error) {
	words := strings.Fields(line)
	first := words[0]
	last := words[len(words)-1]

	// Use regex to extract the numeric part of the string
	re := regexp.MustCompile(`\d`)
	firstNum := re.FindString(first)
	lastNum := re.FindString(last)

	// If digits are not spelled out, use them as is
	if firstNum != "" && lastNum != "" {
		return firstNum, lastNum, nil
	}

	// If digits are spelled out, translate words to digits
	for key, value := range wordToDigit {
		if strings.Contains(first, key) {
			firstNum = value
		}
		if strings.Contains(last, key) {
			lastNum = value
		}
	}

	if firstNum == "" || lastNum == "" {
		return "", "", fmt.Errorf("couldn't find digits in line: %s", line)
	}

	return firstNum, lastNum, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	totalSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit, err := findDigits(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		value, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			continue
		}
		totalSum += value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	fmt.Println(totalSum)
}
