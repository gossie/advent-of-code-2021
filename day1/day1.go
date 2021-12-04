package day1

import "github.com/gossie/adventofcode2021/util"

func NumberOfLargerMeasurements(file string) int {
	raises := 0
	var prev *int = nil
	for _, eachline := range util.ReadLines(file) {
		number := util.IntValue(eachline)
		if prev != nil && number > *prev {
			raises++
		}
		prev = &number
	}

	return raises
}

func NumberOfLargerSumedMeasurements(file string) int {
	lines := util.ReadLines(file)

	raises := 0
	var prev *int = nil
	for i := 0; i < len(lines)-2; i++ {
		number1 := util.IntValue(lines[i])
		number2 := util.IntValue(lines[i+1])
		number3 := util.IntValue(lines[i+2])
		current := number1 + number2 + number3

		if prev != nil && current > *prev {
			raises++
		}

		prev = &current
	}

	return raises
}
