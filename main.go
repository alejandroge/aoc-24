package main

import (
	"flag"
	"fmt"

	"aoc-24/day10"
	"aoc-24/utils"
)

func main() {
	day := flag.Int("day", 1, "Specify the day you want to run")
	part := flag.Int("part", 1, "Specify the part you want to run")
	test := flag.Bool("test", false, "Run with the test input for the current day")
	flag.Parse()

	type dayPartFunc func(test bool) int

	var functions = map[int]map[int]dayPartFunc{
		1: {
			1: day1Part1,
			2: day1Part2,
		},
		2: {
			1: day2Part1,
			2: day2Part2,
		},
		3: {
			1: day3Part1,
			2: day3Part2,
		},
		4: {
			1: day4Part1,
			2: day4Part2,
		},
		5: {
			1: day5Part1,
		},
		6: {
			1: day6Part1,
		},
		10: {
			1: day10.Part1,
		},
	}

	if parts, ok := functions[*day]; ok {
		if fn, ok := parts[*part]; ok {
			result := fn(*test)

			if *test {
				fmt.Printf("%sRunning in test mode. Do not submit this result%s\n", utils.ColorYellow, utils.ColorReset)
			}
			fmt.Println("Result:", result)
		} else {
			fmt.Println("Invalid part: probably I didn't do it 😅")
		}
	} else {
		fmt.Println("Invalid day: probably I didn't do it 😅")
	}
}
