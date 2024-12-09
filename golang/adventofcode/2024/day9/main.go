package main

import (
	"adventofcode/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var DAY = "9"

func main() {
	runBoth()
}

type File struct {
	id   int
	size int
}

func (f File) String() string {
	str := fmt.Sprintf("%d", f.id)
	if f.id == -1 {
		str = "."
	}
	return strings.Repeat(str, f.size)
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}

	input := lines[0]
	if len(input)%2 == 1 {
		input = input + "0"
	}

	var filesSilver []File
	var filesGold []File

	for i := 0; i < len(input); i = i + 2 {
		size, _ := strconv.Atoi(string(input[i]))
		free, _ := strconv.Atoi(string(input[i+1]))
		id := i / 2
		filesSilver = append(filesSilver, slices.Repeat([]File{{id, 1}}, size)...)
		filesSilver = append(filesSilver, slices.Repeat([]File{{-1, 1}}, free)...)
		filesGold = append(filesGold, File{id, size}, File{-1, free})
	}
	fmt.Println(solve(filesSilver))
	fmt.Println(solve(filesGold))
}

func solve(files []File) int {
	for fileIdx := len(files) - 1; fileIdx >= 0; fileIdx-- {
		for freeIdx := 0; freeIdx < fileIdx; freeIdx++ {
			if files[fileIdx].id != -1 && files[freeIdx].id == -1 && files[fileIdx].size <= files[freeIdx].size {
				diff := files[freeIdx].size - files[fileIdx].size
				//fmt.Println(files, diff, len(files))
				files[freeIdx], files[fileIdx] = files[fileIdx], files[freeIdx]
				if diff > 0 {
					files = slices.Insert(files, freeIdx+1, File{-1, diff})
					fileIdx++
					minAdjust := files[fileIdx].size - diff
					if minAdjust >= 0 {
						files[fileIdx].size = minAdjust
					}
				}
				//fmt.Println(files, diff, len(files))
			}
		}
	}
	ans := 0
	idx := 0
	for _, file := range files {
		if file.id != -1 {
			for s := 0; s < file.size; s++ {
				ans += file.id * (idx + s)
			}
		}
		idx = idx + file.size
	}
	return ans
}
