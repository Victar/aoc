
// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0 // Part 1 sum
	sumPartTwo := 0 // Part 2 sum
	re := regexp.MustCompile(`\D+`) // Non-digit characters
	wordToDigit := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			digits := re.ReplaceAllString(word, "")
			if len(digits) >= 2 {
				firstDigit := string(digits[0])
				lastDigit := string(digits[len(digits)-1])
				value, _ := strconv.Atoi(firstDigit + lastDigit)
				sum += value
			}
		}

		// Part Two: replacing written numbers with digits
		lineModified := line
		for word, digit := range wordToDigit {
			lineModified = strings.ReplaceAll(lineModified, word, strconv.Itoa(digit))
		}

		digits := re.ReplaceAllString(lineModified, "")
		if len(digits) >= 2 {
			firstDigit := string(digits[0])
			lastDigit := string(digits[len(digits)-1])
			value, _ := strconv.Atoi(firstDigit + lastDigit)
			sumPartTwo += value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the input file:", err)
		os.Exit(1)
	}

	// Print the answer
	fmt.Println(sum) // Print sum for Part 1
	fmt.Println(sumPartTwo) // Print sum for Part 2
}
