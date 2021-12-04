package main

import (
	"fmt"

	"github.com/gossie/adventofcode2021/day1"
	"github.com/gossie/adventofcode2021/day2"
	"github.com/gossie/adventofcode2021/day3"
	"github.com/gossie/adventofcode2021/day4"
)

func main() {
	fmt.Println("Day 1, task 1: ", day1.NumberOfLargerMeasurements("day1/day1.txt"))
	fmt.Println("Day 1, task 2: ", day1.NumberOfLargerSumedMeasurements("day1/day1.txt"))

	fmt.Println("Day 2, task 1: ", day2.SimplePosition("day2/day2.txt"))
	fmt.Println("Day 2, task 2: ", day2.AimedPosition("day2/day2.txt"))

	fmt.Println("Day 3, task 1: ", day3.PowerConsumption("day3/day3.txt"))
	fmt.Println("Day 3, task 2: ", day3.LifeSupport("day3/day3.txt"))

	fmt.Println("Day 4, task 1: ", day4.BingoFirstWin("day4/day4.txt"))
	fmt.Println("Day 4, task 2: ", day4.BingoLastWin("day4/day4.txt"))
}
