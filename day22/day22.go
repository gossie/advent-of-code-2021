package day22

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type coordinateRange struct {
	min, max int
}

type instruction struct {
	action string
	x      coordinateRange
	y      coordinateRange
	z      coordinateRange
}

func parseCoordinateRange(cr string) (rune, coordinateRange) {
	axis := cr[0]
	minMax := strings.Split(cr[2:], "..")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return rune(axis), coordinateRange{min: min, max: max}
}

func readData(filename string) []instruction {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	instructions := make([]instruction, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		c := strings.Split(line, " ")
		coordinateRanges := strings.Split(c[1], ",")
		var x coordinateRange
		var y coordinateRange
		var z coordinateRange
		for _, cr := range coordinateRanges {
			axis, cRange := parseCoordinateRange(cr)
			switch axis {
			case 'x':
				x = cRange
			case 'y':
				y = cRange
			case 'z':
				z = cRange
			default:
				panic("unknown axis")
			}
		}

		instructions = append(instructions, instruction{action: c[0], x: x, y: y, z: z})
	}

	return instructions
}

func NumberOfEnabledCubes(filename string) int {
	instructions := readData(filename)

	cubes := make(map[point]bool)

	for index := range instructions {
		instruction := instructions[index]
		if instruction.x.min >= -50 && instruction.x.max <= 50 && instruction.y.min >= -50 && instruction.y.max <= 50 && instruction.z.min >= -50 && instruction.z.max <= 50 {
			for x := instruction.x.min; x <= instruction.x.max; x++ {
				for y := instruction.y.min; y <= instruction.y.max; y++ {
					for z := instruction.z.min; z <= instruction.z.max; z++ {
						p := point{x: x, y: y, z: z}
						if instruction.action == "on" {
							cubes[p] = true
						} else {
							cubes[p] = false
						}
					}
				}
			}
		}
	}

	count := 0
	for _, value := range cubes {
		if value {
			count++
		}
	}
	return count
}
