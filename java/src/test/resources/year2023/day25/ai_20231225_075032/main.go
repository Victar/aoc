
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Structure to represent a graph
type Graph map[string][]string

// ReadInput reads the input from the provided file
func ReadInput(filename string) (Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := Graph{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		component := parts[0]
		connected := strings.Fields(parts[1])
		graph[component] = append(graph[component], connected...)
		for _, conn := range connected {
			graph[conn] = append(graph[conn], component)
		}
	}

	return graph, scanner.Err()
}

// DFS performs a depth first search on the graph
func DFS(graph Graph, start string, visited map[string]bool) int {
	visited[start] = true
	size := 1 // Start with 1 to count the current component

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			size += DFS(graph, neighbor, visited)
		}
	}

	return size
}

// Disconnect removes an edge from the graph
func Disconnect(g Graph, node1, node2 string) {
	// Remove node2 from node1's adjacency list
	for i, neighbor := range g[node1] {
		if neighbor == node2 {
			g[node1] = append(g[node1][:i], g[node1][i+1:]...)
			break
		}
	}
	// Remove node1 from node2's adjacency list
	for i, neighbor := range g[node2] {
		if neighbor == node1 {
			g[node2] = append(g[node2][:i], g[node2][i+1:]...)
			break
		}
	}
}

// FindMaximumSplit tries to find the best split of the graph into two groups
func FindMaximumSplit(g Graph) int {
	maxProduct := 0
	var bestDisconnect [3][2]string

	// Try disconnecting each possible pair of nodes
	for node1, neighbors := range g {
		for _, node2 := range neighbors {
			// Disconnect the nodes
			Disconnect(g, node1, node2)

			// Find the sizes of the resulting disconnected graphs
			visited := make(map[string]bool)
			size1 := DFS(g, node1, visited)

			// The complement of size1 in the graph gives size2
			size2 := len(g) - size1
			product := size1 * size2

			// Check if this is the best split found so far
			if product > maxProduct {
				maxProduct = product
				bestDisconnect = [3][2]string{{node1, node2}}
			} else if product == maxProduct && len(bestDisconnect) < 3 {
				bestDisconnect = append(bestDisconnect, [2]string{node1, node2})
			}

			// Reconnect the nodes for the next iteration
			g[node1] = append(g[node1], node2)
			g[node2] = append(g[node2], node1)
		}
	}

	for _, disconnect := range bestDisconnect {
		fmt.Println("Disconnect:", disconnect[0], "/", disconnect[1])
	}

	return maxProduct
}

func main() {
	graph, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := FindMaximumSplit(graph)
	fmt.Println(result)
}
