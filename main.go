package main

import (
	"flag"
	"fmt"
)

func main() {
	// defined the flags that can be used
	day := flag.Int("day", 1, "Specify the day you want to run")
	part := flag.Int("part", 1, "Specify the part you want to run")

	// parse the flags
	flag.Parse()

	switch *day {
	case 1:
		if *part == 1 {
			day1Part1()
		} else {
			day1Part2()
		}
	case 2:
		if *part == 1 {
			day2Part1()
		} else {
			day2Part2()
		}
	case 3:
		if *part == 1 {
			day3Part1()
		} else {
			day3Part2()
		}
	case 4:
		if *part == 1 {
			day4Part1()
		} else {
			day4Part2()
		}
	case 5:
		if *part == 1 {
			day5Part1()
		} else {
			day5Part2()
		}
	case 6:
		if *part == 1 {
			day6Part1()
		} else {
			day6Part2()
		}
	default:
		fmt.Println("Unknown day: probably I didn't do it 😅")
	}
}
