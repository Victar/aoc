
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input data from file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Data structures to hold the input directions and the nodes
	var instructionSet []string
	nodeMap := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			tokens := strings.Split(line, " = ")
			node := strings.TrimSpace(tokens[0])
			connections := strings.Split(tokens[1][1:len(tokens[1])-1], ", ")
			nodeMap[node] = []string{strings.TrimSpace(connections[0]), strings.TrimSpace(connections[1])}
		} else if line != "" {
			// Assuming that the line with no '=' character is the set of instructions
			instructionSet = strings.Split(line, "")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	currentNode := "AAA"
	stepCount := 0

	for currentNode != "ZZZ" {
		direction := instructionSet[stepCount%len(instructionSet)]
		nextIndex := 0
		if direction == "R" {
			nextIndex = 1
		}

		currentNode = nodeMap[currentNode][nextIndex]
		stepCount++
	}

	// Print the number of steps required to reach ZZZ
	fmt.Println(stepCount)
}
