
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Parse input file
	nodes := make(map[string]Node)
	var instructions string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.Split(line, " = ")
			nodeName := strings.TrimSpace(parts[0])
			children := strings.Split(parts[1][1:len(parts[1])-1], ", ")
			nodes[nodeName] = Node{left: children[0], right: children[1]}
		} else {
			instructions = strings.TrimSpace(line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Part 1 - How many steps are required to reach ZZZ?
	steps := 0
	current := "AAA"
	idx := 0
	for current != "ZZZ" {
		if instructions[idx] == 'L' {
			current = nodes[current].left
		} else {
			current = nodes[current].right
		}
		idx = (idx + 1) % len(instructions)
		steps++
	}
	fmt.Println(steps)

	// Part 2 - How many steps does it take before you're only on nodes that end with Z?
	steps = 0
	currentNodes := make(map[string]bool)
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			currentNodes[node] = true
		}
	}
	finished := false
	for !finished {
		nextNodes := make(map[string]bool)
		idx = steps % len(instructions)
		instruction := instructions[idx]

		for node := range currentNodes {
			nextNode := node
			if instruction == 'L' {
				nextNode = nodes[node].left
			} else if instruction == 'R' {
				nextNode = nodes[node].right
			}

			// Check if next node ends with Z
			if strings.HasSuffix(nextNode, "Z") {
				nextNodes[nextNode] = true
			} else if strings.HasSuffix(nextNode, "A") || strings.HasSuffix(nextNode, "B") || strings.HasSuffix(nextNode, "C") {
				// If new nodes are encountered, add them to the list
				nextNodes[nextNode] = true
			}
		}

		finished = true
		for node := range nextNodes {
			if !strings.HasSuffix(node, "Z") {
				finished = false
				break
			}
		}
		steps++
		currentNodes = nextNodes
	}
	fmt.Println(steps)
}
