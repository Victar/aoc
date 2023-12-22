package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "19"

type Rule struct {
	Condition string
	Next      string
}

type Workflow struct {
	Name  string
	Rules []Rule
}
type State struct {
	state string
	r     Range
}

type PartRange struct {
	start, end int
}

type Range struct {
	minX int
	maxX int
	minM int
	maxM int
	minA int
	maxA int
	minS int
	maxS int
}

func main() {
	//runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	workflows := make(map[string]Workflow)
	var parts []map[string]int
	readingParts := false
	for _, line := range lines {
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

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/sample.txt")
	if err != nil {
		panic(err)
	}
	workflows := make(map[string]Workflow)
	var parts []map[string]int
	readingParts := false
	for _, line := range lines {
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
	start := Range{0, 4000, 0, 4000, 0, 4000, 0, 4000}
	fmt.Println(workflows, countRange(start))

}

func countRange(r Range) int {
	return (r.maxX - r.minX) * (r.maxM - r.minM) * (r.maxA - r.minA) * (r.maxS - r.minS)
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
