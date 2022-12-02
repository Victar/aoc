package util

import (
	"bufio"
	"os"
)

var baseDir = "/Users/vkad2506/AdventOfCode/java/src/test/resources/"

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(baseDir + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
