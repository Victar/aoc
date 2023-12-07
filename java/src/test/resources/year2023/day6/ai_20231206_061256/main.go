
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// maxDistance calculates the maximum distance a boat can travel given the time and initial speed.
func maxDistance(time, speed int) int {
	return (time-speed)*speed
}

// waysToWin calculates the number of ways the boat can win the race, given the race time and the record distance.
func waysToWin(time, record int) int {
	ways := 0
	for speed := 1; speed < time; speed++ {
		if maxDistance(time, speed) > record {
			ways++
		}
	}
	return ways
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	multiplyWays := 1 // Initialize with 1 so it can multiply correctly for the first value.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			time, err1 := strconv.Atoi(parts[1])
			record, err2 := strconv.Atoi(parts[3])
			if err1 != nil || err2 != nil {
				fmt.Println("Error converting string to int:", err1, err2)
				return
			}

			multiplyWays *= waysToWin(time, record)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println(multiplyWays)
}
