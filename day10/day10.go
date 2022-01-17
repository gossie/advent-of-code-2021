package day10

import (
	"bufio"
	"os"
	"sort"
)

func readData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	return lines
}

func contains(slice []rune, character rune) bool {
	for _, r := range slice {
		if r == character {
			return true
		}
	}
	return false
}

func broken(stack []rune, character rune, mapping map[rune]rune) bool {
	last := stack[len(stack)-1]
	return character != mapping[last]
}

func ScoreForBrokenLines(filename string) int64 {
	points := map[rune]int64{')': 3, ']': 57, '}': 1197, '>': 25137}
	closing := []rune{')', ']', '}', '>'}
	mapping := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

	score := int64(0)
	for _, line := range readData(filename) {
		stack := make([]rune, 0)
		for _, p := range line {
			if contains(closing, p) {
				if broken(stack, p, mapping) {
					score += points[p]
					break
				} else {
					stack = stack[0 : len(stack)-1]
				}
			} else {
				stack = append(stack, p)
			}
		}
	}
	return score
}

func ScoreForIncompleteLines(filename string) int64 {
	points := map[rune]int64{')': 1, ']': 2, '}': 3, '>': 4}
	closing := []rune{')', ']', '}', '>'}
	mapping := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

	scores := make([]int64, 0)
	for _, line := range readData(filename) {
		score := int64(0)
		stack := make([]rune, 0)
		for cIndex, p := range line {
			if contains(closing, p) {
				if broken(stack, p, mapping) {
					break
				} else {
					stack = stack[0 : len(stack)-1]
				}
			} else {
				stack = append(stack, p)
			}

			if cIndex == len(line)-1 {
				for i := len(stack) - 1; i >= 0; i-- {
					score = 5*score + points[mapping[stack[i]]]
				}
				scores = append(scores, score)
			}
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	return scores[len(scores)/2]
}
