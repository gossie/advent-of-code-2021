package main

import (
	"fmt"

	"github.com/gossie/adventofcode2021/day1"
	"github.com/gossie/adventofcode2021/day10"
	"github.com/gossie/adventofcode2021/day11"
	"github.com/gossie/adventofcode2021/day12"
	"github.com/gossie/adventofcode2021/day13"
	"github.com/gossie/adventofcode2021/day14"
	"github.com/gossie/adventofcode2021/day15"
	"github.com/gossie/adventofcode2021/day16"
	"github.com/gossie/adventofcode2021/day17"
	"github.com/gossie/adventofcode2021/day18"
	"github.com/gossie/adventofcode2021/day19"
	"github.com/gossie/adventofcode2021/day2"
	"github.com/gossie/adventofcode2021/day20"
	"github.com/gossie/adventofcode2021/day21"
	"github.com/gossie/adventofcode2021/day22"
	"github.com/gossie/adventofcode2021/day25"
	"github.com/gossie/adventofcode2021/day3"
	"github.com/gossie/adventofcode2021/day4"
	"github.com/gossie/adventofcode2021/day5"
	"github.com/gossie/adventofcode2021/day6"
	"github.com/gossie/adventofcode2021/day7"
	"github.com/gossie/adventofcode2021/day8"
	"github.com/gossie/adventofcode2021/day9"
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

	fmt.Println("\nPerforming tasks of day 9")
	fmt.Println("Day 9, task 1: ", day9.RiskLevel("day9/day9.txt"))
	fmt.Println("Day 9, task 2: ", day9.Basin("day9/day9.txt"))

	fmt.Println("\nPerforming tasks of day 10")
	fmt.Println("Day 10, task 1: ", day10.ScoreForBrokenLines("day10/day10.txt"))
	fmt.Println("Day 10, task 1: ", day10.ScoreForIncompleteLines("day10/day10.txt"))

	fmt.Println("\nPerforming tasks of day 11")
	fmt.Println("Day 11, task 1: ", day11.NumberOfFlashes("day11/day11.txt"))
	fmt.Println("Day 11, task 2: ", day11.StepWhenAllFlash("day11/day11.txt"))

	fmt.Println("\nPerforming tasks of day 12")
	fmt.Println("Day 12, task 1: ", day12.NumberOfPaths("day12/day12.txt", true))
	fmt.Println("Day 12, task 2: ", day12.NumberOfPaths("day12/day12.txt", false))

	fmt.Println("\nPerforming tasks of day 13")
	fmt.Println("Day 13, task 1: ", day13.AfterOneFold("day13/day13.txt"))
	fmt.Println("Day 13, task 2: ")
	day13.Code("day13/day13.txt")

	fmt.Println("\nPerforming tasks of day 14")
	fmt.Println("Day 14, task 1: ", day14.Quantities("day14/day14.txt", 10))
	fmt.Println("Day 14, task 1: ", day14.Quantities("day14/day14.txt", 40))

	fmt.Println("\nPerforming tasks of day 15")
	fmt.Println("Day 15, task 1: ", day15.MinimalRisk("day15/day15.txt", 1))
	fmt.Println("Day 15, task 2: ", day15.MinimalRisk("day15/day15.txt", 5))

	fmt.Println("\nPerforming tasks of day 16")
	fmt.Println("Day 16, task 1: ", day16.Versions("day16/day16.txt"))
	fmt.Println("Day 16, task 2: ", day16.Calculate("day16/day16.txt"))

	fmt.Println("\nPerforming tasks of day 17")
	fmt.Println("Day 17, task 1: ", day17.Heighest("day17/day17.txt"))
	fmt.Println("Day 17, task 1: ", day17.HowMany("day17/day17.txt"))

	fmt.Println("\nPerforming tasks of day 18")
	fmt.Println("Day 18, task 1: ", day18.Magnitude("day18/day18.txt"))
	fmt.Println("Day 18, task 2: ", day18.LargestMagnitude("day18/day18.txt"))

	fmt.Println("\nPerforming tasks of day 19")
	fmt.Println("Day 19, task 1: ", day19.DistinctBeacons("day19/day19.txt"))
	fmt.Println("Day 19, task 2: ", day19.ManhattenDistance("day19/day19.txt"))

	fmt.Println("\nPerforming tasks of day 20")
	fmt.Println("Day 20, task 1: ", day20.NumberOfLitPixels("day20/day20.txt", 2))
	fmt.Println("Day 20, task 2: ", day20.NumberOfLitPixels("day20/day20.txt", 50))

	fmt.Println("\nPerforming tasks of day 21")
	fmt.Println("Day 21, task 1: ", day21.PlayTestGame())
	fmt.Println("Day 21, task 2: ", day21.MultipleUniverses())

	fmt.Println("\nPerforming tasks of day 22")
	fmt.Println("Day 22, task 1: ", day22.NumberOfEnabledCubes("day22/day22.txt", true))
	// fmt.Println("Day 22, task 2: ", day22.NumberOfEnabledCubes("day22/day22.txt", false))

	fmt.Println("\nPerforming tasks of day 24")
	// fmt.Println("Day 24, task 1: ", day24.ModelNumber("day24/day24.txt"))

	fmt.Println("\nPerforming tasks of day 25")
	fmt.Println("Day 25, task 1: ", day25.WhichStep("day25/day25.txt"))
}
