
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (times []int, distances []int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip 'Time:' label
	scanner.Scan() // read times
	timesStr := strings.Fields(scanner.Text())
	scanner.Scan() // skip 'Distance:' label
	scanner.Scan() // read distances
	distancesStr := strings.Fields(scanner.Text())

	for _, t := range timesStr {
		time, err := strconv.Atoi(t)
		if err != nil {
			return nil, nil, err
		}
		times = append(times, time)
	}

	for _, d := range distancesStr {
		distance, err := strconv.Atoi(d)
		if err != nil {
			return nil, nil, err
		}
		distances = append(distances, distance)
	}

	return times, distances, nil
}

func calculateWays(times []int, distances []int) int {
	totalWays := 1
	for i := range times {
		ways := 0
		for buttonTime := 1; buttonTime < times[i]; buttonTime++ {
			distance := buttonTime * (times[i] - buttonTime)
			if distance > distances[i] {
				ways++
			}
		}
		totalWays *= ways
	}
	return totalWays
}

func main() {
	times, distances, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Part One
	totalWays := calculateWays(times, distances)
	fmt.Println(totalWays)

	// Part Two
	concatenatedTimes, _ := strconv.Atoi(strings.Join(strings.Fields(strings.TrimSpace(string(times))), ""))
	concatenatedDistances, _ := strconv.Atoi(strings.Join(strings.Fields(strings.TrimSpace(string(distances))), ""))
	ways := 0
	for buttonTime := 1; buttonTime < concatenatedTimes; buttonTime++ {
		if concatenatedDistances < buttonTime*(concatenatedTimes-buttonTime) {
			ways++
		}
	}
	fmt.Println(ways)
}
