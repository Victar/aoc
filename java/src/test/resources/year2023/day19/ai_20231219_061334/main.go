
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

func parseCondition(cond string, ratings map[string]int) bool {
	if cond == "A" || cond == "R" || cond == "" {
		return true
	}
	parts := strings.SplitN(cond, ">", 2)
	if len(parts) == 1 {
		parts = strings.SplitN(cond, "<", 2)
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		return ratings[parts[0]] < value
	}
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return ratings[parts[0]] > value
}

func processPart(ratings map[string]int, workflows map[string]Workflow) (bool, int) {
	currentWorkflow := workflows["in"]
	for {
		for _, rule := range currentWorkflow.Rules {
			if parseCondition(rule.Condition, ratings) {
				if rule.Workflow == "A" {
					return true, ratings["x"] + ratings["m"] + ratings["a"] + ratings["s"]
				}
				if rule.Workflow == "R" {
					return false, 0
				}
				currentWorkflow = workflows[rule.Workflow]
				break
			}
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	workflows := make(map[string]Workflow)
	scanner := bufio.NewScanner(file)

	// Parse workflows
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}

		parts := strings.SplitN(line, "{", 2)
		name := parts[0]
		workflowRules := parts[1][:len(parts[1])-1]
		rules := strings.Split(workflowRules, ",")

		var parsedRules []Rule
		for _, rule := range rules {
			ruleParts := strings.SplitN(rule, ":", 2)
			var condition string
			var workflow string
			if len(ruleParts) == 2 {
				condition = ruleParts[0]
				workflow = ruleParts[1]
			} else {
				workflow = ruleParts[0]
			}
			parsedRules = append(parsedRules, Rule{Condition: condition, Workflow: workflow})
		}
		workflows[name] = Workflow{Name: name, Rules: parsedRules}
	}

	// Process parts
	sum := 0
	for scanner.Scan() {
		part := scanner.Text()
		part = strings.Trim(part, "{}")
		ratingsStr := strings.Split(part, ",")
		ratings := make(map[string]int)
		for _, rating := range ratingsStr {
			ratingParts := strings.SplitN(rating, "=", 2)
			category := ratingParts[0]
			value, err := strconv.Atoi(ratingParts[1])
			if err != nil {
				panic(err)
			}
			ratings[category] = value
		}

		accepted, partSum := processPart(ratings, workflows)
		if accepted {
			sum += partSum
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
