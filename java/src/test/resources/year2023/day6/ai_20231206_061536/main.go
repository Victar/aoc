
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var waysToWin []int // Store the number of ways to win each race

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		// Assuming input format is "Time: [time] Distance: [distance]"
		raceTime, _ := strconv.Atoi(parts[1])
		recordDistance, _ := strconv.Atoi(parts[3])

		var winCount int
		for i := 0; i < raceTime; i++ {
			distanceCovered := i * (raceTime - i)
			if distanceCovered > recordDistance {
				winCount++
			}
		}
		waysToWin = append(waysToWin, winCount)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate the product of all the ways to win
	product := 1
	for _, count := range waysToWin {
		product *= count
	}

	fmt.Println(product)
}
