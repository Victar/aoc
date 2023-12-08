
// main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	instructions := ""
	nodes := make(map[string][2]string)

	for scanner.Scan() {
		line := scanner.Text()
		if instructions == "" {
			instructions = line
		} else {
			parts := strings.Fields(line)
			nodes[parts[0]] = [2]string{parts[2], parts[3]}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	currentNode := "AAA"
	steps, instructionLength := 0, len(instructions)

	// Navigate the network following the instructions
	for currentNode != "ZZZ" {
		direction := instructions[steps%instructionLength]
		if direction == 'R' {
			currentNode = nodes[currentNode][1]
		} else if direction == 'L' {
			currentNode = nodes[currentNode][0]
		}
		steps++
	}

	// Print the answer (number of steps)
	fmt.Print(steps)
}
