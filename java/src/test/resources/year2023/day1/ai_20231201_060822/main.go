
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// mapping of spelled digits to their numeric representations
var digitWords = map[string]string{
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

// findDigits extracts the first and last digit from a line
func findDigits(line string) (int, int) {
	// Replace spelled out numbers with digits
	for word, digit := range digitWords {
		line = strings.Replace(line, word, digit, -1)
	}

	first, last := -1, -1
	for _, char := range line {
		if char >= '0' && char <= '9' {
			digit := int(char - '0')
			if first == -1 {
				first = digit
			}
			last = digit
		}
	}

	return first, last
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	sum := 0 // Initial sum is 0

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := findDigits(line)
		if firstDigit != -1 && lastDigit != -1 {
			// Convert the first and last digits to a number and add to sum
			sum += firstDigit*10 + lastDigit
		}
	}

	// Handle any errors that occurred while reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Print the final sum
	fmt.Println(sum)
}
