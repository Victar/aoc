
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func decodeString(str string) string {
	var decoded strings.Builder
	inEscape := false
	inHex := false
	hexChars := 0

	for _, r := range str {
		if inEscape {
			switch r {
			case '\\', '"':
				decoded.WriteRune(r)
				inEscape = false
			case 'x':
				inHex = true
				hexChars = 0
				inEscape = false
				continue
			}
		} else if inHex {
			hexChars++
			if hexChars == 2 {
				inHex = false
			}
			continue
		} else if r == '\\' {
			inEscape = true
			continue
		} else {
			decoded.WriteRune(r)
		}
	}

	return decoded.String()
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	codeChars := 0
	memoryChars := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = line[1 : len(line)-1] // Strip the enclosing quotes

		codeChars += utf8.RuneCountInString(line) + 2 // +2 for the stripped quotes
		decodedString := decodeString(line)
		memoryChars += utf8.RuneCountInString(decodedString)

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println(codeChars - memoryChars)
}
