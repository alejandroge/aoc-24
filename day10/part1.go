package day10

import (
	"fmt"

	"aoc-24/utils"
)

type MapCell struct {
	elevation int
	posX      int
	posY      int
}

func initMap(lines []string) [][]MapCell {
	mapLen := len(lines)

	gameMap := make([][]MapCell, mapLen)
	for i := range gameMap {
		gameMap[i] = make([]MapCell, mapLen)
	}

	for y := 0; y < mapLen; y++ {
		for x := 0; x < mapLen; x++ {
			line := []rune(lines[y])
			rune := line[x]

			gameMap[y][x].elevation = utils.RuneToInt(rune)
			gameMap[y][x].posX = x
			gameMap[y][x].posY = y
		}
	}

	return gameMap
}

func printMapCell(cell MapCell) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	if cell.elevation == 0 {
		fmt.Printf("%s%d%s", colorRed, cell.elevation, colorReset)
	} else {
		fmt.Printf("%d", cell.elevation)
	}
}

func printMap(topoMapPtr *[][]MapCell) {
	topoMap := *topoMapPtr
	mapLen := len(topoMap)

	for y := 0; y < mapLen; y++ {
		for x := 0; x < mapLen; x++ {
			printMapCell(topoMap[y][x])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func Part1(test bool) int {
	path := utils.GetFilePath(10, test)
	lines := utils.ReadFileLines(path)

	topoMap := initMap(lines)
	printMap(&topoMap)

	trailHeadScoresSum := 0
	return trailHeadScoresSum
}
