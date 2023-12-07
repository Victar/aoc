
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseLine reads a line of text and returns the time and record distances.
func parseLine(line string) (int, int) {
	parts := strings.Split(line, " ")
	time, _ := strconv.Atoi(parts[1])
	distance, _ := strconv.Atoi(parts[3])
	return time, distance
}

// countWaysToWin computes the number of ways you can hold the button to win the race.
func countWaysToWin(time, record int) int {
	waysToWin := 0
	for i := 0; i < time; i++ {
		speed := i
		travelTime := time - i
		distance := speed * travelTime
		if distance > record {
			waysToWin++
		}
	}
	return waysToWin
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	product := 1
	for scanner.Scan() {
		line := scanner.Text()
		time, record := parseLine(line)
		waysToWin := countWaysToWin(time, record)
		product *= waysToWin
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println(product)
}
