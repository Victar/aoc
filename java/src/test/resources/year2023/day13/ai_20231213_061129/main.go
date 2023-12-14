package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day13/sample.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstHalf, secondHalf := splitPattern(line)
		sum += reflect(firstHalf, secondHalf)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input:", err)
		return
	}

	fmt.Println(sum)
}

// Split the pattern at the point of symmetry.
func splitPattern(line string) (firstHalf, secondHalf []byte) {
	length := len(line)

	// Check for vertical symmetry
	for mid := 1; mid < length; mid++ {
		if isReflectionVertical(line, mid) {
			return []byte(line[:mid]), []byte(line[length-mid:])
		}
	}

	// Check for horizontal symmetry
	for top := 0; top < length/2; top++ {
		if line[top] != line[length-1-top] {
			return []byte(line[:length/2]), []byte(line[length/2:])
		}
	}

	return []byte(line[:length/2]), []byte(line[length/2:])
}

// Determine if there's a vertical reflection.
func isReflectionVertical(line string, mid int) bool {
	for i, j := mid-1, mid; i >= 0 && j < len(line); i, j = i-1, j+1 {
		if line[i] != line[j] {
			return false
		}
	}
	return true
}

// Reflect the pattern and calculate the summary.
func reflect(firstHalf, secondHalf []byte) int {
	vReflectCount := 0
	hReflectCount := 0

	if len(firstHalf) != len(secondHalf) {
		vReflectCount = len(firstHalf)
	} else {
		hReflectCount = len(firstHalf)
	}

	return vReflectCount + 100*hReflectCount
}
