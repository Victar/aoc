
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Reads a line of input and converts it into a slice of integers.
func readHistory(input string) ([]int, error) {
    parts := strings.Split(strings.TrimSpace(input), " ")
    history := make([]int, len(parts))
    for i, part := range parts {
        num, err := strconv.Atoi(part)
        if err != nil {
            return nil, err
        }
        history[i] = num
    }
    return history, nil
}

// Generates the sequence of differences for a given history.
func generateDifferences(history []int) []int {
    diffs := make([]int, len(history)-1)
    for i := 0; i < len(history)-1; i++ {
        diffs[i] = history[i+1] - history[i]
    }
    return diffs
}

// Generates all sequences of differences until we reach a sequence of all zeroes.
// Returns the last non-zero differences and the number of iterations it took to reach zeroes.
func generateAllDifferences(history []int) ([]int, int) {
    current := history
    generation := 0
    for {
        diffs := generateDifferences(current)
        allZero := true
        for _, diff := range diffs {
            if diff != 0 {
                allZero = false
                break
            }
        }
        if allZero {
            return current, generation
        }
        current = diffs
        generation++
    }
}

// Extrapolates the next or previous value based on the sequence of differences.
func extrapolateValue(diffs []int, iterations int, next bool) int {
    for i := 0; i < iterations; i++ {
        newDiffs := make([]int, len(diffs)+1)
        if next {
            copy(newDiffs[1:], diffs)
        } else {
            copy(newDiffs, diffs)
        }
        diffs = newDiffs

        for j := len(diffs) - 2; j >= 0; j-- {
            diffs[j] = diffs[j] + diffs[j+1]
        }
    }
    if next {
        return diffs[len(diffs)-1]
    } else {
        return diffs[0]
    }
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sumNext := 0
    sumPrevious := 0
    for scanner.Scan() {
        history, err := readHistory(scanner.Text())
        if err != nil {
            fmt.Println("Error reading history:", err)
            return
        }

        diffs, gens := generateAllDifferences(history)
        nextValue := extrapolateValue(diffs, gens, true)
        sumNext += nextValue

        previousValue := extrapolateValue(diffs, gens, false)
        sumPrevious += previousValue
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
        return
    }

    // Uncomment the lines below to see the individual sums
    // fmt.Println("Sum of next values:", sumNext)
    // fmt.Println("Sum of previous values:", sumPrevious)

    // The task asks for the sum of extrapolated values for part one.
    fmt.Println(sumNext) // Should print the sum for part one.
    // If you need to print the sum for part two, uncomment the line below.
    // fmt.Println(sumPrevious) // Should print the sum for part two.
}
