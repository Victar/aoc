
// main.go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    var waysToWin []int

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        if len(parts) == 4 && parts[0] == "Time:" && parts[2] == "Distance:" {
            time, errTime := strconv.Atoi(parts[1])
            distance, errDistance := strconv.Atoi(parts[3])
            if errTime != nil || errDistance != nil {
                fmt.Println("Error parsing input values:", errTime, errDistance)
                return
            }
            waysToWin = append(waysToWin, calculateWays(time, distance))
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
        return
    }

    // Multiply all the ways to win.
    product := 1
    for _, ways := range waysToWin {
        product *= ways
    }

    fmt.Println(product)
}

// calculateWays calculates the number of ways to beat the record in a race.
func calculateWays(time, recordDistance int) int {
    ways := 0
    for hold := 1; hold < time; hold++ {
        distanceCovered := hold * (time - hold) // speed * remaining time
        if distanceCovered > recordDistance {
            ways++
        }
    }
    return ways
}
