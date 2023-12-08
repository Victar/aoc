
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Read the input file
    inputData, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    // Process the input data
    lines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
    instructions := lines[0] // Assumes the first line is the instructions
    nodeMappings := make(map[string][]string)
    for _, line := range lines[1:] {
        parts := strings.Split(line, " ")
        nodeMappings[parts[1]] = []string{parts[3], parts[5]}
    }

    // Start navigation
    currentNode := "AAA"
    stepCount := 0
    instructionIndex := 0
    for currentNode != "ZZZ" {
        // Find the next node based on direction
        direction := instructions[instructionIndex]
        if direction == 'R' {
            currentNode = nodeMappings[currentNode][1]
        } else { // Assumes the direction can only be 'R' or 'L'
            currentNode = nodeMappings[currentNode][0]
        }

        // Increase step count and move to the next instruction
        stepCount++
        instructionIndex = (instructionIndex + 1) % len(instructions)
    }

    // Output the result
    fmt.Println(stepCount)
}
