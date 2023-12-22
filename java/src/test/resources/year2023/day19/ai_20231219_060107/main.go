package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Workflow struct {
	Rules []Rule
}

type Rule struct {
	Condition string
	Action    string
}

type Part struct {
	X int
	M int
	A int
	S int
}

func main() {
	workflows, parts := parseInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day19/input.txt")
	sum := 0
	for _, part := range parts {
		if accepted := processPart(part, "in", workflows); accepted {
			sum += part.X + part.M + part.A + part.S
		}
	}
	fmt.Println(sum)
}

func parseInput(filename string) (map[string]Workflow, []Part) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	workflows := make(map[string]Workflow)
	readingWorkflows := true
	var parts []Part

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingWorkflows = false
			continue
		}

		if readingWorkflows {
			workflowName := strings.Split(line, "{")[0]
			workflowContent := strings.Trim(strings.Split(line, "{")[1], "}")
			workflow := Workflow{parseRules(workflowContent)}
			workflows[workflowName] = workflow
		} else {
			part := parsePart(line)
			parts = append(parts, part)
		}
	}

	return workflows, parts
}

func parseRules(workflowContent string) []Rule {
	ruleStrs := strings.Split(workflowContent, ",")
	var rules []Rule
	for _, rs := range ruleStrs {
		split := strings.Split(rs, ":")
		rules = append(rules, Rule{
			Condition: split[0],
			Action:    split[len(split)-1],
		})
	}
	return rules
}

func parsePart(line string) Part {
	vals := strings.Trim(strings.Split(line, "{")[1], "}")
	vals = strings.Trim(vals, "}")
	valsSplit := strings.Split(vals, ",")

	getValue := func(s string) int {
		val, _ := strconv.Atoi(strings.Split(s, "=")[1])
		return val
	}

	part := Part{
		X: getValue(valsSplit[0]),
		M: getValue(valsSplit[1]),
		A: getValue(valsSplit[2]),
		S: getValue(valsSplit[3]),
	}

	return part
}

func processPart(part Part, workflowName string, workflows map[string]Workflow) bool {
	workflow := workflows[workflowName]
	for _, rule := range workflow.Rules {
		if rule.Condition == "A" {
			return true
		} else if rule.Condition == "R" {
			return false
		}

		if matchesCondition(part, rule.Condition) {
			if rule.Action == "A" {
				return true
			} else if rule.Action == "R" {
				return false
			}

			return processPart(part, rule.Action, workflows)
		}
	}
	return false
}

func matchesCondition(part Part, condition string) bool {
	if condition == "" {
		return true
	}

	var value int
	switch condition[0] {
	case 'x':
		value = part.X
	case 'm':
		value = part.M
	case 'a':
		value = part.A
	case 's':
		value = part.S
	}

	comparator := condition[1]
	condVal, _ := strconv.Atoi(condition[2:])

	switch comparator {
	case '>':
		return value > condVal
	case '<':
		return value < condVal
	}

	return false
}
