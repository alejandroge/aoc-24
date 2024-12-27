package main

import (
	"fmt"

	utils "aoc-24/utils"
)

type Cell struct {
	value  rune
	inWord bool
}

func isXMasValid(crosswordGamePtr *[][]Cell, i int, j int) bool {
	crosswordGame := *crosswordGamePtr

	topLeftCell := crosswordGame[i-1][j-1]
	bottomLeftCell := crosswordGame[i+1][j-1]
	topRightCell := crosswordGame[i-1][j+1]
	bottomRightCell := crosswordGame[i+1][j+1]

	outsideCells := []Cell{topLeftCell, bottomLeftCell, topRightCell, bottomRightCell}
	mCount := 0
	sCount := 0

	for _, cell := range outsideCells {
		if cell.value == 'M' {
			mCount++
		} else if cell.value == 'S' {
			sCount++
		}
	}

	if mCount != 2 || sCount != 2 {
		return false
	}

	mainDiagValid := (topLeftCell.value == 'M' && bottomRightCell.value == 'S') ||
		(bottomRightCell.value == 'M' && topLeftCell.value == 'S')
	revDiagValid := (topRightCell.value == 'M' && bottomLeftCell.value == 'S') ||
		(bottomLeftCell.value == 'M' && topRightCell.value == 'S')

	return mainDiagValid && revDiagValid
}

func markCrossAsValid(crosswordGamePtr *[][]Cell, i int, j int) {
	crosswordGame := *crosswordGamePtr

	crosswordGame[i][j].inWord = true
	crosswordGame[i-1][j-1].inWord = true
	crosswordGame[i+1][j-1].inWord = true
	crosswordGame[i-1][j+1].inWord = true
	crosswordGame[i+1][j+1].inWord = true
}

func scanHorizontalXMas(crosswordGamePtr *[][]Cell) int {
	crosswordGame := *crosswordGamePtr
	rowLen := len(crosswordGame)
	colLen := len(crosswordGame[0])

	count := 0

	for i := 1; i < (rowLen - 1); i++ {
		for j := 1; j < (colLen - 1); j++ {
			cell := crosswordGame[i][j]

			if cell.value == 'A' {
				if isXMasValid(crosswordGamePtr, i, j) {
					markCrossAsValid(crosswordGamePtr, i, j)
					count++
				}
			}
		}
	}

	return count
}

func scanHorizontal(crosswordGamePtr *[][]Cell, matchingString string) int {
	crosswordGame := *crosswordGamePtr
	rowLen := len(crosswordGame)
	colLen := len(crosswordGame[0])

	runesToMatch := []rune(matchingString)

	count := 0

	for i := 0; i < rowLen; i++ {
		j := 0
		for j < (colLen - 3) {
			firstCell := crosswordGame[i][j]

			if firstCell.value != runesToMatch[0] {
				j++
			} else {
				secondCell := crosswordGame[i][j+1]

				if secondCell.value != runesToMatch[1] {
					j++
				} else {
					thirdCell := crosswordGame[i][j+2]

					if thirdCell.value != runesToMatch[2] {
						j = j + 2
					} else {
						fourthCell := crosswordGame[i][j+3]

						if fourthCell.value == runesToMatch[3] {
							count++
							for offset := 0; offset < 4; offset++ {
								crosswordGame[i][j+offset].inWord = true
							}
						}
						j = j + 3
					}
				}
			}
		}
	}

	return count
}

func scanVertical(crosswordGamePtr *[][]Cell, matchingString string) int {
	crosswordGame := *crosswordGamePtr
	rowLen := len(crosswordGame)
	colLen := len(crosswordGame[0])

	runesToMatch := []rune(matchingString)

	count := 0

	for j := 0; j < colLen; j++ {
		i := 0
		for i < (rowLen - 3) {
			firstCell := crosswordGame[i][j]

			if firstCell.value != runesToMatch[0] {
				i++
			} else {
				secondCell := crosswordGame[i+1][j]

				if secondCell.value != runesToMatch[1] {
					i++
				} else {
					thirdCell := crosswordGame[i+2][j]

					if thirdCell.value != runesToMatch[2] {
						i = i + 2
					} else {
						fourthCell := crosswordGame[i+3][j]

						if fourthCell.value == runesToMatch[3] {
							count++
							for offset := 0; offset < 4; offset++ {
								crosswordGame[i+offset][j].inWord = true
							}
						}
						i = i + 3
					}
				}
			}
		}
	}

	return count
}

