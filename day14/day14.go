package day14

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func readData(filename string) (string, map[string]string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	polymer := ""
	mapping := make(map[string]string)

	scanMapping := false
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanMapping = true
			continue
		}

		if !scanMapping {
			polymer = line
		} else {
			row := strings.Split(line, " -> ")
			mapping[row[0]] = row[1]
		}
	}
	return polymer, mapping
}

func initializePairs(polymer string) map[string]int {
	pairs := make(map[string]int)
	for j := 0; j < len(polymer)-1; j++ {
		pairs[polymer[j:j+2]]++
	}
	return pairs
}

func calculateQuantities(polymer string, pairs map[string]int, mapping map[string]string, iterations int) map[rune]int {
	letterQuantities := make(map[rune]int)
	for _, letter := range polymer {
		letterQuantities[letter]++
	}

	for i := 0; i < iterations; i++ {
		newPairs := make(map[string]int)
		for pair, quantity := range pairs {
			toInsert := mapping[pair]
			letterQuantities[rune(toInsert[0])] += quantity
			newPairs[string(pair[0])+toInsert] += quantity
			newPairs[toInsert+string(pair[1])] += quantity
		}
		pairs = newPairs
	}
	return letterQuantities
}

func calculateMinAndMax(letterQuantities map[rune]int) (int, int) {
	min := math.MaxInt
	max := math.MinInt
	for _, value := range letterQuantities {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func Quantities(filename string, iterations int) int {
	polymer, mapping := readData(filename)

	min, max := calculateMinAndMax(
		calculateQuantities(
			polymer,
			initializePairs(polymer),
			mapping,
			iterations,
		),
	)

	return max - min
}
