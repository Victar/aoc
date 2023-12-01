
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// map for converting spelled-out numbers to digits
var numberMap = map[string]string{
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

func parseAndSumLines(scanner *bufio.Scanner) int {
	total := 0

	// regex to find sequences that represent numbers (both text and digits)
	numberRegex := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)

	for scanner.Scan() {
		line := scanner.Text()

		// find all sequences representing numbers from the line
		matches := numberRegex.FindAllString(line, -1)

		if len(matches) > 1 {
			// get the first and last "digit" in the line
			first := convertToDigit(matches[0])
			last := convertToDigit(matches[len(matches)-1])

			// combine them into a two-digit number
			calibrationValue, _ := strconv.Atoi(first + last)
			total += calibrationValue
		}
	}

	return total
}

func convertToDigit(s string) string {
	if val, ok := numberMap[s]; ok {
		return val
	}
	return s
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := parseAndSumLines(scanner)
	fmt.Println(sum)
}
