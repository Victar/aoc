
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
	scanner.Scan()
	timeStr := strings.Replace(scanner.Text(), "Time: ", "", 1)
	for _, t := range strings.Fields(timeStr) {
		time, _ := strconv.Atoi(t)
		times = append(times, time)
	}

	scanner.Scan()
	distanceStr := strings.Replace(scanner.Text(), "Distance: ", "", 1)
	for _, d := range strings.Fields(distanceStr) {
		distance, _ := strconv.Atoi(d)
		distances = append(distances, distance)
	}

	return times, distances, scanner.Err()
}

func calculateNumberOfWays(time, distance int) int {
	ways := 0
	for chargeTime := 1; chargeTime <= time; chargeTime++ {
		totalDistance := chargeTime * (time - chargeTime)
		if totalDistance > distance {
			ways++
		}
	}
	return ways
}

func calculateProductOfWays(times, distances []int) int {
	product := 1
	for i := 0; i < len(times); i++ {
		ways := calculateNumberOfWays(times[i], distances[i])
		product *= ways
	}
	return product
}

func main() {
	times, distances, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part One
	product := calculateProductOfWays(times, distances)
	fmt.Println(product)

	// Part Two
	// For part two, we will now treat the data as having no spaces and create a single, long race
	timeStr := strings.Join(strings.Fields(strings.Join(strings.Split(fmt.Sprint(times), " "), "")), "")
	distanceStr := strings.Join(strings.Fields(strings.Join(strings.Split(fmt.Sprint(distances), " "), "")), "")

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	ways := calculateNumberOfWays(time, distance)
	fmt.Println(ways)
}
