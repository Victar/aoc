
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph map[string][]string

func (g Graph) Disconnect(compA, compB string) {
	// Remove compB from compA's list
	for i, comp := range g[compA] {
		if comp == compB {
			g[compA] = append(g[compA][:i], g[compA][i+1:]...)
			break
		}
	}

	// Remove compA from compB's list
	for i, comp := range g[compB] {
		if comp == compA {
			g[compB] = append(g[compB][:i], g[compB][i+1:]...)
			break
		}
	}
}

func (g Graph) DFS(start string, visited map[string]bool) {
	visited[start] = true
	for _, neighbor := range g[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

func countComponents(g Graph) int {
	visited := make(map[string]bool)
	count := 0
	for node := range g {
		if !visited[node] {
			count++
			g.DFS(node, visited)
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer file.Close()

	graph := make(Graph)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		node := parts[0]
		neighbors := strings.Fields(parts[1])
		graph[node] = neighbors
		for _, neighbor := range neighbors {
			graph[neighbor] = append(graph[neighbor], node)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}

	// Try disconnecting every pair of connected components
	bestMultiplication := -1

	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			// Disconnect one pair of components
			tmpGraph := make(Graph)
			for k, v := range graph {
				tmpGraph[k] = append([]string(nil), v...)
			}
			tmpGraph.Disconnect(node, neighbor)

			// Find the groups after disconnection
			visited := make(map[string]bool)
			var groupSizes []int
			for comp := range tmpGraph {
				if !visited[comp] {
					size := 0
					tmpGraph.DFS(comp, visited)
					for _, checked := range visited {
						if checked {
							size++
						}
					}
					if size > 0 {
						groupSizes = append(groupSizes, size)
					}
				}
			}

			if len(groupSizes) == 2 && (bestMultiplication == -1 || bestMultiplication > groupSizes[0]*groupSizes[1]) {
				bestMultiplication = groupSizes[0] * groupSizes[1]
			}
		}
	}

	if bestMultiplication != -1 {
		fmt.Println(bestMultiplication)
	} else {
		fmt.Println("No solution found.")
	}
}
