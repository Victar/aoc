
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var waysToWin []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 2 {
			continue
		}
		timeAllowed, err1 := strconv.Atoi(parts[0])
		bestDistance, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		// Calculate ways to win for this race
		ways := calculateWaysToWin(timeAllowed, bestDistance)
		waysToWin = append(waysToWin, ways)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Multiply the number of ways to win for each race
	result := int64(1)
	for _, ways := range waysToWin {
		result *= ways
	}

	fmt.Println(result)
}

func calculateWaysToWin(timeAllowed int, bestDistance int) int64 {
	var waysToWin int64 = 0
	for holdTime := 0; holdTime <= timeAllowed; holdTime++ {
		distanceTraveled := (timeAllowed - holdTime) * holdTime
		if distanceTraveled > bestDistance {
			waysToWin++
		}
	}
	return waysToWin
}
