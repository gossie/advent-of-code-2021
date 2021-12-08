package main

import (
	"fmt"

	"github.com/gossie/adventofcode2021/day1"
	"github.com/gossie/adventofcode2021/day2"
	"github.com/gossie/adventofcode2021/day3"
	"github.com/gossie/adventofcode2021/day4"
	"github.com/gossie/adventofcode2021/day5"
	"github.com/gossie/adventofcode2021/day6"
	"github.com/gossie/adventofcode2021/day7"
	"github.com/gossie/adventofcode2021/day8"
)

func main() {
	fmt.Println("Performing tasks of day 1")
	fmt.Println("Day 1, task 1: ", day1.NumberOfLargerMeasurements("day1/day1.txt"))
	fmt.Println("Day 1, task 2: ", day1.NumberOfLargerSumedMeasurements("day1/day1.txt"))

	fmt.Println("\nPerforming tasks of day 2")
	fmt.Println("Day 2, task 1: ", day2.SimplePosition("day2/day2.txt"))
	fmt.Println("Day 2, task 2: ", day2.AimedPosition("day2/day2.txt"))

	fmt.Println("\nPerforming tasks of day 3")
	fmt.Println("Day 3, task 1: ", day3.PowerConsumption("day3/day3.txt"))
	fmt.Println("Day 3, task 2: ", day3.LifeSupport("day3/day3.txt"))

	fmt.Println("\nPerforming tasks of day 4")
	fmt.Println("Day 4, task 1: ", day4.BingoFirstWin("day4/day4.txt"))
	fmt.Println("Day 4, task 2: ", day4.BingoLastWin("day4/day4.txt"))

	fmt.Println("\nPerforming tasks of day 5")
	fmt.Println("Day 5, task 1: ", day5.AvoidDangerousArea("day5/day5.txt", false))
	fmt.Println("Day 5, task 2: ", day5.AvoidDangerousArea("day5/day5.txt", true))

	fmt.Println("\nPerforming tasks of day 6")
	fmt.Println("Day 6, task 1: ", day6.FishPopulation("day6/day6.txt", 80))
	fmt.Println("Day 6, task 2: ", day6.FishPopulation("day6/day6.txt", 256))

	fmt.Println("\nPerforming tasks of day 7")
	fmt.Println("Day 7, task 1: ", day7.FuelConstantConsumptions("day7/day7.txt"))
	fmt.Println("Day 7, task 2: ", day7.FuelLinearConsumption("day7/day7.txt"))

	fmt.Println("\nPerforming tasks of day 8")
	fmt.Println("Day 8, task 1: ", day8.NumberOf1And4And7And8("day8/day8.txt"))
	fmt.Println("Day 8, task 2: ", day8.OutputSum("day8/day8.txt"))
}
