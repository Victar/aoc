
// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numberWords = map[string]string{
	"one": "1", "two": "2", "three": "3",
	"four": "4", "five": "5", "six": "6",
	"seven": "7", "eight": "8", "nine": "9",
}

func extractDigits(line string) (int, error) {
	// Convert text numbers to actual digits within the line.
	for text, digit := range numberWords {
		line = strings.Replace(line, text, digit, -1)
	}

	// Extract first and last digit from the line.
	var firstDigit, lastDigit rune
	foundFirst := false
	for _, r := range line {
		if r >= '0' && r <= '9' {
			if !foundFirst {
				firstDigit = r
				foundFirst = true
			}
			lastDigit = r
		}
	}

	// Combine the digits to form a two-digit number.
	twoDigitValue := string(firstDigit) + string(lastDigit)
	return strconv.Atoi(twoDigitValue)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		value, err := extractDigits(line)
		if err != nil {
			fmt.Println("Error extracting digits:", err)
			return
		}
		sum += value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(sum)
}
