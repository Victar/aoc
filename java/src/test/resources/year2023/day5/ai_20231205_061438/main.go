
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type RangeMap map[int]int

func parseMap(line string) RangeMap {
    rm := make(RangeMap)
    nums := strings.Split(line, " ")
    for i := 0; i < len(nums); i += 3 {
        destStart, _ := strconv.Atoi(nums[i])
        srcStart, _ := strconv.Atoi(nums[i+1])
        length, _ := strconv.Atoi(nums[i+2])
        for j := 0; j < length; j++ {
            rm[srcStart+j] = destStart + j
        }
    }
    return rm
}

func mapValue(rm RangeMap, value int) int {
    if dest, ok := rm[value]; ok {
        return dest
    }
    return value
}

func parseSeedRanges(line string) []int {
    nums := strings.Fields(line)
    var seeds []int
    for i := 0; i < len(nums); i += 2 {
        start, _ := strconv.Atoi(nums[i])
        length, _ := strconv.Atoi(nums[i+1])
        for j := 0; j < length; j++ {
            seeds = append(seeds, start+j)
        }
    }
    return seeds
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error opening file:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Read seeds from the file
    scanner.Scan()
    seedRanges := parseSeedRanges(strings.TrimPrefix(scanner.Text(), "seeds: "))

    // Read maps
    maps := make([]RangeMap, 7)
    for i := range maps {
        scanner.Scan()
        maps[i] = parseMap(strings.SplitN(scanner.Text(), ": ", 2)[1])
    }

    lowestLocation := -1

    for _, seed := range seedRanges {
        soil := mapValue(maps[0], seed)
        fertilizer := mapValue(maps[1], soil)
        water := mapValue(maps[2], fertilizer)
        light := mapValue(maps[3], water)
        temperature := mapValue(maps[4], light)
        humidity := mapValue(maps[5], temperature)
        location := mapValue(maps[6], humidity)

        if lowestLocation == -1 || location < lowestLocation {
            lowestLocation = location
        }
    }

    fmt.Println(lowestLocation)
}
