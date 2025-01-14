package main

import (
	"fmt"
	"sort"
	"strings"

	"aoc-24/utils"
)

func columns(lines []string) ([]int, []int) {
	var leftColumn []int
	var rightColumn []int

	for _, line := range lines {
		columns := strings.Fields(line)

		leftColumn = append(leftColumn, utils.StringToInt(columns[0]))
		rightColumn = append(rightColumn, utils.StringToInt(columns[1]))
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	return leftColumn, rightColumn
}

func countOccurrences(n int, col []int) int {
	repetitions := 0

	for _, locationId := range col {
		if n < locationId {
			return repetitions
		}

		if n == locationId {
			repetitions++
		}
	}

	return repetitions
}

func day1Part1(test bool) int {
	lines := utils.ReadFileLines("./inputs/day1.input.txt")
	leftColumn, rightColumn := columns(lines)

	totalDistance := 0
	for i := 0; i < len(leftColumn); i++ {
		lineDistance := utils.Absolute(leftColumn[i] - rightColumn[i])

		totalDistance += lineDistance
	}

	fmt.Println(totalDistance)
	return totalDistance
}

func day1Part2(test bool) int {
	similarityScore := 0

	lines := utils.ReadFileLines("./inputs/day1.input.txt")
	leftColumn, rightColumn := columns(lines)

	for _, locationId := range leftColumn {
		repetitions := countOccurrences(locationId, rightColumn)

		if repetitions > 0 {
			similarityScore += (locationId * repetitions)
		}
	}

	fmt.Println(similarityScore)
	return similarityScore
}
