package main

import (
	"strings"
)

func simplePosition() int {
	lines := readLines("day2.txt")
	horizontalPosition := 0
	depth := 0

	for _, line := range lines {
		components := strings.Split(line, " ")
		if components[0] == "forward" {
			horizontalPosition += intValue(components[1])
		} else if components[0] == "up" {
			depth -= intValue(components[1])
		} else {
			depth += intValue(components[1])
		}
	}
	return horizontalPosition * depth
}

func aimedPosition() int {
	lines := readLines("day2.txt")
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		components := strings.Split(line, " ")
		value := intValue(components[1])
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
