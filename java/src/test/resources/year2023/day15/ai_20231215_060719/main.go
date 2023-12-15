
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label      string
	FocalLength int
}

func main() {
	// Read the input from "input.txt".
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	// Split the steps by commas.
	steps := strings.Split(input, ",")

	// Implement the described HASHMAP process.
	var boxes [256][]Lens
	totalFocusingPower := 0
	for _, step := range steps {
		label := strings.TrimRight(step, "-0123456789=")
		boxIndex := hash(label)

		switch {
		case strings.Contains(step, "-"):
			// Remove the lens with the current label from the box, if present.
			lensRemoved := false
			for i, lens := range boxes[boxIndex] {
				if lens.Label == label {
					// Remove the lens by slicing it out of the array.
					boxes[boxIndex] = append(boxes[boxIndex][:i], boxes[boxIndex][i+1:]...)
					lensRemoved = true
					break
				}
			}

			// Recalculate focusing power if lens was removed.
			if lensRemoved {
				totalFocusingPower = calculateFocusingPower(boxes)
			}

		case strings.Contains(step, "="):
			// Parse the focal length from step.
			focalLengthStr := strings.Split(step, "=")[1]
			focalLength, _ := strconv.Atoi(focalLengthStr)

			// Find if a lens with the same label exists and replace it.
			lensReplaced := false
			for i, lens := range boxes[boxIndex] {
				if lens.Label == label {
					boxes[boxIndex][i].FocalLength = focalLength
					lensReplaced = true
					break
				}
			}

			// If the lens was not replaced, add new lens to the box.
			if !lensReplaced {
				boxes[boxIndex] = append(boxes[boxIndex], Lens{Label: label, FocalLength: focalLength})
			}

			totalFocusingPower = calculateFocusingPower(boxes)
		}
	}

	fmt.Println(totalFocusingPower)
}

// HASH algorithm implementation.
func hash(s string) int {
	currentValue := 0
	for _, ch := range s {
		currentValue = ((currentValue + int(ch)) * 17) % 256
	}
	return currentValue
}

// Calculate the total focusing power for all boxes with lenses.
func calculateFocusingPower(boxes [256][]Lens) int {
	totalPower := 0
	for boxNum, lenses := range boxes {
		for slotNum, lens := range lenses {
			totalPower += (boxNum + 1) * (slotNum + 1) * lens.FocalLength
		}
	}
	return totalPower
}
