
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const numBoxes = 256

type Lens struct {
	Label string
	Focal int
}

func hash(s string) int {
	val := 0
	for i := 0; i < len(s); i++ {
		val = (val + int(s[i])) * 17 % 256
	}
	return val
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	steps := strings.Split(input, ",")
	totalHashSum := 0

	// Part 1 - Calculating the sum of HASH algorithm results
	for _, step := range steps {
		result := hash(step)
		totalHashSum += result
	}

	// Part 2 - Calculating the focusing power of the lens configuration
	lenses := make([][]Lens, numBoxes)
	for _, step := range steps {
		label := strings.TrimRight(step, "-=0123456789")
		index := hash(label)
		switch {
		case strings.Contains(step, "="):
			focalLength := int(step[len(step)-1] - '0')
			found := false
			for i, lens := range lenses[index] {
				if lens.Label == label {
					lenses[index][i].Focal = focalLength // Replace the lens
					found = true
					break
				}
			}
			if !found {
				lenses[index] = append(lenses[index], Lens{label, focalLength}) // Add new lens
			}
		case strings.Contains(step, "-"):
			for i, lens := range lenses[index] {
				if lens.Label == label {
					copy(lenses[index][i:], lenses[index][i+1:])
					lenses[index] = lenses[index][:len(lenses[index])-1] // Remove the lens
					break
				}
			}
		}
	}

	var totalFocusingPower int
	for boxIndex, box := range lenses {
		for slotIndex, lens := range box {
			totalFocusingPower += (boxIndex + 1) * (slotIndex + 1) * lens.Focal
		}
	}
	
	fmt.Println(totalHashSum)
	fmt.Println(totalFocusingPower)
}
