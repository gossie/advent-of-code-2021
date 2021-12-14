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

func Quantities(filename string, iterations int) int {
	polymer, mapping := readData(filename)

	pairs := make(map[string]int)
	for j := 0; j < len(polymer)-1; j++ {
		pairs[polymer[j:j+2]]++
	}

	for i := 0; i < iterations; i++ {
		newPairs := make(map[string]int)
		for pair, quantity := range pairs {
			toInsert := mapping[pair]
			newPairs[string(pair[0])+toInsert] += quantity
			newPairs[toInsert+string(pair[1])] += quantity
		}
		pairs = newPairs
	}

	letterQuatities := make(map[rune]int)
	for pair, quantity := range pairs {
		for _, letter := range pair {
			letterQuatities[letter] += quantity
		}
	}

	min := math.MaxInt
	max := math.MinInt

	for letter, value := range letterQuatities {
		actualValue := value / 2
		if letter == rune(polymer[0]) || letter == rune(polymer[len(polymer)-1]) {
			actualValue++
		}

		if actualValue < min {
			min = actualValue
		}

		if actualValue > max {
			max = actualValue
		}
	}

	return max - min
}
