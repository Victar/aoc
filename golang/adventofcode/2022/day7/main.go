package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

var DAY = "7"

func main() {
	runSilver()
	runGold()
}

type dir struct {
	name        string
	parent      *dir
	children    []*dir
	childrenMap map[string]*dir
	files       map[string]int
}

func readRoot() *dir {
	lines, err := util.ReadFile("year2022/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	root := newDir("/", nil)
	currentDir := root
	currentLS := false
	for i := 1; i < len(lines); i++ {
		input := lines[i]
		if currentLS && !strings.HasPrefix(input, "$") {
			currentDir.addFileOrFolder(input)
		}
		if strings.HasPrefix(input, "$ ls") {
			currentLS = true
		}
		if strings.HasPrefix(input, "$ cd") {
			currentLS = false
			currentDir = currentDir.getDirByName(input)
		}
	}
	return root
}
func newDir(name string, parent *dir) *dir {
	return &dir{name: name,
		parent:      parent,
		children:    make([]*dir, 0),
		childrenMap: make(map[string]*dir),
		files:       make(map[string]int)}
}

func (current dir) addFileOrFolder(input string) {
	parts := strings.Fields(input)
	if "dir" == parts[0] {
		subDir := newDir(parts[1], &current)
		current.children = append(current.children, subDir)
		current.childrenMap[parts[1]] = subDir
	} else {
		current.files[parts[1]], _ = strconv.Atoi(parts[0])
	}
}

func (current dir) countSilver(limit int) int {
	answer := 0
	if current.getSize() < limit {
		answer += current.getSize()
	}
	for _, subdir := range current.childrenMap {
		answer += subdir.countSilver(limit)
	}
	return answer
}

func (current dir) countGold(limit int) int {
	answer := 0
	if current.getSize() >= limit {
		answer = current.getSize()
	}
	for _, subdir := range current.childrenMap {
		childAnswer := subdir.countGold(limit)
		if childAnswer > 0 && childAnswer < answer {
			answer = childAnswer
		}
	}
	return answer
}

func (current dir) getDirByName(input string) *dir {
	dirName := strings.Fields(input)[2]
	if ".." == dirName {
		return current.parent
	} else {
		for _, subdir := range current.childrenMap {
			if dirName == subdir.name {
				return subdir
			}
		}
	}
	return nil
}

func (current dir) getSize() int {
	size := 0
	for _, s := range current.files {
		size += s
	}
	for _, subdir := range current.childrenMap {
		size += subdir.getSize()
	}
	return size
}

func runSilver() {
	root := readRoot()
	println(root.countSilver(100000))
}

func runGold() {
	root := readRoot()
	total := 70000000
	needSpace := 30000000
	toDelete := needSpace - (total - root.getSize())
	println(root.countGold(toDelete))
}
