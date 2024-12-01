package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFileLines() []string {
	var lines []string

	// Open the file
	file, err := os.Open("day1.input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return lines
	}
	defer file.Close() // Ensure the file is closed when the function ends

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return lines
}

func stringToInt(s string) int {
	number, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Error converting the string to a number")
		return -1
	}

	return number
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func columns(lines []string) ([]int, []int) {
	var leftColumn	[]int
	var rightColumn	[]int

	for _, line := range(lines) {
		columns := strings.Fields(line)

		leftColumn = append(leftColumn, stringToInt(columns[0]))
		rightColumn = append(rightColumn, stringToInt(columns[1]))
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	return leftColumn, rightColumn
}

func countOccurrences(n int, col []int) int {
	repetitions := 0

	for _, locationId := range(col) {
		if n < locationId {
			return repetitions
		}

		if n == locationId {
			repetitions++
		}
	}

	return repetitions
}

func day1Part1() int {
	lines := readFileLines()
	leftColumn, rightColumn := columns(lines)

	totalDistance := 0
	for i := 0; i < len(leftColumn); i++ {
		lineDistance := absolute(leftColumn[i] - rightColumn[i])

		totalDistance += lineDistance
	}

	fmt.Println(totalDistance)
	return totalDistance
}

func day1Part2() int {
	similarityScore := 0

	lines := readFileLines()
	leftColumn, rightColumn := columns(lines)

	for _, locationId := range(leftColumn) {
		repetitions := countOccurrences(locationId, rightColumn)

		if repetitions > 0 {
			similarityScore += (locationId * repetitions)
		}
	}

	fmt.Println(similarityScore)
	return similarityScore
}
