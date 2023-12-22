
package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    moves         = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // All possible moves.
    visitedPlots  = make(map[[2]int]bool)                        // Visited plots.
    startingPoint [2]int                                         // Starting position (S).
)

func readInput(filename string) ([][]rune, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var grid [][]rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        row := []rune(scanner.Text())
        grid = append(grid, row)
    }

    return grid, scanner.Err()
}

// findStartingPoint locates the starting point (S) on the grid.
func findStartingPoint(grid [][]rune) [2]int {
    for y, row := range grid {
        for x, tile := range row {
            if tile == 'S' {
                return [2]int{x, y}
            }
        }
    }
    return [2]int{}
}

// countSteps counts how many unique garden plots are reachable in exactly steps remaining.
func countSteps(grid [][]rune, location [2]int, stepsRemaining int) int {
    if stepsRemaining == 0 {
        // Mark the starting point as visited when counting starts.
        if _, found := visitedPlots[location]; !found {
            visitedPlots[location] = true
            return 1 // Counting the current plot as reachable.
        }
        return 0
    }

    count := 0
    x, y := location[0], location[1]

    // Check all possible moves.
    for _, move := range moves {
        newX, newY := x+move[0], y+move[1]

        // Check if within grid and not a rock.
        if newX >= 0 && newY >= 0 && newY < len(grid) && newX < len(grid[newY]) && grid[newY][newX] == '.' {
            newLocation := [2]int{newX, newY}
            if _, found := visitedPlots[newLocation]; !found {
                count += countSteps(grid, newLocation, stepsRemaining-1)
            }
        }
    }

    return count
}

func main() {
    grid, err := readInput("input.txt")
    if err != nil {
        panic(err)
    }

    startingPoint = findStartingPoint(grid)
    result := countSteps(grid, startingPoint, 64)
    fmt.Println(result)
}
