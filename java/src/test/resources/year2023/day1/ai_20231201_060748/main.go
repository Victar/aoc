
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
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
