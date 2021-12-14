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

	letterQuantities := make(map[rune]int)
	pairs := make(map[string]int)
	for j := 0; j < len(polymer)-1; j++ {
		pairs[polymer[j:j+2]]++
		letterQuantities[rune(polymer[j])]++
		if j == len(polymer)-2 {
			letterQuantities[rune(polymer[j+1])]++
		}
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

	return max - min
}
