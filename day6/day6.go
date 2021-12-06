package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func startPopulation(filename string) map[int]int {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	scanner := bufio.NewScanner(file)

	fishes := make(map[int]int)

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

func FishPopulation(file string, days int) int {
	fishes := startPopulation(file)
	for i := 0; i < days; i++ {
		nextGeneration := make(map[int]int)
		for i := 1; i < 9; i++ {
			nextGeneration[i-1] = fishes[i]
		}
		nextGeneration[6] += fishes[0]
		nextGeneration[8] = fishes[0]
		fishes = nextGeneration
	}

	sum := 0
	for _, amount := range fishes {
		sum += amount
	}

	return sum
}
