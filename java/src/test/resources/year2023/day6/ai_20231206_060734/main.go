
// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	result := 1
	for _, race := range input {
		ways := calculateWaysToWin(race.time, race.distance)
		result *= ways
	}

	fmt.Println(result)
}

// Race represents a single race's time limit and the record distance to beat.
type Race struct {
	time     int
	distance int
}

// readInput reads the input file and returns a slice of Race structs.
func readInput(filename string) ([]Race, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var races []Race
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) != 2 {
			continue
		}

		time, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}

		distance, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		races = append(races, Race{time: time, distance: distance})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return races, nil
}

// calculateWaysToWin calculates how many ways there are to win the race by beating the record.
func calculateWaysToWin(time, distance int) int {
	var waysToWin int
	for holdTime := 1; holdTime < time; holdTime++ {
		moveTime := time - holdTime
		totalDistance := holdTime * moveTime
		if totalDistance > distance {
			waysToWin++
		}
	}
	return waysToWin
}
