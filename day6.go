package main

import "fmt"

type MapCell struct {
	status      rune
	posX        int
	posY        int
	visitNumber int
}

type Guard struct {
	positionPtr *MapCell
	orientation rune
}

func initMap(lines []string) ([][]MapCell, Guard) {
	mapLen := len(lines)

	gameMap := make([][]MapCell, mapLen)
	for i := range gameMap {
		gameMap[i] = make([]MapCell, mapLen)
	}

	guard := Guard{positionPtr: nil, orientation: 'u'}

	for y := 0; y < mapLen; y++ {
		for x := 0; x < mapLen; x++ {
			line := []rune(lines[y])
			rune := line[x]

			gameMap[y][x].posX = x
			gameMap[y][x].posY = y
			gameMap[y][x].status = rune
			gameMap[y][x].visitNumber = 0

			if rune == '^' {
				gameMap[y][x].visitNumber = 1

				guard.positionPtr = &gameMap[y][x]
				guard.orientation = 'u'
			}
		}
	}

	return gameMap, guard
}

func printMapCell(cell MapCell) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	if cell.status == '#' {
		fmt.Printf("%s%c%s", colorRed, cell.status, colorReset)
	} else {
		fmt.Printf("%c", cell.status)
	}
}

func printMap(gameMapPtr *[][]MapCell) {
	gameMap := *gameMapPtr
	mapLen := len(gameMap)

	for y := 0; y < mapLen; y++ {
		for x := 0; x < mapLen; x++ {
			printMapCell(gameMap[y][x])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func getNextCell(guardPtr *Guard, gameMapPtr *[][]MapCell) (*MapCell, error) {
	gameMap := *gameMapPtr
	mapLen := len(gameMap)

	guardCurrentCellPtr := guardPtr.positionPtr
	y := guardCurrentCellPtr.posY
	x := guardCurrentCellPtr.posX

	switch guardPtr.orientation {
	case 'u':
		if (y - 1) < 0 {
			return nil, fmt.Errorf("tried to move out of bounds")
		}
		return &(gameMap[y-1][x]), nil
	case 'r':
		if (x + 1) >= mapLen {
			return nil, fmt.Errorf("tried to move out of bounds")
		}
		return &(gameMap[y][x+1]), nil
	case 'd':
		if (y + 1) >= mapLen {
			return nil, fmt.Errorf("tried to move out of bounds")
		}
		return &(gameMap[y+1][x]), nil
	case 'l':
		if (x - 1) < 0 {
			return nil, fmt.Errorf("tried to move out of bounds")
		}
		return &(gameMap[y][x-1]), nil
	default:
		return nil, fmt.Errorf("unexpected orientation")
	}
}

func moveGuardForward(guardPtr *Guard, gameMapPtr *[][]MapCell) error {
	nextCellPtr, err := getNextCell(guardPtr, gameMapPtr)

	if err != nil {
		return fmt.Errorf("tried to move out of bounds, play should be over")
	}

	for nextCellPtr.status != '#' {
		guardPtr.positionPtr.status = 'X'
		guardPtr.positionPtr = nextCellPtr

		nextCellPtr.visitNumber++
		nextCellPtr.status = '^'

		nextCellPtr, err = getNextCell(guardPtr, gameMapPtr)
		if err != nil {
			return fmt.Errorf("tried to move out of bounds, play should be over")
		}
	}

	return nil
}

func turnGuard(guardPtr *Guard) error {
	switch guardPtr.orientation {
	case 'u':
		guardPtr.orientation = 'r'
	case 'r':
		guardPtr.orientation = 'd'
	case 'd':
		guardPtr.orientation = 'l'
	case 'l':
		guardPtr.orientation = 'u'
	default:
		return fmt.Errorf("something went wrong, this is not an expected orientation")
	}
	return nil
}

func getNumberOfVisitedCells(gameMapPtr *[][]MapCell) int {
	gameMap := *gameMapPtr
	mapLen := len(gameMap)

	countOfVisited := 0
	for y := 0; y < mapLen; y++ {
		for x := 0; x < mapLen; x++ {
			currentCell := gameMap[y][x]
			if currentCell.visitNumber > 0 {
				countOfVisited++
			}
		}
	}

	return countOfVisited
}

func day6Part1() {
	lines := readFileLines("./inputs/day6.input.txt")
	gameMap, guard := initMap(lines)

	for {
		if err := moveGuardForward(&guard, &gameMap); err != nil {
			// tried to move outside of the map, should stop the loop at this point
			break
		}
		turnGuard(&guard)
	}

	printMap(&gameMap)
	numberOfVisitedCells := getNumberOfVisitedCells(&gameMap)

	fmt.Println("visited", numberOfVisitedCells, "different cells")
}

func day6Part2() {
	lines := readFileLines("./inputs/day6.test.txt")
	gameMap, guard := initMap(lines)

	// main loop for moving the guard around
	for i := 0; i < 100; i++ {
		if err := moveGuardForward(&guard, &gameMap); err != nil {
			// tried to move outside of the map, should stop the loop at this point
			break
		}
		turnGuard(&guard)
	}

	printMap(&gameMap)

	fmt.Println(guard)
}
