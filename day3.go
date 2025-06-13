package main

import (
	"fmt"
	"regexp"

	"aoc-24/utils"
)

func day3Part1(test bool) int {
	lines := utils.ReadFileLines("./inputs/day3.input.txt")

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("whoopsie")
		return 0
	}

	numberOfMatches := 0
	sum := 0
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)

		numberOfMatches += len(matches)
		for i := 0; i < len(matches); i++ {
			match := matches[i]
			matchResult := utils.StringToInt(match[1]) * utils.StringToInt(match[2])
			sum += matchResult
		}
	}

	fmt.Println("total matches:", numberOfMatches)
	fmt.Println("total sum:", sum)
	return sum
}

func day3Part2(test bool) int {
	lines := utils.ReadFileLines("./inputs/day3.input.txt")

	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don\'t\(\)`
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("whoopsie")
		return 0
	}

	numberOfMatches := 0
	enabled := true
	sum := 0

	enable_command := "do()"
	disable_command := "don't()"

	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)

		numberOfMatches += len(matches)
		for i := 0; i < len(matches); i++ {
			match := matches[i]

			switch match[0] {
			case enable_command:
				enabled = true
			case disable_command:
				enabled = false
			default:
				if enabled {
					matchResult := utils.StringToInt(match[1]) * utils.StringToInt(match[2])
					sum += matchResult
				}
			}
		}
	}

	fmt.Println("total matches:", numberOfMatches)
	fmt.Println("total sum:", sum)
	return sum
}
