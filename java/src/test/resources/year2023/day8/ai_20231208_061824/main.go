
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left, right string
}

func main() {
	// Open the input file.
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read instructions line
	instructions := strings.Split(scanner.Text(), "")

	// Map to hold nodes.
	nodes := make(map[string]Node)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 5 {
			name, left, right := parts[0], parts[2], parts[4]
			nodes[name] = Node{left, right}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Find all nodes that end with "A" to start with.
	startingNodes := getAllStartingNodes(nodes)

	// Simulate all the movements from the starting nodes.
	count := simulateMovements(startingNodes, instructions, nodes)

	fmt.Println(count)
}

func simulateMovements(startingNodes []string, instructions []string, nodes map[string]Node) int {
	count, idx := 0, 0
	positions := make(map[string]bool)

	// Initialize positions with starting nodes.
	for _, node := range startingNodes {
		positions[node] = true
	}

	for len(positions) > 0 {
		instruction := instructions[idx%len(instructions)]
		newPositions := make(map[string]bool)

		for position := range positions {
			left, right := nodes[position].left, nodes[position].right
			nextNode := right
			if instruction == "L" {
				nextNode = left
			}

			if strings.HasSuffix(nextNode, "Z") {
				newPositions[nextNode] = true
			} else {
				newPositions[nextNode] = true
			}
		}

		positions = newPositions
		idx++
		if len(newPositions) == 1 {
			// If only nodes ending with Z are left, we stop counting.
			allZ := true
			for pos := range newPositions {
				if !strings.HasSuffix(pos, "Z") {
					allZ = false
					break
				}
			}
			if allZ {
				break
			}
		}
		count++
	}

	return count
}

func getAllStartingNodes(nodes map[string]Node) []string {
	startingNodes := []string{}
	for nodeName := range nodes {
		if strings.HasSuffix(nodeName, "A") {
			startingNodes = append(startingNodes, nodeName)
		}
	}
	return startingNodes
}
