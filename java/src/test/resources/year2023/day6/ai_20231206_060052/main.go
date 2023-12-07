package main

import (
	"bufio"
	"os"
)

//
//func main() {
//	input, err := readInput("input.txt")
//	if err != nil {
//		log.Fatalf("Failed to read input: %v\n", err)
//	}
//
//	totalWays := 1
//
//
//	for _, line := range input {
//		parts := strings.Split(line, " ")
//		time, err := strconv.Atoi(parts[1])
//		if err != nil {
//			log.Fatalf("Failed to parse time: %s", parts[1])
//		}
//		distanceRecord, err := strconv.Atoi(parts[3])
//		if err != nil {
//			log.Fatalf("Failed to parse distance record: %s", parts[3])
//		}
//
//		ways := calculateWaysToWin(time, distanceRecord)
//		totalWays *= ways
//	}
//
//	fmt.Println(totalWays)
//}

func readInput(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func calculateWaysToWin(time int, distanceRecord int) int {
	ways := 0
	for buttonMillis := 0; buttonMillis < time; buttonMillis++ {
		distance := buttonMillis * (time - buttonMillis)
		if distance > distanceRecord {
			ways++
		}
	}
	return ways
}
