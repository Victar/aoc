
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	workflows, parts := parseInput("input.txt")
	sum := processAndSum(workflows, parts)
	fmt.Println(sum)
}

type Rule struct {
	Condition string
	Next      string
}

func parseInput(filename string) (map[string][]Rule, []map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	workflows := make(map[string][]Rule)
	parts := []map[string]int{}

	parsingWorkflows := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingWorkflows = false
			continue
		}

		if parsingWorkflows {
			components := strings.Split(line, "{")
			workflowName := components[0]
			ruleDefs := components[1][:len(components[1])-1]
			rules := strings.Split(ruleDefs, ",")

			for _, rule := range rules {
				parts := strings.Split(rule, ":")
				var condition string
				destination := parts[1]
				if len(parts) == 2 && parts[0] != "A" && parts[0] != "R" {
					condition = parts[0]
				}
				r := Rule{Condition: condition, Next: destination}
				workflows[workflowName] = append(workflows[workflowName], r)
			}
		} else {
			part := map[string]int{}
			ratings := strings.Trim(line, "{}")
			for _, pair := range strings.Split(ratings, ",") {
				kv := strings.Split(pair, "=")
				key := kv[0]
				value, _ := strconv.Atoi(kv[1])
				part[key] = value
			}
			parts = append(parts, part)
		}
	}
	return workflows, parts
}

func evaluateCondition(cond string, part map[string]int) bool {
	if cond == "" {
		return true
	}
	parts := strings.Split(cond, ">")
	if len(parts) < 2 {
		parts = strings.Split(cond, "<")
		attr := parts[0]
		value, _ := strconv.Atoi(parts[1])
		return part[attr] < value
	} else {
		attr := parts[0]
		value, _ := strconv.Atoi(parts[1])
		return part[attr] > value
	}
}

func processAndSum(workflows map[string][]Rule, parts []map[string]int) int {
	sum := 0
	for _, part := range parts {
		currentWorkflow := "in"
		accepted := false
		for {
			if currentWorkflow == "A" {
				sum += part["x"] + part["m"] + part["a"] + part["s"]
				accepted = true
				break
			} else if currentWorkflow == "R" {
				break
			}
			rules := workflows[currentWorkflow]
			for _, rule := range rules {
				if evaluateCondition(rule.Condition, part) {
					currentWorkflow = rule.Next
					break
				}
			}
		}
		if !accepted {
			continue
		}
	}
	return sum
}

func main() {
	workflows, parts := parseInput("input.txt")
	fmt.Println(processAndSum(workflows, parts))
}
