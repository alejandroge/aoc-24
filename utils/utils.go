package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFilePath(day int, test bool) string {
	fileKind := "input"
	if test {
		fileKind = "test"
	}

	return fmt.Sprintf("./inputs/day%d.%s.txt", day, fileKind)
}

func ReadFileLines(filePath string) []string {
	var lines []string

	// Open the file
	file, err := os.Open(filePath)
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

func StringToInt(s string) int {
	number, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Error converting the string to a number")
		return -1
	}

	return number
}

func RuneToInt(r rune) int {
	return int(r - '0')
}

func Absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
