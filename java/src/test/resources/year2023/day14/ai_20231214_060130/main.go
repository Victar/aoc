
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer input.Close()

    var lines []string
    scanner := bufio.NewScanner(input)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err = scanner.Err(); err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    totalLoad := 0
    rowCount := len(lines)

    // Starting from the bottom row, add the load for each rounded rock.
    for rowIndex, line := range lines {
        distanceFromSouth := rowCount - rowIndex
        for _, char := range line {
            if char == 'O' {
                totalLoad += distanceFromSouth
            }
        }
    }

    fmt.Println(totalLoad)
}
