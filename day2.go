package main

import (
	"fmt"
	"strings"
)

type Row []int

func getIntRows(lines []string) []Row {
	var rows []Row

	for _, line := range(lines) {
		levels := strings.Fields(line)

		var row Row
		for _, level := range(levels) {
			row = append(row, stringToInt(level))
		}

		rows = append(rows, row)
	}

	return rows
}

func isDirectionValid(ascending bool, a int, b int) bool {
	if ascending {
		return a < b
	} else {
		return b < a
	}
}

func isIncreaseValid(a int, b int) bool {
	difference := absolute(a - b)
	if difference < 1 || difference > 3 {
		return false
	}
	return true
}

func isAscending(row Row) bool {
	positiveDeltaCount := 0
	negativeDeltaCount := 0

	for i := 1; i < len(row); i++ {
		upperValue := row[i]
		lowerValue := row[i - 1]

		delta := upperValue - lowerValue

		if delta > 0 {
			positiveDeltaCount++
		} else if delta < 0 {
			negativeDeltaCount++
		}
	}

	return positiveDeltaCount > negativeDeltaCount
}

func increasedLowerIndex(lowerIndex int, upperIndex int, unsafeLevels []int) int {
	potentialNewIndex := lowerIndex + 1

	for potentialNewIndex < upperIndex {
		valid := true

		for _, unsafeIndex := range(unsafeLevels) {
			if potentialNewIndex == unsafeIndex {
				valid = false
			}
		}

		if valid {
			return potentialNewIndex
		}

		potentialNewIndex++
	}

	return potentialNewIndex
}

func areLevelsValid(lowerValue int, upperValue int, ascending bool) bool {
	directionValid := isDirectionValid(ascending, lowerValue, upperValue)
	increaseValid := isIncreaseValid(lowerValue, upperValue)
	return directionValid && increaseValid
}

func isRowValid(row Row, unsafetyThreshold int) bool {
	ascending := isAscending(row)

	var unsafeLevels []int

	lowerIndex := 0
	upperIndex := 1

	for (len(unsafeLevels) <= unsafetyThreshold) && (upperIndex < len(row)) {
		valid := true

		lowerValue := row[lowerIndex]
		upperValue := row[upperIndex]

		valid = areLevelsValid(lowerValue, upperValue, ascending)

		if valid {
			lowerIndex = increasedLowerIndex(lowerIndex, upperIndex, unsafeLevels)
		} else {
			if (upperIndex + 1) < len(row) {
				// which one is unsafe? could be that the first one is unsafe ...
				lowerValue = row[lowerIndex]
				upperValue = row[upperIndex + 1]
				valid = areLevelsValid(lowerValue, upperValue, ascending)

				if valid {
					unsafeLevels = append(unsafeLevels, upperIndex)
				} else {
					if lowerIndex == 0 {
						unsafeLevels = append(unsafeLevels, lowerIndex)
					} else {
						unsafeLevels = append(unsafeLevels, upperIndex)
						unsafeLevels = append(unsafeLevels, upperIndex + 1)
					}
				}
			} else {
				unsafeLevels = append(unsafeLevels, upperIndex)
			}

			lowerIndex = increasedLowerIndex(lowerIndex, upperIndex, unsafeLevels)
		}
		upperIndex = lowerIndex + 1
	}

	return len(unsafeLevels) <= unsafetyThreshold
}

func day2Part1() int {
	lines := readFileLines("./day2.input.txt")
	rows := getIntRows(lines)

	validRowsCounter := 0
	for i := 0; i < len(rows); i++ {
		if isRowValid(rows[i], 0) {
			validRowsCounter++
		}
	}

	fmt.Println(validRowsCounter)
	return validRowsCounter
}

func day2Part2() int {
	// 668 too low
	// 671 too low
	// 679 too high
	lines := readFileLines("./day2.input.txt")
	rows := getIntRows(lines)

	validRowsCounter := 0
	for i := 0; i < 5; i++ {
		fmt.Print(rows[i])
		if isRowValid(rows[i], 1) {
			fmt.Print(" valid!!")
			validRowsCounter++
		}
		fmt.Println("")
	}

	fmt.Println(validRowsCounter)
	return validRowsCounter
}

