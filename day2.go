package main

import (
	"fmt"
	"strings"
)

type Row []int

func getIntRows(lines []string) []Row {
	var rows []Row

	for _, line := range lines {
		levels := strings.Fields(line)

		var row Row
		for _, level := range levels {
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
		lowerValue := row[i-1]

		delta := upperValue - lowerValue

		if delta > 0 {
			positiveDeltaCount++
		} else if delta < 0 {
			negativeDeltaCount++
		}
	}
	return positiveDeltaCount > negativeDeltaCount
}

func increasedLowerIndex(lowerIndex int, arrayLen int, unsafeLevels []int) int {
	potentialNewIndex := lowerIndex + 1

	for potentialNewIndex < arrayLen {
		valid := true

		for _, unsafeIndex := range unsafeLevels {
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

func moveIndexesWindow(indexesWindow []int, invalidIndexes []int) []int {
	lowestIndex := indexesWindow[0]
	newLowIndex := increasedLowerIndex(lowestIndex, 100, invalidIndexes)
	newMidIndex := increasedLowerIndex(newLowIndex, 100, invalidIndexes)

	return []int{newLowIndex, newMidIndex, newMidIndex + 1}
}

func isRowValidWindow(row Row) bool {
	ascending := isAscending(row)
	indexesWindow := []int{0, 1, 2}
	var unsafeLevels []int

	for (indexesWindow[1] < len(row)) && (len(unsafeLevels) <= 1) {
		lowValue := row[indexesWindow[0]]
		midValue := row[indexesWindow[1]]

		firstPairValid := areLevelsValid(lowValue, midValue, ascending)

		endOfSlice := indexesWindow[1]+1 == len(row)
		if endOfSlice {
			if !firstPairValid {
				unsafeLevels = append(unsafeLevels, indexesWindow[1])
			}
		} else {
			topValue := row[indexesWindow[2]]
			secondPairValid := areLevelsValid(midValue, topValue, ascending)

			bothValid := firstPairValid && secondPairValid
			if !bothValid {
				if secondPairValid {
					unsafeLevels = append(unsafeLevels, indexesWindow[0])
				} else if firstPairValid {
					outsidePairValid := areLevelsValid(lowValue, topValue, ascending)

					if outsidePairValid {
						unsafeLevels = append(unsafeLevels, indexesWindow[1])
					} else {
						unsafeLevels = append(unsafeLevels, indexesWindow[2])
					}
				} else {
					outsidePairValid := areLevelsValid(lowValue, topValue, ascending)
					if outsidePairValid {
						unsafeLevels = append(unsafeLevels, indexesWindow[1])
					} else {
						unsafeLevels = append(unsafeLevels, indexesWindow[0])
						unsafeLevels = append(unsafeLevels, indexesWindow[1])
					}
				}
			}
		}

		indexesWindow = moveIndexesWindow(indexesWindow, unsafeLevels)
	}

	return len(unsafeLevels) <= 1
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
				upperValue = row[upperIndex+1]

				valid = areLevelsValid(lowerValue, upperValue, ascending)

				if valid {
					unsafeLevels = append(unsafeLevels, upperIndex)
				} else {
					if lowerIndex == 0 {
						unsafeLevels = append(unsafeLevels, lowerIndex)
					} else {
						unsafeLevels = append(unsafeLevels, upperIndex)
						unsafeLevels = append(unsafeLevels, upperIndex+1)
					}
				}
			} else {
				unsafeLevels = append(unsafeLevels, upperIndex)
			}

			lowerIndex = increasedLowerIndex(lowerIndex, len(row), unsafeLevels)
		}
		upperIndex = lowerIndex + 1
	}

	return len(unsafeLevels) <= unsafetyThreshold
}

func day2Part1() int {
	lines := readFileLines("./inputs/day2.test.txt")
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
	lines := readFileLines("./inputs/day2.input.txt")
	rows := getIntRows(lines)

	validRowsCounter := 0
	for i := 0; i < len(rows); i++ {
		fmt.Print(i+1, "	")

		valid := isRowValidWindow(rows[i])

		if valid {
			fmt.Print("✅ ")
			validRowsCounter++
		} else {
			fmt.Print("❌ ")
		}
		fmt.Println(rows[i])
	}

	fmt.Println(validRowsCounter)
	return validRowsCounter
}
