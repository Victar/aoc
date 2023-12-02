
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct to hold color counts
type ColorCount struct {
	red, green, blue int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPossibleGamesSum := 0

	// Create the reference bag with its count of cubes
	referenceBag := ColorCount{red: 12, green: 13, blue: 14}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Extract the game ID
		parts := strings.SplitN(line, ":", 2)
		gameID, _ := strconv.Atoi(strings.Fields(parts[0])[1])

		// Process the sets of cubes for the current game
		cubeSets := strings.Split(parts[1], ";")
		gamePossible := true

		for _, setCubes := range cubeSets {
			cubes := strings.Split(setCubes, ",")
			cubeCounts := ColorCount{red: 0, green: 0, blue: 0}

			for _, cube := range cubes {
				cubeData := strings.Fields(strings.TrimSpace(cube))

				if len(cubeData) < 2 { // Skip if data is not well formed
					continue
				}

				count, _ := strconv.Atoi(cubeData[0])
				color := cubeData[1]

				switch color {
				case "red":
					cubeCounts.red += count
					if cubeCounts.red > referenceBag.red {
						gamePossible = false
					}
				case "green":
					cubeCounts.green += count
					if cubeCounts.green > referenceBag.green {
						gamePossible = false
					}
				case "blue":
					cubeCounts.blue += count
					if cubeCounts.blue > referenceBag.blue {
						gamePossible = false
					}
				}
				
				if !gamePossible {
					break
				}
			}

			if !gamePossible {
				break
			}
		}

		if gamePossible {
			totalPossibleGamesSum += gameID
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input.txt:", err)
	} else {
		fmt.Println(totalPossibleGamesSum)
	}
}
