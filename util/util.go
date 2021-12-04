package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
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
