package day11

import (
	"bufio"
	"os"
	"strconv"
)

type octopus struct {
	energyLevel int
	flashed     bool
}

func (o *octopus) increaseEnergyLevel() {
	if !o.flashed {
		o.energyLevel++
	}
}

func (o *octopus) flash() {
	o.flashed = true
	o.energyLevel = 0
}

func readData(filename string) [][]octopus {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	octopuses := make([][]octopus, 0, 10)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := make([]octopus, 0, 10)
		text := scanner.Text()
		for i := 0; i < len(text); i++ {
			level, _ := strconv.Atoi(text[i : i+1])
			row = append(row, octopus{energyLevel: level})
		}
		octopuses = append(octopuses, row)

	}
	return octopuses
}

func flash(octopuses [][]octopus, y, x int) {
	if y > 0 && y < len(octopuses)-1 && x > 0 && x < len(octopuses[y])-1 {
		octopuses[y-1][x].increaseEnergyLevel()
		octopuses[y-1][x+1].increaseEnergyLevel()
		octopuses[y][x+1].increaseEnergyLevel()
		octopuses[y+1][x+1].increaseEnergyLevel()
		octopuses[y+1][x].increaseEnergyLevel()
		octopuses[y+1][x-1].increaseEnergyLevel()
		octopuses[y][x-1].increaseEnergyLevel()
		octopuses[y-1][x-1].increaseEnergyLevel()
	} else if y == 0 && x > 0 && x < len(octopuses[y])-1 {
		octopuses[y][x+1].increaseEnergyLevel()
		octopuses[y+1][x+1].increaseEnergyLevel()
		octopuses[y+1][x].increaseEnergyLevel()
		octopuses[y+1][x-1].increaseEnergyLevel()
		octopuses[y][x-1].increaseEnergyLevel()
	} else if y == len(octopuses)-1 && x > 0 && x < len(octopuses[y])-1 {
		octopuses[y-1][x].increaseEnergyLevel()
		octopuses[y-1][x+1].increaseEnergyLevel()
		octopuses[y][x+1].increaseEnergyLevel()
		octopuses[y][x-1].increaseEnergyLevel()
		octopuses[y-1][x-1].increaseEnergyLevel()
	} else if y > 0 && y < len(octopuses)-1 && x == 0 {
		octopuses[y-1][x].increaseEnergyLevel()
		octopuses[y-1][x+1].increaseEnergyLevel()
		octopuses[y][x+1].increaseEnergyLevel()
		octopuses[y+1][x+1].increaseEnergyLevel()
		octopuses[y+1][x].increaseEnergyLevel()
	} else if y > 0 && y < len(octopuses)-1 && x == len(octopuses[y])-1 {
		octopuses[y-1][x].increaseEnergyLevel()
		octopuses[y+1][x].increaseEnergyLevel()
		octopuses[y+1][x-1].increaseEnergyLevel()
		octopuses[y][x-1].increaseEnergyLevel()
		octopuses[y-1][x-1].increaseEnergyLevel()
	} else if y == 0 && x == 0 {
		octopuses[y][x+1].increaseEnergyLevel()
		octopuses[y+1][x+1].increaseEnergyLevel()
		octopuses[y+1][x].increaseEnergyLevel()
	} else if y == len(octopuses)-1 && x == len(octopuses[y])-1 {
		octopuses[y-1][x].increaseEnergyLevel()
		octopuses[y][x-1].increaseEnergyLevel()
		octopuses[y-1][x-1].increaseEnergyLevel()
	} else if y == 0 && x == len(octopuses[y])-1 {
		octopuses[y+1][x].increaseEnergyLevel()
		octopuses[y+1][x-1].increaseEnergyLevel()
		octopuses[y][x-1].increaseEnergyLevel()
	} else if y == len(octopuses)-1 && x == 0 {
		octopuses[y-1][x].increaseEnergyLevel()
		octopuses[y-1][x+1].increaseEnergyLevel()
		octopuses[y][x+1].increaseEnergyLevel()
	}
	octopuses[y][x].flash()
}

func NumberOfFlashes(filename string) int {
	octopuses := readData(filename)
	sum := 0
	for i := 0; i < 100; i++ {
		for y, _ := range octopuses {
			for x, _ := range octopuses[y] {
				octopuses[y][x].increaseEnergyLevel()
			}
		}

		oldSum := -1
		for oldSum != sum {
			oldSum = sum
			for y, _ := range octopuses {
				for x, _ := range octopuses[y] {
					if octopuses[y][x].energyLevel > 9 {
						sum++
						flash(octopuses, y, x)
					}
				}
			}
		}

		for y, _ := range octopuses {
			for x, _ := range octopuses[y] {
				octopuses[y][x].flashed = false
			}
		}
	}
	return sum
}

func allFlashed(octopuses [][]octopus) bool {
	for _, row := range octopuses {
		for _, octopus := range row {
			if !octopus.flashed {
				return false
			}
		}
	}
	return true
}

func StepWhenAllFlash(filename string) int {
	octopuses := readData(filename)
	sum := 0
	for i := 0; ; i++ {
		for y, _ := range octopuses {
			for x, _ := range octopuses[y] {
				octopuses[y][x].increaseEnergyLevel()
			}
		}

		oldSum := -1
		for oldSum != sum {
			oldSum = sum
			for y, _ := range octopuses {
				for x, _ := range octopuses[y] {
					if octopuses[y][x].energyLevel > 9 {
						sum++
						flash(octopuses, y, x)
					}
				}
			}
		}

		if allFlashed(octopuses) {
			return i + 1
		} else {
			for y, _ := range octopuses {
				for x, _ := range octopuses[y] {
					octopuses[y][x].flashed = false
				}
			}
		}
	}
}
