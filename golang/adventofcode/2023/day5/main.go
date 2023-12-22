package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var DAY = "5"

type Converter struct {
	from  string
	to    string
	rules []Rule
}

type Range struct {
	start int
	end   int
}
type Rule struct {
	source      int
	destination int
	size        int
	start       int
	end         int
}

func (r *Rule) printRule() {
	fmt.Println("   ", r.source, r.source+r.size, r)
}
func (c *Converter) convert(sourceNumber int) int {
	for _, rule := range c.rules {
		if rule.source <= sourceNumber && rule.source+rule.size > sourceNumber {
			return rule.destination + sourceNumber - rule.source
		}
	}
	return sourceNumber
}

func (c *Converter) convertRange(ranges []Range) []Range {
	newRanges := []Range{}
	for _, r := range ranges {
		fmt.Println(r)
		for _, rule := range c.rules {
			// .............. [79 ... 93)... // range.
			//.....[55...68).................// range
			//...................[98, 100)...// rule
			// ..[50..............98)........// rule
			// Check if there is an overlap between the range and the rule
			if r.start < rule.end && r.end > rule.start {
				// Split the range based on the overlap area
				if r.start < rule.start {
					newRanges = append(newRanges, Range{r.start, rule.start})
				}
				if r.end > rule.end {
					newRanges = append(newRanges, Range{rule.end, r.end})
				}
			} else {
				// If there is no overlap, add the original range to the result
				newRanges = append(newRanges, r)
			}
		}
	}
	fmt.Println(newRanges)
	return newRanges
}

func logTime(methodName string) func() {
	startTime := time.Now()
	return func() {
		elapsed := time.Since(startTime)
		fmt.Printf("%s took %s\n", methodName, elapsed)
	}
}

func main() {
	runSilver()
	runGoldBF()
}

func runSilver() {
	defer logTime("runSilver")()

	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	seeds := []int{}
	seedsstr := strings.Split(lines[0], " ")
	for i := 1; i < len(seedsstr); i++ {
		seed, _ := strconv.Atoi(seedsstr[i])
		seeds = append(seeds, seed)
	}

	converterMap := make(map[string]Converter)
	cur := Converter{}
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			if cur.from != "" {
				converterMap[cur.from] = cur
				cur = Converter{}
			}
		} else if strings.Contains(line, "-to-") {
			strs := strings.Split(line, " ")
			name := strings.Split(strs[0], "-to-")
			cur.from = name[0]
			cur.to = name[1]
		} else {
			strs := strings.Split(line, " ")
			source, _ := strconv.Atoi(strs[1])
			dest, _ := strconv.Atoi(strs[0])
			size, _ := strconv.Atoi(strs[2])
			cur.rules = append(cur.rules, Rule{source, dest, size, source, source + size})
		}
	}
	if cur.from != "" {
		converterMap[cur.from] = cur
		cur = Converter{}
	}
	minSeed := 100000000000
	for _, seed := range seeds {
		cur := converterMap["seed"]
		for cur.to != "location" {
			seed = cur.convert(seed)
			cur = converterMap[cur.to]
		}
		seed = cur.convert(seed)
		if seed < minSeed {
			minSeed = seed
		}
	}
	fmt.Println(minSeed)
}

func runGoldBF() {
	defer logTime("runGoldBF")()
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	ranges := []Range{}
	seedsstr := strings.Split(lines[0], " ")
	for i := 1; i < len(seedsstr); i = i + 2 {
		start, _ := strconv.Atoi(seedsstr[i])
		size, _ := strconv.Atoi(seedsstr[i+1])
		ranges = append(ranges, Range{start, start + size})
	}
	fmt.Println(ranges)

	converterMap := make(map[string]Converter)
	cur := Converter{}
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			if cur.from != "" {
				converterMap[cur.from] = cur
				cur = Converter{}
			}
		} else if strings.Contains(line, "-to-") {
			strs := strings.Split(line, " ")
			name := strings.Split(strs[0], "-to-")
			cur.from = name[0]
			cur.to = name[1]
		} else {
			strs := strings.Split(line, " ")
			source, _ := strconv.Atoi(strs[1])
			dest, _ := strconv.Atoi(strs[0])
			size, _ := strconv.Atoi(strs[2])
			cur.rules = append(cur.rules, Rule{source, dest, size, source, source + size})
		}
	}
	if cur.from != "" {
		converterMap[cur.from] = cur
		cur = Converter{}
	}
	minSeed := 100000000000
	for _, r := range ranges {
		fmt.Println(r, r.end-r.start)
		for cur := r.start; cur < r.end; cur++ {
			seed := cur
			cur := converterMap["seed"]
			for cur.to != "location" {
				seed = cur.convert(seed)
				cur = converterMap[cur.to]
			}
			seed = cur.convert(seed)
			if seed < minSeed {
				fmt.Println("new min", seed, minSeed)
				minSeed = seed
			}
		}
	}
	fmt.Println(minSeed)
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/sample.txt")
	if err != nil {
		panic(err)
	}
	ranges := []Range{}
	seedsstr := strings.Split(lines[0], " ")
	for i := 1; i < len(seedsstr); i = i + 2 {
		start, _ := strconv.Atoi(seedsstr[i])
		size, _ := strconv.Atoi(seedsstr[i+1])
		ranges = append(ranges, Range{start, start + size})
	}
	fmt.Println(ranges)

	converterMap := make(map[string]Converter)
	cur := Converter{}
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			if cur.from != "" {
				converterMap[cur.from] = cur
				cur = Converter{}
			}
		} else if strings.Contains(line, "-to-") {
			strs := strings.Split(line, " ")
			name := strings.Split(strs[0], "-to-")
			cur.from = name[0]
			cur.to = name[1]
		} else {
			strs := strings.Split(line, " ")
			source, _ := strconv.Atoi(strs[1])
			dest, _ := strconv.Atoi(strs[0])
			size, _ := strconv.Atoi(strs[2])
			cur.rules = append(cur.rules, Rule{source, dest, size, source, source + size})
		}
	}
	if cur.from != "" {
		converterMap[cur.from] = cur
		cur = Converter{}
	}
	cur = converterMap["seed"]
	//for cur.to != "location" {
	//	ranges = cur.convertRange(ranges)
	//	cur = converterMap[cur.to]
	//}
	ranges = cur.convertRange(ranges)
	fmt.Println(ranges)
}
