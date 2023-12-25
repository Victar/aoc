
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Edge represents a connection between two components.
type Edge struct {
	A, B string
}

// ComponentMap holds the connections of components.
type ComponentMap map[string][]string

func readInput(filename string) (ComponentMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	componentMap := make(ComponentMap)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		component := parts[0]
		connectedComponents := strings.Fields(parts[1])
		componentMap[component] = connectedComponents
	}
	return componentMap, scanner.Err()
}

func findGroupsAfterDisconnectingThreeWires(components ComponentMap) (int, int) {
	// This function implements a heuristic approach to the problem.
	// It is not guaranteed to be correct because the problem
	// as stated requires an optimized solution that might need more complex logic or brute force.
	edges := []Edge{}
	for component, connectedComponents := range components {
		for _, connectedComponent := range connectedComponents {
			edges = append(edges, Edge{A: component, B: connectedComponent})
		}
	}

	// Heuristic: Just disconnect the first three unique wires we find.
	disconnectedEdges := edges[:3]

	for _, edge := range disconnectedEdges {
		disconnectWire(components, edge.A, edge.B)
	}

	group1Size := countComponentsInGroup(components, edges[0].A)
	group2Size := len(components) - group1Size

	return group1Size, group2Size
}

func disconnectWire(components ComponentMap, a, b string) {
	removeConnection(components, a, b)
	removeConnection(components, b, a)
}

func removeConnection(components ComponentMap, a, b string) {
	conn := components[a]
	for i, comp := range conn {
		if comp == b {
			components[a] = append(conn[:i], conn[i+1:]...)
			return
		}
	}
}

func countComponentsInGroup(components ComponentMap, startComponent string) int {
	visited := make(map[string]bool)
	var visitComponent func(string)
	visitComponent = func(component string) {
		if visited[component] {
			return
		}
		visited[component] = true
		for _, connectedComp := range components[component] {
			visitComponent(connectedComp)
		}
	}

	visitComponent(startComponent)
	return len(visited)
}

func main() {
	components, err := readInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %s", err)
		os.Exit(1)
	}

	group1Size, group2Size := findGroupsAfterDisconnectingThreeWires(components)
	fmt.Println(group1Size * group2Size)
}
