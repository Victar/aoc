
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

type Rule struct {
    condition string
    next      string
}

type Workflow struct {
    rules []Rule
}

func main() {
    workflows, parts := readInput("input.txt")
    sum := 0
    for _, part := range parts {
        if accepted, ratings := processPart(workflows, part); accepted {
            sum += sumRatings(ratings)
        }
    }
    fmt.Println(sum)
}

func processPart(workflows map[string]Workflow, part map[string]int) (bool, map[string]int) {
    currentWorkflow := workflows["in"]
    for {
        accepted, rejected, nextWorkflow := applyRules(currentWorkflow, part)
        if accepted {
            return true, part
        }
        if rejected {
            return false, nil
        }
        currentWorkflow = workflows[nextWorkflow]
    }
}

func applyRules(workflow Workflow, part map[string]int) (accepted bool, rejected bool, nextWorkflow string) {
    for _, rule := range workflow.rules {
        if rule.condition == "" {
            if rule.next == "A" {
                return true, false, ""
            } else if rule.next == "R" {
                return false, true, ""
            }
        } else {
            matches := regexp.MustCompile(`([xmas])>?=?<?(\d+)*`).FindStringSubmatch(rule.condition)
            attribute, valueStr := matches[1], matches[2]
            conditionValue, _ := strconv.Atoi(valueStr)
            partValue := part[attribute]

            if rule.condition[1] == '<' && partValue < conditionValue {
                return false, false, rule.next
            } else if rule.condition[1] == '>' && partValue > conditionValue {
                return false, false, rule.next
            }
        }
    }
    return false, false, ""
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

    readingWorkflows := true
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            readingWorkflows = false
            continue
        }

        if readingWorkflows {
            name, rules := parseWorkflow(line)
            workflows[name] = rules
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

func parseWorkflow(line string) (string, Workflow) {
    parts := regexp.MustCompile(`(\w+)\{(.*)\}`).FindStringSubmatch(line)
    if len(parts) < 3 {
        panic(fmt.Sprintf("Invalid workflow format: %s", line))
    }

    name := parts[1]
    rawRules := parts[2]
    ruleStrs := regexp.MustCompile(`,`).Split(rawRules, -1)

    var rules []Rule
    for _, ruleStr := range ruleStrs {
        if ruleStr == "A" || ruleStr == "R" {
            rules = append(rules, Rule{"", ruleStr})
        } else {
            parts := regexp.MustCompile(`([xmas><=\d]+):(\w+)`).FindStringSubmatch(ruleStr)
            rules = append(rules, Rule{parts[1], parts[2]})
        }
    }

    return name, Workflow{rules}
}

func parsePart(line string) map[string]int {
    matches := regexp.MustCompile(`\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}`).FindStringSubmatch(line)
    if len(matches) < 5 {
        panic(fmt.Sprintf("Invalid part format: %s", line))
    }

    x, _ := strconv.Atoi(matches[1])
    m, _ := strconv.Atoi(matches[2])
    a, _ := strconv.Atoi(matches[3])
    s, _ := strconv.Atoi(matches[4])

    return map[string]int{
        "x": x,
        "m": m,
        "a": a,
        "s": s,
    }
}

func sumRatings(ratings map[string]int) int {
    sum := 0
    for _, value := range ratings {
        sum += value
    }
    return sum
}
