
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
		log.Fatalf("Error opening input file: %s", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var raceTimes []int
	var distances []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) >= 2 {
			time, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("Error parsing race time '%s': %s", parts[1], err.Error())
			}
			raceTimes = append(raceTimes, time)

			distance, err := strconv.Atoi(parts[3])
			if err != nil {
				log.Fatalf("Error parsing record distance '%s': %s", parts[3], err.Error())
			}
			distances = append(distances, distance)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file: %s", err.Error())
	}

	totalWays := 1
	for idx, raceTime := range raceTimes {
		recordDistance := distances[idx]
		waysToWin := 0
		for holdTime := 1; (raceTime - holdTime) * holdTime > recordDistance; holdTime++ {
			waysToWin++
		}
		totalWays *= waysToWin
	}

	fmt.Println(totalWays)
}
