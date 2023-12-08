
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// Parse the instructions and the map
	scanner := bufio.NewScanner(file)
	instructions := ""
	mapping := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		if len(instructions) == 0 {
			instructions = line
			continue
		}
		parts := strings.Split(line, " = ")
		nodes := strings.Trim(parts[1], "()")
		mapping[parts[0]] = strings.Split(nodes, ", ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Starting point is always AAA according to the task
	currentNode := "AAA"
	instructionsIndex := 0
	steps := 0

	// Repeat the instruction sequence until ZZZ is reached
	for currentNode != "ZZZ" {
		instruction := string(instructions[instructionsIndex])
		// Use modulo to loop over the instructions
		instructionsIndex = (instructionsIndex + 1) % len(instructions)

		// Left (L) is the first element, Right (R) is the second element
		switch instruction {
		case "L":
			currentNode = mapping[currentNode][0]
		case "R":
			currentNode = mapping[currentNode][1]
		default:
			fmt.Println("Invalid instruction:", instruction)
			return
		}

		steps++
	}

	fmt.Println(steps)
}
