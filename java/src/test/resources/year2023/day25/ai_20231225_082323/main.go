
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read wiring diagram from input.txt
	components, err := readComponentsFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Find the sizes of the two groups
	group1, group2 := disconnectAndFindGroups(components)
	fmt.Println(group1 * group2)
}

// Component structure
type Component struct {
	Name        string
	Connections map[string]*Component
}

// readComponentsFromFile reads the wiring diagram from the input file and returns a map of components
func readComponentsFromFile(filename string) (map[string]*Component, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	components := make(map[string]*Component)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		name := parts[0]
		connections := strings.Split(parts[1], " ")

		if _, exists := components[name]; !exists {
			components[name] = &Component{Name: name, Connections: make(map[string]*Component)}
		}

		for _, connection := range connections {
			if _, exists := components[connection]; !exists {
				components[connection] = &Component{Name: connection, Connections: make(map[string]*Component)}
			}
			components[name].Connections[connection] = components[connection]
			components[connection].Connections[name] = components[name]
		}
	}

	return components, scanner.Err()
}

// disconnectAndFindGroups disconnects components and calculates group sizes
func disconnectAndFindGroups(components map[string]*Component) (int, int) {
	var maxGroupSize int = len(components) / 2
	var bestSplit [3][2]string
	var bestGroup1Size, bestGroup2Size int = 0, len(components)

	// This is a naive approach; a more efficient solution may be possible
	for c1 := range components {
		for c2 := range components[c1].Connections {
			// disconnect c1 and c2
			disconnectedComponent := components[c1].Connections[c2]
			delete(components[c1].Connections, c2)
			delete(disconnectedComponent.Connections, c1)

			for c3 := range components {
				for c4 := range components[c3].Connections {
					if c3 == c1 && c4 == c2 {
						continue // Skip the first disconnected components
					}
					// disconnect c3 and c4
					disconnectedComponent := components[c3].Connections[c4]
					delete(components[c3].Connections, c4)
					delete(disconnectedComponent.Connections, c3)

					for c5 := range components {
						for c6 := range components[c5].Connections {
							if (c5 == c1 && c6 == c2) || (c5 == c3 && c6 == c4) {
								continue // Skip already disconnected components
							}
							// disconnect c5 and c6
							disconnectedComponent := components[c5].Connections[c6]
							delete(components[c5].Connections, c6)
							delete(disconnectedComponent.Connections, c5)

							group1, group2 := calculateGroups(components)
							if abs(group1-group2) < abs(bestGroup1Size-bestGroup2Size) {
								bestGroup1Size, bestGroup2Size = group1, group2
								bestSplit = [3][2]string{{c1, c2}, {c3, c4}, {c5, c6}}
							}

							// reconnect c5 and c6
							components[c5].Connections[c6] = disconnectedComponent
							disconnectedComponent.Connections[c5] = components[c5]
						}
					}

					// reconnect c3 and c4
					components[c3].Connections[c4] = disconnectedComponent
					disconnectedComponent.Connections[c3] = components[c3]
				}
			}

			// reconnect c1 and c2
			components[c1].Connections[c2] = disconnectedComponent
			disconnectedComponent.Connections[c1] = components[c1]
		}
	}

	return bestGroup1Size, bestGroup2Size
}

// calculateGroups calculates the sizes of two groups by doing a DFS on the graph
func calculateGroups(components map[string]*Component) (int, int) {
	visited := make(map[string]bool)
	var groupSize int

	for _, component := range components {
		if visited[component.Name] {
			continue
		}
		groupSize = 0
		dfs(component, visited, &groupSize)
		break // Only need the size of the first group; the other is len(components) - groupSize
	}

	return groupSize, len(components) - groupSize
}

func dfs(component *Component, visited map[string]bool, groupSize *int) {
	if visited[component.Name] {
		return
	}

	visited[component.Name] = true
	*groupSize++

	for _, connectedComponent := range component.Connections {
		dfs(connectedComponent, visited, groupSize)
	}
}

// abs calculates the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
