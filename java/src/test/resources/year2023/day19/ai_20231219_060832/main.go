
// We are assuming the input file's syntax is well-formed.
// Parsing and error checking are minimal due to the constraints of the problem statement.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Condition func(map[string]int) bool
	Next      string
}

type Workflow struct {
	Name  string
	Rules []Rule
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer file.Close()

	workflows := make(map[string][]Rule)
	var parts []map[string]int

	parseCondition := func(cond string) func(map[string]int) bool {
		if cond == "A" {
			return nil // no condition; the rule should always apply
		}
		re := regexp.MustCompile(`([xmas])\s*([<=>]+)\s*(\d+)`)
		matches := re.FindStringSubmatch(cond)
		attribute, comparator, valueStr := matches[1], matches[2], matches[3]
		value, _ := strconv.Atoi(valueStr)

		return func(part map[string]int) bool {
			partValue := part[attribute]
			switch comparator {
			case "<":
				return partValue < value
			case "<=":
				return partValue <= value
			case ">":
				return partValue > value
			case ">=":
				return partValue >= value
			case "=":
				return partValue == value
			case "!=":
				return partValue != value
			default:
				return false
			}
		}
	}

	scanner := bufio.NewScanner(file)
	parsingWorkflows := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingWorkflows = false
			continue
		}

		if parsingWorkflows {
			parts := strings.SplitN(line, "{", 2)
			name := parts[0]
			ruleParts := strings.Trim(parts[1], "}")
			rules := strings.Split(ruleParts, ",")
			var workflowRules []Rule

			for _, r := range rules {
				conditionAndNext := strings.Split(r, ":")
				var rule Rule
				if len(conditionAndNext) == 2 {
					cond := conditionAndNext[0]
					rule = Rule{
						Condition: parseCondition(cond),
						Next:      conditionAndNext[1],
					}
				} else if len(conditionAndNext) == 1 && conditionAndNext[0] == "A" {
					// Accept unconditionally
					rule = Rule{
						Condition: nil,
						Next:      "A",
					}
				} else {
					rule = Rule{
						Condition: nil,
						Next:      "R",
					}
				}
				workflowRules = append(workflowRules, rule)
			}
			workflows[name] = workflowRules

		} else {
			partRatingStr := strings.Trim(line, "{}")
			partRatings := strings.Split(partRatingStr, ",")
			part := make(map[string]int)
			for _, r := range partRatings {
				ratingParts := strings.Split(r, "=")
				attr := ratingParts[0]
				value, _ := strconv.Atoi(ratingParts[1])
				part[attr] = value
			}
			parts = append(parts, part)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	total := 0
	for _, part := range parts {
		nextWorkflowName := "in"
		for nextWorkflowName != "A" && nextWorkflowName != "R" {
			workflow := workflows[nextWorkflowName]
			matched := false
			for _, rule := range workflow {
				if rule.Condition == nil || rule.Condition(part) {
					nextWorkflowName = rule.Next
					matched = true
					break
				}
			}
			if !matched {
				break // This means that none of the rules applied, which should not happen
			}
		}

		if nextWorkflowName == "A" {
			for _, value := range part {
				total += value
			}
		}
	}
	fmt.Println(total)
}
