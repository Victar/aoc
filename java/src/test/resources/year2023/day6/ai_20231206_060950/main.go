
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
	var totalWays int64 = 1
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) < 2 {
			continue
		}
		time, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("failed to parse time: %v", err)
		}
		distance, err := strconv.Atoi(line[3])
		if err != nil {
			log.Fatalf("failed to parse distance: %v", err)
		}
		
		ways := calculateWaysToWin(time, distance)
		totalWays *= int64(ways)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalWays)
}

func calculateWaysToWin(time, record int) int {
	var ways int
	for holdTime := 0; holdTime < time; holdTime++ {
		travelTime := time - holdTime
		distance := holdTime * travelTime
		if distance > record {
			ways++
		}
	}
	return ways
}
