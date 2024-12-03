package main

import (
	"fmt"
	"regexp"
)

func day3Part1() int {
	lines := readFileLines("./inputs/day3.input.txt")

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("whoopsie")
		return 0
	}

	numberOfMatches := 0
	sum := 0
	for _, line := range(lines) {
		matches := r.FindAllStringSubmatch(line, -1)

		numberOfMatches += len(matches)
		for i := 0; i < len(matches); i++ {
			match := matches[i]
			matchResult := stringToInt(match[1]) * stringToInt(match[2])
			sum += matchResult
		}
	}

	fmt.Println("total matches:", numberOfMatches)
	fmt.Println("total sum:", sum)
	return sum
}

func day3Part2() {
}