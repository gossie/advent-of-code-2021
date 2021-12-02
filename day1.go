package main

func numberOfLargerMeasurements() int {
	raises := 0
	var prev *int = nil
	for _, eachline := range readLines("day1.txt") {
		number := intValue(eachline)
		if prev != nil && number > *prev {
			raises++
		}
		prev = &number
	}

	return raises
}

func numberOfLargerSumedMeasurements() int {
	lines := readLines("day1.txt")

	raises := 0
	var prev *int = nil
	for i := 0; i < len(lines)-2; i++ {
		number1 := intValue(lines[i])
		number2 := intValue(lines[i+1])
		number3 := intValue(lines[i+2])
		current := number1 + number2 + number3

		if prev != nil && current > *prev {
			raises++
		}

		prev = &current
	}

	return raises
}
