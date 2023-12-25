
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Hailstone struct {
    Position [3]int
    Velocity [3]int
}

func parseInput(filename string) ([]Hailstone, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var hailstones []Hailstone
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Each line in the format "px, py, pz @ vx, vy, vz"
        parts := strings.Split(scanner.Text(), "@")
        if len(parts) != 2 {
            return nil, fmt.Errorf("invalid input line format")
        }

        posStr := strings.Split(strings.TrimSpace(parts[0]), ",")
        velStr := strings.Split(strings.TrimSpace(parts[1]), ",")

        var pos [3]int
        var vel [3]int
        for i := 0; i < 3; i++ {
            pos[i], err = strconv.Atoi(strings.TrimSpace(posStr[i]))
            if err != nil {
                return nil, err
            }
            vel[i], err = strconv.Atoi(strings.TrimSpace(velStr[i]))
            if err != nil {
                return nil, err
            }
        }

        hailstones = append(hailstones, Hailstone{
            Position: pos,
            Velocity: vel,
        })
    }
    return hailstones, scanner.Err()
}

func main() {
    // Parse hailstones from the input file
    hailstones, err := parseInput("input.txt")
    if err != nil {
        fmt.Println("Error reading input:", err)
        os.Exit(1)
    }

    // Perform calculations for parts one and two.
    partOneAnswer := calculatePartOne(hailstones)
    partTwoAnswer := calculatePartTwo(hailstones)

    // Print the results as required.
    fmt.Println(partOneAnswer)
    fmt.Println(partTwoAnswer)
}

func calculatePartOne(hailstones []Hailstone) int {
    // Insert code to calculate the number of intersections
    // within the given bounds on the XY plane here.
    return 0
}

func calculatePartTwo(hailstones []Hailstone) int {
    // Insert code to determine the initial rock position and velocity.
    return 0
}
