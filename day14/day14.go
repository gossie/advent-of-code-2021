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

	for i := 0; i < iterations; i++ {
		newPolymer := ""
		for j := 0; j < len(polymer)-1; j++ {
			key := polymer[j : j+2]
			toInsert := mapping[key]
			newPolymer = newPolymer + string(polymer[j]) + toInsert
			if j == len(polymer)-2 {
				newPolymer += string(polymer[j+1])
			}
		}
		polymer = newPolymer
	}

	quantities := make(map[rune]int)
	for _, letter := range polymer {
		quantities[letter]++
	}

	min := math.MaxInt
	max := math.MinInt

	for _, value := range quantities {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return max - min
}
