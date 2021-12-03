package main

import (
	"math"
	"strings"
)

type selector func([]int, []int) []int

func distributionOfBits(binary string, numberOfOnesPerPlace map[int]int) {
	for i, digit := range binary {
		if digit == '1' {
			numberOfOnesPerPlace[i] = numberOfOnesPerPlace[i] + 1
		}
	}
}

func binaryToDecimal(binary string) int {
	decimal := 0
	for i, digit := range binary {
		if digit == '1' {
			decimal += int(math.Pow(2, float64((len(binary)-1)-i)))
		}
	}
	return decimal
}

func powerConsumption() int {
	lines := readLines("day3.txt")
	numberOfLines := len(lines)
	numberOfOnesPerPlace := make(map[int]int)
	lineLength := 0

	for _, line := range lines {
		lineLength = len(line)
		distributionOfBits(strings.Trim(line, " "), numberOfOnesPerPlace)
	}

	epsilon := 0
	gamma := 0
	for bitIndex, numberOfOnes := range numberOfOnesPerPlace {
		if numberOfOnes > numberOfLines/2 {
			gamma |= int(math.Pow(2, float64((lineLength-1)-bitIndex)))
		} else {
			epsilon |= int(math.Pow(2, float64((lineLength-1)-bitIndex)))
		}
	}
	return epsilon * gamma
}

func lifeSupport() int {
	lines := readLines("day3.txt")
	bitmasks := make(map[int]int)
	decimalNumbers := make([]int, 0, len(lines))

	bitmasks[0] = 0b100000000000
	bitmasks[1] = 0b010000000000
	bitmasks[2] = 0b001000000000
	bitmasks[3] = 0b000100000000
	bitmasks[4] = 0b000010000000
	bitmasks[5] = 0b000001000000
	bitmasks[6] = 0b000000100000
	bitmasks[7] = 0b000000010000
	bitmasks[8] = 0b000000001000
	bitmasks[9] = 0b000000000100
	bitmasks[10] = 0b000000000010
	bitmasks[11] = 0b000000000001

	for _, line := range lines {
		decimalNumbers = append(decimalNumbers, binaryToDecimal(strings.Trim(line, " ")))
	}

	oxygenGeneratorRating := searchRating(decimalNumbers, bitmasks, 0, oxygenRatingSelector)
	co2ScrubberRating := searchRating(decimalNumbers, bitmasks, 0, co2RatingSelector)

	return oxygenGeneratorRating * co2ScrubberRating
}

func searchRating(decimalNumbers []int, bitmasks map[int]int, maskIndex int, newNumbersSelector selector) int {
	if len(decimalNumbers) == 1 {
		return decimalNumbers[0]
	}

	return searchRating(
		newNumbersSelector(divideNumbers(decimalNumbers, bitmasks, maskIndex)),
		bitmasks,
		maskIndex+1,
		newNumbersSelector)
}

func divideNumbers(decimalNumbers []int, bitmasks map[int]int, maskIndex int) ([]int, []int) {
	zeros := make([]int, 0)
	ones := make([]int, 0)
	for _, decimal := range decimalNumbers {
		if decimal&bitmasks[maskIndex] == 0 {
			zeros = append(zeros, decimal)
		} else {
			ones = append(ones, decimal)
		}
	}
	return zeros, ones
}

func co2RatingSelector(zeros []int, ones []int) []int {
	if len(ones) < len(zeros) {
		return ones
	}
	return zeros
}

func oxygenRatingSelector(zeros []int, ones []int) []int {
	if len(zeros) > len(ones) {
		return zeros
	}
	return ones
}
