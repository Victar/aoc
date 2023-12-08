
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read the input data from the file "input.txt"
	instructions, network, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Process the instructions and network to find the number of steps required
	steps := followInstructions(instructions, network)
	fmt.Println(steps)
}

func readInput(filename string) (instructions string, network map[string][2]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	network = make(map[string][2]string)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue // skip empty lines
		}

		parts := strings.Split(line, " ")
		node := parts[1]
		targets := strings.Split(parts[3], ",")

		// Ensure the format of the inputs matches expectations
		if len(parts) != 5 || len(targets) != 2 || len(node) != 3 {
			return "", nil, fmt.Errorf("invalid input format: %s", line)
		}

		// Remove parentheses
		targets[0] = strings.Trim(targets[0], "(")
		targets[1] = strings.Trim(targets[1], ")")

		network[node] = [2]string{targets[0], targets[1]}
	}

	if scanner.Err() != nil {
		return "", nil, scanner.Err()
	}

	// The first non-empty line contains the sequence of instructions
	instructions = scanner.Text()
	return instructions, network, nil
}

func followInstructions(instructions string, network map[string][2]string) int {
	current := "AAA"
	steps := 0

	instructions = strings.ReplaceAll(instructions, " ", "")

	for {
		for _, direction := range instructions {
			if current == "ZZZ" {
				return steps
			}

			nodes := network[current]
			if direction == 'R' {
				current = nodes[1]
			} else if direction == 'L' {
				current = nodes[0]
			} else {
				log.Fatalf("invalid direction: %c", direction)
			}

			steps++
		}
	}
}
