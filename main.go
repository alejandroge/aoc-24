package main

import (
	"fmt"
	"flag"
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
	default:
		fmt.Println("Unknown day: probably I didn't do it 😅")
	}
}
