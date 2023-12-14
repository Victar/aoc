
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	pattern := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if len(pattern) > 0 {
				sum += analyzePattern(pattern)
				pattern = []string{}
			}
		} else {
			pattern = append(pattern, line)
		}
	}
	if len(pattern) > 0 {
		sum += analyzePattern(pattern)
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func analyzePattern(pattern []string) int {
	mirrorType, index := findReflection(pattern)
	if mirrorType == "vertical" {
		return index
	}
	return 100 * index
}

// findReflection finds a line of reflection; it returns the type of mirror ("vertical" or "horizontal") and the index.
func findReflection(pattern []string) (string, int) {
	for i := 0; i < len(pattern[0])/2; i++ {
		mirror := true
		for _, row := range pattern {
			if row[i] != row[len(row)-1-i] {
				mirror = false
				break
			}
		}
		if mirror {
			return "vertical", i
		}
	}

	for i := 0; i < len(pattern)/2; i++ {
		mirror := true
		for j, row := range pattern {
			if row[j] != pattern[len(pattern)-1-j][j] {
				mirror = false
				break
			}
		}
		if mirror {
			return "horizontal", i
		}
	}

	return "", 0
}
