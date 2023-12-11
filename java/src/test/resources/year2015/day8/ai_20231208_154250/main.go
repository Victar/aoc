
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	codeCharacters := 0
	memoryCharacters := 0

	for scanner.Scan() {
		line := scanner.Text()
		codeCharacters += len(line)

		// Unquote the string to calculate the in-memory characters.
		// The strconv.Unquote function handles the escape sequences
		// correctly according to the Go language specification.
		unquoted, err := strconv.Unquote(line)
		if err != nil {
			log.Fatal(err)
		}
		memoryCharacters += len(unquoted)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(codeCharacters - memoryCharacters)
}
