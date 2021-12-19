package parser

import (
	"bufio"
	"os"
	"strconv"

	"github.com/gossie/adventofcode2021/day18/tree"
)

func indexOfSeperatingComma(number string) int {
	numberOfOpensBeforeClose := 0
	for i, c := range number {
		if c == ']' {
			numberOfOpensBeforeClose--
			if numberOfOpensBeforeClose == 0 {
				return i + 1
			}
		}
		if c == '[' {
			numberOfOpensBeforeClose++
		}
	}
	panic("no comma")
}

func parsePair(number string) *tree.Node {
	index := 1
	if number[0] == '[' {
		index = indexOfSeperatingComma(number)
	}
	left := parseNumber(number[0:index])
	var right *tree.Node
	if number[index+1] == '[' {
		right = parseNumber(number[index+1:])
	} else {
		right = parseNumber(number[index+1:])
	}
	newNode := tree.Node{Left: left, Right: right}
	newNode.Left.Parent = &newNode
	newNode.Right.Parent = &newNode
	return &newNode
}

func parseNumber(line string) *tree.Node {
	if line[0] == '[' {
		return parsePair(line[1 : len(line)-1])
	} else {
		n, _ := strconv.Atoi(line)
		return &tree.Node{Value: n}
	}
}

func ReadData(filename string) []*tree.Node {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	numbers := []*tree.Node{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		numbers = append(numbers, parseNumber(scanner.Text()))
	}

	return numbers
}
