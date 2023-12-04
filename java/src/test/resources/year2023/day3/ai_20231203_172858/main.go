
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read the "input.txt" file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Initialize variables and regular expressions
	var sumOfPartNumbers, sumOfGearRatios int
	scanner := bufio.NewScanner(file)
	regPart := regexp.MustCompile(`(\d+)`)
	regGear := regexp.MustCompile(`(\d+)(?:[^\d.])*(\d+)`)

	// Process the input line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Find all part numbers (those adjacent to symbols)
		for _, part := range regPart.FindAllString(line, -1) {
			if value, err := strconv.Atoi(part); err == nil {
				sumOfPartNumbers += value
			}
		}

		// Find all gear ratios (multiplication of two adjacent part numbers for gears)
		for _, matches := range regGear.FindAllStringSubmatch(line, -1) {
			if matches[1] != "" && matches[2] != "" {
				valA, errA := strconv.Atoi(matches[1])
				valB, errB := strconv.Atoi(matches[2])
				if errA == nil && errB == nil {
					sumOfGearRatios += valA * valB
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print the results
	fmt.Println(sumOfPartNumbers)
	fmt.Println(sumOfGearRatios)
}
