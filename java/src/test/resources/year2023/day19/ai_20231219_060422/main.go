package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Condition string
	Next      string
}

type Workflow struct {
	Name  string
	Rules []Rule
}

func main() {
	workflows, parts := readInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day19/sample.txt")
	acceptedSum := 0
	for _, part := range parts {
		wf := "in" // Start with "in" workflow as per the task description
		for wf != "A" && wf != "R" {
			wf = processPart(workflows[wf], part) // Process part through current workflow
		}
		if wf == "A" {
			acceptedSum += partSum(part) // Add ratings if accepted
		}
	}
	println(acceptedSum)
}

func readInput(filename string) (map[string]Workflow, []map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	workflows := make(map[string]Workflow)
	var parts []map[string]int
	readingParts := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingParts = true
			continue
		}
		if !readingParts {
			wf := parseWorkflow(line)
			workflows[wf.Name] = wf
		} else {
			part := parsePart(line)
			parts = append(parts, part)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return workflows, parts
}

func parseWorkflow(line string) Workflow {
	parts := strings.Split(strings.Trim(line, "{}"), "{")
	name := parts[0]
	rulesStr := strings.Split(parts[1], ",")

	var rules []Rule
	for _, ruleStr := range rulesStr {
		ruleParts := strings.Split(ruleStr, ":")
		condition := ""
		if len(ruleParts) == 2 {
			condition = ruleParts[0]
		}
		rules = append(rules, Rule{condition, ruleParts[len(ruleParts)-1]})
	}

	return Workflow{name, rules}
}

func parsePart(line string) map[string]int {
	part := make(map[string]int)
	line = strings.Trim(line, "{}")
	ratings := strings.Split(line, ",")

	for _, rating := range ratings {
		ratingParts := strings.Split(rating, "=")
		value, err := strconv.Atoi(ratingParts[1])
		if err != nil {
			panic(err)
		}
		part[ratingParts[0]] = value
	}
	return part
}

func checkCondition(condition string, part map[string]int) bool {
	if condition == "" {
		return true
	}
	condParts := strings.Split(condition, ">")
	if len(condParts) == 2 {
		value, err := strconv.Atoi(condParts[1])
		if err != nil {
			panic(err)
		}
		return part[condParts[0]] > value
	}

	condParts = strings.Split(condition, "<")
	if len(condParts) == 2 {
		value, err := strconv.Atoi(condParts[1])
		if err != nil {
			panic(err)
		}
		return part[condParts[0]] < value
	}
	return false
}

func processPart(wf Workflow, part map[string]int) string {
	for _, rule := range wf.Rules {
		if checkCondition(rule.Condition, part) {
			return rule.Next
		}
	}
	return "A"
}

func partSum(part map[string]int) int {
	sum := 0
	for _, value := range part {
		sum += value
	}
	return sum
}
