
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func winningWaysSingleRace(time, distance int) int {
	// Determine the minimum and maximum hold time to beat the distance record
	minHoldTime := ((8*distance + 1) - 1)/2*time;
	maxHoldTime := time - 1;

	// Calculate the number of ways to win
	waysToWin := maxHoldTime - minHoldTime + 1
	if waysToWin < 0 {
		return 0
	}
	return waysToWin
}

func main() {
	// Open input data file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeLine := scanner.Text()
	scanner.Scan()
	distanceLine := scanner.Text()

	// Remove "Time: " and "Distance: " prefixes, split each line into times and distances
	timeParts := strings.Fields(strings.TrimPrefix(timeLine, "Time: "))
	distanceParts := strings.Fields(strings.TrimPrefix(distanceLine, "Distance: "))

	// Parse the combined race data without spaces
	combinedTime, _ := strconv.Atoi(strings.Join(timeParts, ""))
	combinedDistance, _ := strconv.Atoi(strings.Join(distanceParts, ""))

	// Calculate and print the number of ways to win the combined race
	combinedWaysToWin := winningWaysSingleRace(combinedTime, combinedDistance)
	fmt.Println(combinedWaysToWin)
}