func scanDiagonalTopToBottom(crosswordGamePtr *[][]Cell, matchingString string) int {
	crosswordGame := *crosswordGamePtr
	rowLen := len(crosswordGame)
	colLen := len(crosswordGame[0])

	runesToMatch := []rune(matchingString)

	count := 0

	for i := 0; i < (rowLen - 3); i++ {
		j := 0
		for j < (colLen - 3) {
			firstCell := crosswordGame[i][j]

			if firstCell.value != runesToMatch[0] {
				j++
			} else {
				secondCell := crosswordGame[i+1][j+1]

				if secondCell.value != runesToMatch[1] {
					j++
				} else {
					thirdCell := crosswordGame[i+2][j+2]

					if thirdCell.value != runesToMatch[2] {
						j++
					} else {
						fourthCell := crosswordGame[i+3][j+3]

						if fourthCell.value == runesToMatch[3] {
							count++
							for offset := 0; offset < 4; offset++ {
								crosswordGame[i+offset][j+offset].inWord = true
							}
						}

						j++
					}
				}
			}
		}
	}

	return count
}

func scanDiagonalBottomToTop(crosswordGamePtr *[][]Cell, matchingString string) int {
	crosswordGame := *crosswordGamePtr
	rowLen := len(crosswordGame)
	colLen := len(crosswordGame[0])

	runesToMatch := []rune(matchingString)

	count := 0

	for i := (rowLen - 1); i > 2; i-- {
		j := 0
		for j < (colLen - 3) {
			firstCell := crosswordGame[i][j]

			if firstCell.value != runesToMatch[0] {
				j++
			} else {
				secondCell := crosswordGame[i-1][j+1]

				if secondCell.value != runesToMatch[1] {
					j++
				} else {
					thirdCell := crosswordGame[i-2][j+2]

					if thirdCell.value != runesToMatch[2] {
						// nothing, jump ahead
						j++
					} else {
						fourthCell := crosswordGame[i-3][j+3]

						if fourthCell.value == runesToMatch[3] {
							count++
							for offset := 0; offset < 4; offset++ {
								crosswordGame[i-offset][j+offset].inWord = true
							}
						}

						j++
					}
				}
			}
		}
	}

	return count
}

func initCrossWord(lines []string) [][]Cell {
	rowLen := len(lines)
	colLen := len(lines[0])

	crosswordGame := make([][]Cell, rowLen)
	for i := range crosswordGame {
		crosswordGame[i] = make([]Cell, colLen)
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			line := []rune(lines[i])
			rune := line[j]
			crosswordGame[i][j].value = rune
		}
	}

	return crosswordGame
}

func printCell(cell Cell) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	if cell.inWord {
		fmt.Printf("%s%c%s", colorRed, cell.value, colorReset)
	} else {
		fmt.Printf("%c", cell.value)
	}
}

func day4Part1(test bool) int {
	lines := utils.ReadFileLines("./inputs/day4.input.txt")
	count := 0

	crosswordGame := initCrossWord(lines)

	count += scanHorizontal(&crosswordGame, "XMAS")
	count += scanHorizontal(&crosswordGame, "SAMX")

	count += scanVertical(&crosswordGame, "XMAS")
	count += scanVertical(&crosswordGame, "SAMX")

	count += scanDiagonalTopToBottom(&crosswordGame, "XMAS")
	count += scanDiagonalTopToBottom(&crosswordGame, "SAMX")

	count += scanDiagonalBottomToTop(&crosswordGame, "XMAS")
	count += scanDiagonalBottomToTop(&crosswordGame, "SAMX")

	columnLen := len(lines[0])
	rowLen := len(lines)

	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			printCell(crosswordGame[i][j])
		}
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("found", count, "words")
	fmt.Println("colNo:", columnLen, "rowNo:", rowLen)
	return count
}

func day4Part2(test bool) int {
	lines := utils.ReadFileLines("./inputs/day4.input.txt")
	crosswordGame := initCrossWord(lines)
	count := scanHorizontalXMas(&crosswordGame)

	colLen := len(lines[0])
	rowLen := len(lines)

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			printCell(crosswordGame[i][j])
		}
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("found", count, "words")
	fmt.Println("colNo:", colLen, "rowNo:", rowLen)
	return count
}
