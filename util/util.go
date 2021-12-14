package util

import (
	"bufio"
	"os"
	"strconv"
)

func LoadFile(filename string) *os.File {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}
	return file
}

func ReadLines(filename string) []string {
	scanner := bufio.NewScanner(LoadFile(filename))
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func IntValue(line string) int {
	number, err := strconv.Atoi(line)
	if err != nil {
		panic("Line is not a number")
	}
	return number
}
