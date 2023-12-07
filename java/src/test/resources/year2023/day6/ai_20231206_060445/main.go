
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
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var races [][]int
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 4 {
			continue
		}
		time, _ := strconv.Atoi(parts[1])
		distance, _ := strconv.Atoi(parts[3])
		races = append(races, []int{time, distance})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalWays := 1
	for _, race := range races {
		totalWays *= calculateWaysToWin(race[0], race[1])
	}

	fmt.Println(totalWays)
}

func calculateWaysToWin(time, record int) int {
	waysToWin := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		distance := holdTime * (time - holdTime)
		if distance > record {
			waysToWin++
		}
	}
	return waysToWin
}
