package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func startPopulation(filename string) map[int]int64 {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	fishes := make(map[int]int64)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, fishAge := range strings.Split(scanner.Text(), ",") {
		fishAgeAsInt, err := strconv.Atoi(fishAge)
		if err != nil {
			panic("fish cannot be parsed")
		}
		fishes[fishAgeAsInt]++
	}

	return fishes
}

func FishPopulation(file string, days int) int64 {
	fishes := startPopulation(file)
	for i := 0; i < days; i++ {
		nextGeneration := make(map[int]int64)
		for i := 1; i < 9; i++ {
			nextGeneration[i-1] = fishes[i]
		}
		nextGeneration[6] += fishes[0]
		nextGeneration[8] = fishes[0]
		fishes = nextGeneration
	}

	var sum int64 = 0
	for _, amount := range fishes {
		sum += amount
	}

	return sum
}
