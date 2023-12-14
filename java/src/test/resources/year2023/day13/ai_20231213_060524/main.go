
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var patterns [][]string
	currentPattern := []string{}

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if len(currentPattern) > 0 {
				patterns = append(patterns, currentPattern)
				currentPattern = []string{}
			}
		} else {
			currentPattern = append(currentPattern, line)
		}
	}

	if len(currentPattern) > 0 {
		patterns = append(patterns, currentPattern)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, p := range patterns {
		if len(p) > 0 {
			vertical, horizontal := findReflectionPoints(p)
			sum += vertical + 100*horizontal
		}
	}

	fmt.Println(sum)
}

func findReflectionPoints(pattern []string) (vertical int, horizontal int) {
	// Check for vertical reflection
	for i := 0; i < len(pattern[0])/2; i++ {
		reflected := true
		for _, row := range pattern {
			if row[i] != row[len(row)-1-i] {
				reflected = false
				break
			}
		}
		if reflected {
			vertical = i
			break
		}
	}

	// Check for horizontal reflection
	for i := 0; i < len(pattern)/2; i++ {
		reflected := true
		for col := 0; col < len(pattern[0]); col++ {
			if pattern[i][col] != pattern[len(pattern)-1-i][col] {
				reflected = false
				break
			}
		}
		if reflected {
			horizontal = i
			break
		}
	}

	return
}
