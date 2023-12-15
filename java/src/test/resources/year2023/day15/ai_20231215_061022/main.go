
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxBoxes = 256

type Lens struct {
	Label    string
	FocalLen int
}

func hashAlgorithm(label string) int {
	value := 0
	for _, c := range label {
		value += int(c)
		value *= 17
		value %= 256
	}
	return value
}

func readInput(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	return strings.ReplaceAll(input, "\n", ""), nil
}

func main() {
	// Read initialization sequence from input.txt
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	boxes := make([][]Lens, maxBoxes)
	steps := strings.Split(input, ",")

	sumHash := 0
	totalFocusingPower := 0
	for _, step := range steps {
		label, op, focalLenStr := parseStep(step)
		hash := hashAlgorithm(label)
		sumHash += hash

		switch op {
		case "-":
			removeLens(boxes, label, hash)
		case "=":
			focalLen, _ := strconv.Atoi(focalLenStr)
			insertLens(boxes, Lens{label, focalLen}, hash)
		}
	}

	for boxIdx, lenses := range boxes {
		for slotIdx, lens := range lenses {
			focusingPower := (boxIdx + 1) * (slotIdx + 1) * lens.FocalLen
			totalFocusingPower += focusingPower
		}
	}

	fmt.Println(totalFocusingPower)
}

func parseStep(step string) (label, op, focalLenStr string) {
	if strings.Contains(step, "=") {
		parts := strings.Split(step, "=")
		return parts[0], "=", parts[1]
	}
	parts := strings.Split(step, "-")
	return parts[0], "-", ""
}

func removeLens(boxes [][]Lens, label string, boxIdx int) {
	lensIdx := -1
	for i, lens := range boxes[boxIdx] {
		if lens.Label == label {
			lensIdx = i
			break
		}
	}
	if lensIdx != -1 {
		boxes[boxIdx] = append(boxes[boxIdx][:lensIdx], boxes[boxIdx][lensIdx+1:]...)
	}
}

func insertLens(boxes [][]Lens, lens Lens, boxIdx int) {
	for i, existingLens := range boxes[boxIdx] {
		if existingLens.Label == lens.Label {
			boxes[boxIdx][i] = lens
			return
		}
	}
	boxes[boxIdx] = append(boxes[boxIdx], lens)
}
