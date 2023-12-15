
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type LensSlot struct {
	Label      string
	FocalLength int
}

func hashAlgorithm(input string) int {
	currentValue := 0
	for _, character := range input {
		currentValue += int(character)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func parseInput(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\n", ""), ",")
}

func partOne(operations []string) int {
	sum := 0
	for _, operation := range operations {
		parts := strings.Split(operation, "=")
		if len(parts) == 1 {
			parts = strings.Split(operation, "-")
		}
		sum += hashAlgorithm(parts[0])
	}
	return sum
}

func partTwo(operations []string) int {
	boxes := make([][]LensSlot, 256)
	focusingPower := 0

	for _, operation := range operations {
		parts := strings.Split(operation, "=")
		removeOp := false
		if len(parts) == 1 {
			parts = strings.Split(operation, "-")
			removeOp = true
		}
		boxIndex := hashAlgorithm(parts[0])
		label := parts[0]

		if removeOp {
			// Remove the lens
			newSlots := []LensSlot{}
			for _, slot := range boxes[boxIndex] {
				if slot.Label != label {
					newSlots = append(newSlots, slot)
				}
			}
			boxes[boxIndex] = newSlots
		} else {
			focalLength, _ := strconv.Atoi(parts[1])
			replaced := false
			for i, slot := range boxes[boxIndex] {
				if slot.Label == label {
					boxes[boxIndex][i].FocalLength = focalLength
					replaced = true
					break
				}
			}
			if !replaced {
				boxes[boxIndex] = append(boxes[boxIndex], LensSlot{Label: label, FocalLength: focalLength})
			}
		}
	}

	// Calculate focusing power
	for boxIndex, slots := range boxes {
		for slotIndex, slot := range slots {
			focusingPower += (boxIndex + 1) * (slotIndex + 1) * slot.FocalLength
		}
	}

	return focusingPower
}

func main() {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read input file: %v", err)
		return
	}

	operations := parseInput(string(fileContent))

	fmt.Println(partOne(operations))
	fmt.Println(partTwo(operations))
}
