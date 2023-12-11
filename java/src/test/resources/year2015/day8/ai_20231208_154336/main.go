
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	var codeChars, memoryChars int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		codeChars += len(line)

		unquoted, err := strconv.Unquote(line) // Using strconv.Unquote to handle escape sequences
		if err != nil {
			fmt.Println("Error unquoting string:", err, line)
			return
		}
		memoryChars += len(unquoted)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println(codeChars - memoryChars)
}
