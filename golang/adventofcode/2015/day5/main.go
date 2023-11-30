package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func main() {
	runAny()
}

func runAny() {
	ar, err := util.ReadFile("year2015/day5/input.txt")
	if err != nil {
		panic(err)
	}

	niceCount := 0
	for _, str := range ar {
		if isNiceString(str) {
			niceCount++
		}
	}

	fmt.Printf("Number of nice strings: %d\n", niceCount)
}

func isNiceString1(s string) bool {
	// Rule 1: Check for at least three vowels
	vowelCount := 0
	vowels := "aeiou"
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		}
	}
	if vowelCount < 3 {
		return false
	}

	// Rule 2: Check for a double letter
	hasDoubleLetter := false
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			hasDoubleLetter = true
			break
		}
	}
	if !hasDoubleLetter {
		return false
	}

	// Rule 3: Check for disallowed substrings
	disallowedSubstrings := []string{"ab", "cd", "pq", "xy"}
	for _, substring := range disallowedSubstrings {
		if strings.Contains(s, substring) {
			return false
		}
	}

	// All rules passed, it's a nice string
	return true
}

func isNiceString(s string) bool {
	// Rule 1: Check for a pair of two letters that appears at least twice without overlapping
	hasNonOverlappingPair := false
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Count(s, pair) >= 2 && strings.Index(s, pair) != strings.LastIndex(s, pair) {
			hasNonOverlappingPair = true
			break
		}
	}

	// Rule 2: Check for a letter that repeats with exactly one letter between them
	hasRepeatingLetter := false
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			hasRepeatingLetter = true
			break
		}
	}

	// A nice string meets both rules
	return hasNonOverlappingPair && hasRepeatingLetter
}
