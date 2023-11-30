package main

import (
	"adventofcode/util"
	"fmt"
)

func main() {
	runSilver()
	runGold()
}
func runSilver() {
	ar, err := util.ReadFile("year2015/day2/input.txt")
	if err != nil {
		panic(err)
	}
	var sum = 0
	for _, str := range ar {
		var l, w, h int
		_, err := fmt.Sscanf(str, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			panic(err)
		}
		var s1 = l * w
		var s2 = w * h
		var s3 = h * l
		var m = s1
		if s2 < m {
			m = s2
		}
		if s3 < m {
			m = s3
		}
		var res = 2*s1 + 2*s2 + 2*s3 + m
		sum += res
	}
	println(sum)
}

func runGold() {
	ar, err := util.ReadFile("year2015/day2/input.txt")
	if err != nil {
		panic(err)
	}
	var sum = 0
	for _, str := range ar {
		var l, w, h int
		_, err := fmt.Sscanf(str, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			panic(err)
		}
		var m = l
		if m < w {
			m = w
		}
		if m < h {
			m = h
		}
		var res = l*w*h + 2*(l+w+h-m)
		sum += res
	}
	println(sum)
}
