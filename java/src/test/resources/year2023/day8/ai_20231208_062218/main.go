
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	networkMap := make(map[string][2]string)
	var pattern string

	// Read input and build the network map
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.Split(line, " = ")
			node := strings.TrimSpace(parts[0])
			connections := strings.Trim(parts[1], "()")
			nodes := strings.Split(connections, ", ")
			networkMap[node] = [2]string{strings.TrimSpace(nodes[0]), strings.TrimSpace(nodes[1])}
		} else {
			pattern = strings.TrimSpace(line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Part One: Navigate from AAA to ZZZ
	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		direction := string(pattern[steps%len(pattern)])
		nextIndex := 0
		if direction == "R" {
			nextIndex = 1
		}
		currentNode = networkMap[currentNode][nextIndex]
		steps++
	}
	fmt.Println(steps)

	// Part Two: Navigate simultaneously from nodes ending with A to Z
	nodes := make(map[string]struct{})
	for node := range networkMap {
		if strings.HasSuffix(node, "A") {
			nodes[node] = struct{}{}
		}
	}

	multiSteps := 0
	for {
		newNodes := make(map[string]struct{})
		for node := range nodes {
			if strings.HasSuffix(node, "Z") {
				newNodes[node] = struct{}{}
				continue
			}
			direction := string(pattern[multiSteps%len(pattern)])
			nextIndex := 0
			if direction == "R" {
				nextIndex = 1
			}
			nextNode := networkMap[node][nextIndex]
			newNodes[nextNode] = struct{}{}
		}
		nodes = newNodes

		allZ := true
		for node := range nodes {
			if !strings.HasSuffix(node, "Z") {
				allZ = false
				break
			}
		}

		if allZ {
			break
		}
		multiSteps++
	}
	fmt.Println(multiSteps)
}
