
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    grid := strings.Split(strings.TrimSpace(string(data)), "\n")
    fmt.Println(countEnergizedTiles(grid))
}

func countEnergizedTiles(grid []string) int {
    energizedTilesCount := 0
    energizedTiles := make(map[[2]int]struct{})

    var traverse func(x, y, dx, dy int)
    traverse = func(x, y, dx, dy int) {
        for x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
            switch grid[y][x] {
            case '.':
                // If already energized, just continue
                if _, exists := energizedTiles[[2]int{x, y}]; exists {
                    x += dx
                    y += dy
                    continue
                }
                energizedTiles[[2]int{x, y}] = struct{}{}
                energizedTilesCount++
                x += dx
                y += dy
            case '/':
                dx, dy = dy, dx
                x += dx
                y += dy
            case '\\':
                dx, dy = -dy, -dx
                x += dx
                y += dy
            case '|', '-':
                pos := [2]int{x, y}
                if _, exists := energizedTiles[pos]; !exists {
                    energizedTiles[pos] = struct{}{}
                    energizedTilesCount++
                }

                if (grid[y][x] == '|' && dx != 0) || (grid[y][x] == '-' && dy != 0) {
                    // Split the beam into two
                    traverse(x, y, 0, 1)
                    traverse(x, y, 0, -1)
                    return
                }

                // If the beam hits the pointy end, treat it like empty space
                x += dx
                y += dy
            default:
                // Unrecognized character, end the beam.
                return
            }
        }
    }

    // Start the beam going right from the top-left corner
    traverse(0, 0, 1, 0)
    return energizedTilesCount
}
