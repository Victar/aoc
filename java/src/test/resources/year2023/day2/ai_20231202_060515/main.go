
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var gamesPossibleSum, powerSum int
	for scanner.Scan() {
		line := scanner.Text()
		gameId, r, g, b := parseGame(line)
		if isPossible(12, 13, 14, r, g, b) {
			gamesPossibleSum += gameId
		}
		minRed, minGreen, minBlue := findMinimumCubes(r, g, b)
		powerSum += minRed * minGreen * minBlue
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("The sum of the IDs of games possible with 12 red, 13 green, 14 blue cubes: %d\n", gamesPossibleSum)
	fmt.Printf("The sum of the power of the minimum sets of cubes: %d\n", powerSum)
}

func parseGame(line string) (gameId int, red, green, blue []int) {
	parts := strings.Split(line, ":")
	id := strings.TrimSpace(strings.Split(parts[0], " ")[1])
	gameId, _ = strconv.Atoi(id)

	revealedSets := strings.Split(parts[1], ";")
	for _, set := range revealedSets {
		set = strings.TrimSpace(set)
		cubes := strings.Split(set, ",")
		var r, g, b int
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			colorInfo := strings.Split(cube, " ")
			count, _ := strconv.Atoi(colorInfo[0])
			switch colorInfo[1] {
			case "red":
				r += count
			case "green":
				g += count
			case "blue":
				b += count
			}
		}
		red = append(red, r)
		green = append(green, g)
		blue = append(blue, b)
	}
	return
}

func isPossible(red, green, blue int, r, g, b []int) bool {
	for i := range r {
		if r[i] > red || g[i] > green || b[i] > blue {
			return false
		}
	}
	return true
}

func findMinimumCubes(r, g, b []int) (minRed, minGreen, minBlue int) {
	for _, val := range r {
		if val > minRed {
			minRed = val
		}
	}
	for _, val := range g {
		if val > minGreen {
			minGreen = val
		}
	}
	for _, val := range b {
		if val > minBlue {
			minBlue = val
		}
	}
	return
}
