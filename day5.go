package main

import (
	"fmt"
	"strings"

	"aoc-24/utils"
)

func splitAndConvertLine(line string, delimiter string) []int {
	var result []int

	splittedStrings := strings.Split(line, delimiter)

	for _, numberAsString := range splittedStrings {
		result = append(result, utils.StringToInt(numberAsString))
	}

	return result
}

func splitRulesAndUpdates(lines []string) ([][]int, [][]int) {
	var rules [][]int
	var updates [][]int

	inRulesSection := true

	for _, line := range lines {
		if line == "" {
			inRulesSection = false
			continue
		}

		if inRulesSection {
			convertedRule := splitAndConvertLine(line, "|")
			rules = append(rules, convertedRule)
		} else {
			convertedUpdate := splitAndConvertLine(line, ",")
			updates = append(updates, convertedUpdate)
		}
	}

	return rules, updates
}

func getApplyingRules(page int, rules [][]int) ([]int, []int) {
	var beforePages []int
	var afterPages []int

	for _, rule := range rules {
		if rule[0] == page {
			beforePages = append(beforePages, rule[1])
		}

		if rule[1] == page {
			afterPages = append(afterPages, rule[0])
		}
	}

	return beforePages, afterPages
}

func contains(numbersArray []int, target int) bool {
	isPresent := false

	for _, elem := range numbersArray {
		if elem == target {
			isPresent = true
			break
		}
	}

	return isPresent
}

func checkPages(index int, update []int, allowedNumbers []int, direction string) bool {
	switch direction {
	case "asc":
		valid := true

		for i := index + 1; i < len(update); i++ {
			pageToCheck := update[i]

			if !contains(allowedNumbers, pageToCheck) {
				valid = false
				break
			}
		}

		return valid
	case "desc":
		valid := true
		for i := index - 1; i >= 0; i-- {
			pageToCheck := update[i]

			if !contains(allowedNumbers, pageToCheck) {
				valid = false
				break
			}
		}
		return valid
	default:
		fmt.Println("bad direction, so you get false")
		return false
	}
}

func isUpdateValid(update []int, rules [][]int) bool {
	updateValid := true

	for idx, pageNumber := range update {
		beforePages, afterPages := getApplyingRules(pageNumber, rules)
		beforeValid := true
		afterValid := true

		if idx > 0 {
			if len(afterPages) > 0 {
				afterValid = checkPages(idx, update, afterPages, "desc")
			}
		}

		if idx < len(update)-1 {
			if len(beforePages) == 0 {
				beforeValid = true
			} else {
				beforeValid = checkPages(idx, update, beforePages, "asc")
			}
		}

		updateValid = beforeValid && afterValid

		if !updateValid {
			// fmt.Println("Reporting error on position idx:", idx, "checkingPage", pageNumber)
			// fmt.Println("beforeValid", beforeValid, "afterValid", afterValid)

			// fmt.Println("beforePages", beforePages)
			// fmt.Println("afterPages", afterPages)

			break
		}
	}

	return updateValid
}

func day5Part1(test bool) int {
	lines := utils.ReadFileLines("./inputs/day5.input.txt")
	rules, updates := splitRulesAndUpdates(lines)

	var validUpdates [][]int
	for updateIdx, update := range updates {
		valid := isUpdateValid(update, rules)

		if valid {
			fmt.Print(updateIdx, ": ✅\t")
			validUpdates = append(validUpdates, update)
		} else {
			fmt.Print(updateIdx, ": ❌\t")
		}

		fmt.Println(update)

	}

	count := 0
	for _, update := range validUpdates {
		updateLength := len(update)
		middleIndex := updateLength / 2
		// fmt.Println("odd len?", updateLength, "middle idx", updateLength / 2)

		count += update[middleIndex]
	}

	fmt.Println("Result:", count) // 4872
	return count
}
