package day2

import (
	"strings"

	"github.com/gossie/adventofcode2021/util"
)

func SimplePosition(file string) int {
	lines := util.ReadLines(file)
	horizontalPosition := 0
	depth := 0

	for _, line := range lines {
		components := strings.Split(line, " ")
		if components[0] == "forward" {
			horizontalPosition += util.IntValue(components[1])
		} else if components[0] == "up" {
			depth -= util.IntValue(components[1])
		} else {
			depth += util.IntValue(components[1])
		}
	}
	return horizontalPosition * depth
}

func AimedPosition(file string) int {
	lines := util.ReadLines(file)
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		components := strings.Split(line, " ")
		value := util.IntValue(components[1])
		if components[0] == "forward" {
			horizontalPosition += value
			depth += aim * value
		} else if components[0] == "up" {
			aim -= value
		} else {
			aim += value
		}
	}
	return horizontalPosition * depth
}
