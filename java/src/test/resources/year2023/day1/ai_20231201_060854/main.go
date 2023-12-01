
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberWordsMap = map[string]string{
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

// SpellOutDigits replaces spelled out numbers with their digit representation.
func SpellOutDigits(line string) string {
	for word, digit := range numberWordsMap {
		line = strings.ReplaceAll(line, word, digit)
	}
	return line
}

// ExtractFirstAndLastDigit extracts the first and last digit from a string.
func ExtractFirstAndLastDigit(s string) int {
	// Find all digits in the string using regex.
	re := regexp.MustCompile(`\d`)
	digits := re.FindAllString(s, -1)

	// Convert the first and last digits to a two-digit number.
	firstDigit, _ := strconv.Atoi(digits[0])
	lastDigit, _ := strconv.Atoi(digits[len(digits)-1])
	number := firstDigit*10 + lastDigit

	return number
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = SpellOutDigits(line) // Takes care of Part Two
		sum += ExtractFirstAndLastDigit(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}

	fmt.Println(sum)
}
