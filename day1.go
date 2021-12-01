package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readLines() []string {
	file, err := os.Open("puzzle-input.txt")
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

func intValue(lines []string, index int) int {
	number, err := strconv.Atoi(lines[index])
	if err != nil {
		panic("Line is not a number")
	}
	return number
}

func determineNumberOfLargerMeasurements() int {
	raises := 0
	var prev *int = nil
	for _, eachline := range readLines() {
		number, err := strconv.Atoi(eachline)
		if err != nil {
			panic("Line is not a number")
		}

		if prev != nil && number > *prev {
			raises++
		}

		prev = &number
	}

	return raises
}

func determineNumberOfLargerSumedMeasurements() int {
	lines := readLines()

	raises := 0
	var prev *int = nil
	for i := 0; i < len(lines)-2; i++ {
		number1 := intValue(lines, i)
		number2 := intValue(lines, i+1)
		number3 := intValue(lines, i+2)
		current := number1 + number2 + number3

		if prev != nil && current > *prev {
			raises++
		}

		prev = &current
	}

	return raises
}
