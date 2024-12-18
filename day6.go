package main

import "fmt"

type MapCell struct {
    // don't love this actually, status works kinda nicely for blocked, or initial position of guard, but feels a bit
    // weird when combined with "visited" logic, tho the visit counter kinda helps making it work. Might just leave like
    // that
	status      rune
    visitNumber int
}

type Guard struct {
    xPosition   int
    yPosition   int
    orientation rune
}

func initMap(lines []string) ([][]MapCell, Guard) {
	mapLen	:= len(lines)

	gameMap := make([][]MapCell, mapLen)
	for i := range(gameMap) {
		gameMap[i] = make([]MapCell, mapLen)
	}

    guard := Guard{xPosition: -1, yPosition: -1, orientation: 'u'}

	for y := 0; y < mapLen; y++ {
		for x := 0; x < mapLen; x++ {
			line := []rune(lines[y])
			rune := line[x]
			gameMap[y][x].status = rune
			gameMap[y][x].visitNumber = 0

            if rune == '^' {
                gameMap[y][x].visitNumber = 1

                guard.xPosition = x
                guard.yPosition = y
                guard.orientation = 'u'
            }
		}
	}

	return gameMap, guard
}

func printMapCell(cell MapCell) {
    fmt.Printf("%c", cell.status)
}

func nextCell(guard *Guard, gameMap *[][]GameCell, orientation rune) MapCell {
    mapLen := len(*gameMap)
    y, x := (*guard.yPosition, *guard.xPosition)

    switch *guard.orientation {
    case 'u':
        return *gameMap[y-1][x];
    case 'r':
        return *gameMap[y][x+1];
    case 'd':
        return *gameMap[y+1][x];
    case 'l':
        return *gameMap[y][x-1];
    }
}

func moveGuardForward(guard *Guard, gameMap *[][]MapCell) error {
    return nil
}

func turnGuard(guard *Guard) {
    currentOrientation := *guard.orientation

    switch currentOrientation {
    case 'u':
        *guard.orientation = 'r'
    case 'r':
        *guard.orientation = 'd'
    case 'd':
        *guard.orientation = 'l'
    case 'l':
        *guard.orientation = 'u'
    default:
        fmt.Println("something went wrong, this is not an expected orientation")
    }
}

func day6Part1() {
    lines := readFileLines("./inputs/day6.test.txt")
    gameMap, guard := initMap(lines)

    for y := 0; y < len(gameMap); y++ {
        for x := 0; x < len(gameMap); x++ {
            printMapCell(gameMap[y][x])
        }
        fmt.Println("")
    }
    fmt.Println("")

    // main loop for moving the guard around. Limit to 100 now, so I don't loop infinitely all the time during testing
    for i := 0; i < 100; i++ {
        if err := moveGuardForward(&guard, &gameMap); err != nil {
            // tried to move outside of the map, should stop the loop at this point
        }
        turnGuard(&guard)
    }

    visitedAtLeastOnce := getVisitedCells(&gameMap)

    fmt.Println("guard is on x:", guard.xPosition, "and y:", guard.yPosition)
    fmt.Println("visited", len(visitedCells), "different cells")
}

func day6Part2() {
    lines := readFileLines("./inputs/day6.test.txt")
    gameMap, guard := initMap(lines)

    for y := 0; y < len(gameMap); y++ {
        for x := 0; x < len(gameMap); x++ {
            printMapCell(gameMap[y][x])
        }
        fmt.Println("")
    }
    fmt.Println("")

    // main loop for moving the guard around
    for i := 0; i < 100; i++ {
        if err := moveGuardForward(&guard, &gameMap); err != nil {
            // tried to move outside of the map, should stop the loop at this point
        }
        turnGuard(&guard)
    }

    fmt.Println(guard)
}
