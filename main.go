package main

import (
	"flag"
	"fmt"
)

func main() {
	// defined the flags that can be used
	day := flag.Int("d", 1, "Specify the day you want to run")
	part := flag.Int("p", 1, "Specify the part you want to run")
	test := flag.Bool("t", false, "Run with the test input for the current day")

	// parse the flags
	flag.Parse()

	// Define the function signature for your day/part functions
	type dayPartFunc func(test bool) int

	// Create a map to store day/part functions
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
		// Add more days as needed
	}

	// Lookup and execute the corresponding function
	if parts, ok := functions[*day]; ok {
		if fn, ok := parts[*part]; ok {
			fn(*test) // Pass the test flag down
		} else {
			fmt.Println("Invalid part: probably I didn't do it 😅")
		}
	} else {
		fmt.Println("Invalid day: probably I didn't do it 😅")
	}
}
