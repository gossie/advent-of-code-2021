package day23

import (
	"bufio"
	"os"
)

func parse(filename string) [][]rune {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	field := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := make([]rune, 0)
		for _, c := range scanner.Text() {
			row = append(row, c)
		}
		field = append(field, row)
	}
	return field
}

func finished(field [][]rune) bool {
	return field[2][3] == 'A' && field[3][3] == 'A' && field[2][5] == 'B' && field[3][5] == 'B' && field[2][7] == 'C' && field[3][7] == 'C' && field[2][9] == 'D' && field[3][9] == 'D'
}

func LeastEnergy(filename string) int {
	field := parse(filename)
	// energyMapping := map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
	for !finished(field) {

	}

	return 0
}
