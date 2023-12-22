
// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	condition string
	workflow  string
}

type Workflow struct {
	rules []Rule
}

// Parse the workflows from the input and return a map of workflow name to workflow
func parseWorkflows(lines []string) map[string]Workflow {
	workflows := make(map[string]Workflow)
	for _, line := range lines {
		if line == "" {
			break
		}
		parts := strings.SplitN(line, "{", 2)
		name := parts[0]
		rulesStr := strings.TrimRight(parts[1], "}")
		rulesStrParts := strings.Split(rulesStr, ",")
		var rules []Rule
		for _, r := range rulesStrParts {
			ruleParts := strings.SplitN(r, ":", 2)
			var rule Rule
			if len(ruleParts) == 2 {
				rule = Rule{condition: ruleParts[0], workflow: ruleParts[1]}
			} else {
				rule = Rule{workflow: ruleParts[0]}
			}
			rules = append(rules, rule)
		}
		workflows[name] = Workflow{rules: rules}
	}
	return workflows
}

// Apply the conditions and workflows and return whether the part is accepted, and its ratings
func applyWorkflows(workflows map[string]Workflow, part string, name string) (bool, map[string]int) {
	for {
		workflow, present := workflows[name]
		if !present {
			break
		}
		accepted := false
		for _, rule := range workflow.rules {
			if rule.condition == "" {
				if rule.workflow == "A" {
					accepted = true
				}
				return accepted, parseRatings(part)
			}

			attribute := rule.condition[0:1]
			valueStr := rule.condition[2:len(rule.condition) - 2]
			value, _ := strconv.Atoi(valueStr)
			op := rule.condition[1:2]
			rating := parseRatings(part)[attribute]

			var match bool
			switch op {
			case ">":
				match = rating > value
			case "<":
				match = rating < value
			}

			if match {
				name = rule.workflow
				break
			}
		}
		if accepted || name == "R" {
			return false, nil
		}
	}
	return false, nil
}

// Parse the part ratings from the string
func parseRatings(part string) map[string]int {
	ratings := map[string]int{}
	part = strings.Trim(part, "{} ")
	ratingStrs := strings.Split(part, ",")
	for _, r := range ratingStrs {
		ratings[r[0:1]], _ = strconv.Atoi(r[2:])
	}
	return ratings
}

// Read the file and parse it into a list of lines
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// Read the input
	lines, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Parse workflows
	workflows := parseWorkflows(lines)

	sum := 0
	for _, part := range lines[len(workflows)+1:] {
		accepted, ratings := applyWorkflows(workflows, part, "in")
		if accepted {
			for _, rating := range ratings {
				sum += rating
			}
		}
	}
	fmt.Println(sum)
}
