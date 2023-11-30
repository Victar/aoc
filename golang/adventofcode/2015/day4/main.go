package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	runAny(5)
	runAny(6)

}

func runAny(length int) {
	secretKey := "ckczppom"
	number := 1
	pattern := strings.Repeat("0", length)
	for {
		data := []byte(secretKey + strconv.Itoa(number))
		hash := md5.Sum(data)
		hashStr := hex.EncodeToString(hash[:])

		if hashStr[:length] == pattern {
			fmt.Printf("Lowest number: %d\n", number)
			break
		}

		number++
	}
}
