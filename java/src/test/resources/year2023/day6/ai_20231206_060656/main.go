
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxDistance(time, chargeTime int) int {
	return (time-chargeTime)*(time-chargeTime+1)/2
}

func waysToWin(time, record int) int {
	ways := 0
	for chargeTime := 0; chargeTime < time; chargeTime++ {
		if maxDistance(time, chargeTime) > record {
			ways++
		}
	}
	return ways
}

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var times, records []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			line = strings.TrimPrefix(line, "Time:")
			parts := strings.Fields(line)
			for _, part := range parts {
				time, _ := strconv.Atoi(part)
				times = append(times, time)
			}
		} else if strings.HasPrefix(line, "Distance:") {
			line = strings.TrimPrefix(line, "Distance:")
			parts := strings.Fields(line)
			for _, part := range parts {
				record, _ := strconv.Atoi(part)
				records = append(records, record)
			}
		}
	}
	return times, records, scanner.Err()
}

func main() {
	times, records, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	product := 1
	for i := 0; i < len(times); i++ {
		wins := waysToWin(times[i], records[i])
		product *= wins
	}

	fmt.Println(product)
}
