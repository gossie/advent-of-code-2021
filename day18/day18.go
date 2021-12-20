package day18

import (
	"math"

	"github.com/gossie/adventofcode2021/day18/parser"
)

func Magnitude(filename string) int {
	numbers := parser.ReadData(filename)
	current := numbers[0]

	for i := 1; i < len(numbers); i++ {
		current = current.Add(numbers[i])
		current = current.ReduceCompletey()
	}
	return current.Magnitude()
}

func LargestMagnitude(filename string) int {
	numbers := parser.ReadData(filename)

	max := math.MinInt
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				numbers = parser.ReadData(filename) // that sucks! imutability for the win
				magnitude := numbers[i].Add(numbers[j]).ReduceCompletey().Magnitude()
				max = int(math.Max(float64(magnitude), float64(max)))

				numbers = parser.ReadData(filename)
				magnitude = numbers[j].Add(numbers[i]).ReduceCompletey().Magnitude()
				max = int(math.Max(float64(magnitude), float64(max)))
			}
		}
	}
	return max
}
