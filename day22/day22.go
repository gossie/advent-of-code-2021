package day22

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type cube struct {
	x, y, z coordinateRange
	factor  int
}

func (c *cube) intersects(other *cube) *cube {
	x := coordinateRange{int(math.Max(float64(c.x.min), float64(other.x.min))), int(math.Min(float64(c.x.max), float64(other.x.max)))}
	y := coordinateRange{int(math.Max(float64(c.y.min), float64(other.y.min))), int(math.Min(float64(c.y.max), float64(other.y.max)))}
	z := coordinateRange{int(math.Max(float64(c.z.min), float64(other.z.min))), int(math.Min(float64(c.z.max), float64(other.z.max)))}
	intersection := cube{x: x, y: y, z: z}
	if (intersection.x.min > intersection.x.max) || (intersection.y.min > intersection.y.max) || (intersection.z.min > intersection.z.max) {
		return nil
	} else {
		return &intersection
	}
}

func (c *cube) volume() uint64 {
	length := uint64(c.x.max - c.x.min + 1)
	width := uint64(c.y.max - c.y.min + 1)
	height := uint64(c.z.max - c.z.min + 1)
	return length * width * height
}

type coordinateRange struct {
	min, max int
}

type instruction struct {
	action  string
	x, y, z coordinateRange
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

func NumberOfEnabledCubes(filename string, limit bool) uint64 {
	instructions := readData(filename)

	cubes := make([]*cube, 0)

	for index := range instructions {
		instruction := instructions[index]
		if !limit || instruction.x.min >= -50 && instruction.x.max <= 50 && instruction.y.min >= -50 && instruction.y.max <= 50 && instruction.z.min >= -50 && instruction.z.max <= 50 {
			cubesToAdd := make([]*cube, 0)
			newCube := cube{x: instruction.x, y: instruction.y, z: instruction.z}
			if instruction.action == "on" {
				newCube.factor = 1
				cubesToAdd = append(cubesToAdd, &newCube)
			}

			for _, c := range cubes {
				intersection := c.intersects(&newCube)
				if intersection != nil {
					intersection.factor = -c.factor
					cubesToAdd = append(cubesToAdd, intersection)
				}
			}
			cubes = append(cubes, cubesToAdd...)
		}
	}

	count := uint64(0)
	for _, value := range cubes {
		count += uint64(value.factor) * value.volume()
	}
	return count
}
