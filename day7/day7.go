package day7

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func readData(filename string) []int {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	startPositions := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, p := range strings.Split(scanner.Text(), ",") {
		startPosition, err := strconv.Atoi(p)
		if err != nil {
			panic("position cannot be parsed")
		}
		startPositions = append(startPositions, startPosition)
	}

	return startPositions
}

func FuelConstantConsumptions(file string) int {
	startPositions := readData(file)

	maxPosition := math.MinInt
	for _, p := range startPositions {
		maxPosition = int(math.Max(float64(p), float64(maxPosition)))
	}

	fuel := math.MaxInt
	for i := 0; i <= maxPosition; i++ {
		current := 0
		for _, p := range startPositions {
			current += int(math.Abs(float64((p - i))))
		}
		fuel = int(math.Min(float64(fuel), float64(current)))
	}
	return fuel
}

func FuelLinearConsumption(file string) int {
	startPositions := readData(file)

	minPosition := math.MaxInt
	maxPosition := math.MinInt
	for _, p := range startPositions {
		minPosition = int(math.Min(float64(p), float64(minPosition)))
		maxPosition = int(math.Max(float64(p), float64(maxPosition)))
	}

	fuel := math.MaxInt
	for i := minPosition; i <= maxPosition; i++ {
		current := 0
		for _, p := range startPositions {
			diff := int(math.Abs(float64((p - i))))
			current += (diff * (diff + 1)) / 2
		}
		fuel = int(math.Min(float64(fuel), float64(current)))
	}
	return fuel
}
