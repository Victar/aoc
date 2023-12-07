
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

	scanner := bufio.NewScanner(file)

	var raceCount int
	var result int64 = 1 // initialize product as 1; we'll be multiplying against it later

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 2 {
			continue // skip any malformed line
		}

		time, errTime := strconv.Atoi(parts[1])
		distance, errDistance := strconv.Atoi(parts[3])
		if errTime != nil || errDistance != nil {
			log.Fatal("Error converting input to integers", errTime, errDistance)
		}

		waysToWin := calculateWaysToWin(time, distance)
		result *= int64(waysToWin) // accumulate the result through multiplication
		raceCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func calculateWaysToWin(totalTime int, recordDistance int) int {
	possibleWays := 0
	for holdTime := 1; holdTime < totalTime; holdTime++ {
		remainingTime := totalTime - holdTime
		distance := holdTime * remainingTime
		if distance > recordDistance {
			possibleWays++
		}
	}
	return possibleWays
}
