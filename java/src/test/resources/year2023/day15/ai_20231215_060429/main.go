
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// hashAlgorithm computes the HASH algorithm as described in the task
func hashAlgorithm(str string) int {
	currentValue := 0
	for i := 0; i < len(str); i++ {
		asciiValue := int(str[i])
		currentValue += asciiValue
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

// readInput reads the content from the input.txt file
func readInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func main() {
	const inputFileName = "input.txt"
	inputString := readInput(inputFileName)

	// Removing any newline characters that might be present in the input
	inputString = strings.ReplaceAll(inputString, "\n", "")

	// Splitting the input string by commas to handle each step in the sequence
	steps := strings.Split(inputString, ",")

	sum := 0
	for _, step := range steps {
		result := hashAlgorithm(step)
		sum += result
	}

	fmt.Println(sum)
}
