
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Condition string
	Workflow  string
}

type Workflow struct {
	Name  string
	Rules []Rule
}

func main() {
	// Placeholder to store workflows
	workflows := make(map[string]Workflow)

	// Read file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	processingRules := true
	acceptedPartsSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			processingRules = false
			continue
		}

		if processingRules {
			processWorkflowLine(line, workflows)
		} else {
			acceptedPartsSum += processPartLine(line, workflows)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(acceptedPartsSum)
}

func processWorkflowLine(line string, workflows map[string]Workflow) {
	// TODO: Implement parsing of rule line and insertion into 'workflows' map
}

func processPartLine(line string, workflows map[string]Workflow) int {
	// TODO: Implement processing of part line based on 'workflows' rules
	parts := parsePartLine(line)
	return evaluatePart(workflows, parts, "in")
}

func parsePartLine(line string) map[string]int {
	parts := make(map[string]int)
	values := strings.Split(strings.Trim(line, "{}"), ",")
	for _, v := range values {
		split := strings.Split(v, "=")
		if len(split) == 2 {
			key := strings.TrimSpace(split[0])
			value, _ := strconv.Atoi(split[1])
			parts[key] = value
		}
	}
	return parts
}

func evaluatePart(workflows map[string]Workflow, parts map[string]int, currentWorkflow string) int {
	// TODO: Implement evaluation logic for part based on the current workflow
	// This should recursively evaluate the part through workflows
	// and return the sum of accepted parts ratings or 0 if rejected.
	return 0
}

// NOTE: Completing the TODO sections will require implementing proper rule parsing and evaluating a part's values against the rules.
