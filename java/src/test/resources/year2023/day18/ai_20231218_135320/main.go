
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	var x, y, minX, minY, maxX, maxY int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var direction rune
		var command string
		fmt.Sscanf(line, "%s", &command)
		colorCode := command[4 : len(command)-1]                  // Extract the color code
		distance, _ := strconv.ParseInt(colorCode[:5], 16, 64)     // Parse the distance from hex
		directionCode, _ := strconv.ParseInt(colorCode[5:6], 16, 64) // Parse the direction code

		// Map the last digit of color code to a direction
		switch directionCode {
		case 0:
			direction = 'R'
		case 1:
			direction = 'D'
		case 2:
			direction = 'L'
		case 3:
			direction = 'U'
		default:
			log.Fatalf("invalid direction code: %d", directionCode)
		}

		// Move in the specified direction
		for i := int64(0); i < distance; i++ {
			switch direction {
			case 'R':
				x++
			case 'D':
				y--
			case 'L':
				x--
			case 'U':
				y++
			}

			// Update min and max coordinates
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate the area covered by the trench using the largest rectangle that contains the trench
	// Since area of lava equals the area of the trench itself plus 1 for the starting cube, here we add +1 to both width and height
	area := (maxX - minX + 1) * (maxY - minY + 1)
	fmt.Println(area)
}
