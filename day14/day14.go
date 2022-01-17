package day14

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func readData(filename string) (string, map[string]string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	polymer := ""
	rules := make(map[string]string)

	scanRules := false
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanRules = true
			continue
		}

		if !scanRules {
			polymer = line
		} else {
			row := strings.Split(line, " -> ")
			rules[row[0]] = row[1]
		}
	}
	return polymer, rules
}

func initializePairs(polymer string) map[string]int64 {
	pairs := make(map[string]int64)
	for j := 0; j < len(polymer)-1; j++ {
		pairs[polymer[j:j+2]]++
	}
	return pairs
}

func calculateQuantities(polymer string, pairs map[string]int64, rules map[string]string, iterations int) map[rune]int64 {
	letterQuantities := make(map[rune]int64)
	for _, letter := range polymer {
		letterQuantities[letter]++
	}

	for i := 0; i < iterations; i++ {
		newPairs := make(map[string]int64)
		for pair, quantity := range pairs {
			toInsert := rules[pair]
			letterQuantities[rune(toInsert[0])] += quantity
			newPairs[string(pair[0])+toInsert] += quantity
			newPairs[toInsert+string(pair[1])] += quantity
		}
		pairs = newPairs
	}
	return letterQuantities
}

func calculateMinAndMax(letterQuantities map[rune]int64) (int64, int64) {
	var min int64 = math.MaxInt
	var max int64 = math.MinInt
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

func Quantities(filename string, iterations int) int64 {
	polymer, rules := readData(filename)

	min, max := calculateMinAndMax(
		calculateQuantities(
			polymer,
			initializePairs(polymer),
			rules,
			iterations,
		),
	)

	return max - min
}
