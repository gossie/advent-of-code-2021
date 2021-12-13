package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type foldInstruction struct {
	axis string
	line int
}

func readData(filename string) ([][]string, []foldInstruction) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	points := make([]point, 0)
	foldInstructions := make([]foldInstruction, 0)

	maxX := 0
	maxY := 0

	scanCoordinates := true
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanCoordinates = false
			continue
		}
		if scanCoordinates {
			coordinates := strings.Split(line, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])
			maxX = int(math.Max(float64(x), float64(maxX)))
			maxY = int(math.Max(float64(y), float64(maxY)))
			points = append(points, point{x: x, y: y})
		} else {
			instruction := strings.Split(line[11:], "=")
			foldLine, _ := strconv.Atoi(instruction[1])
			foldInstructions = append(foldInstructions, foldInstruction{axis: instruction[0], line: foldLine})
		}
	}

	sheet := make([][]string, 0, maxY+1)
	for y := 0; y <= maxY; y++ {
		row := make([]string, 0, maxX+1)
		for x := 0; x <= maxX; x++ {
			row = append(row, ".")
		}
		sheet = append(sheet, row)
	}

	for _, point := range points {
		sheet[point.y][point.x] = "#"
	}

	return sheet, foldInstructions
}

func AfterOneFold(filename string) int {
	sheet, foldInstructions := readData(filename)

	if foldInstructions[0].axis == "x" {
		sheet = foldHorizontally(sheet, foldInstructions[0].line)
	} else {
		sheet = foldVertically(sheet, foldInstructions[0].line)
	}

	sum := 0
	for _, row := range sheet {
		for _, dot := range row {
			if dot == "#" {
				sum++
			}
		}
	}
	return sum
}

func foldHorizontally(sheet [][]string, foldLine int) [][]string {
	newSheet := make([][]string, 0)
	for y := 0; y < len(sheet); y++ {
		row := make([]string, 0)
		longerSlice := sheet[y][0:foldLine]
		shorterSlice := sheet[y][foldLine+1:]
		if len(longerSlice) < len(shorterSlice) {
			tmp := longerSlice
			longerSlice = shorterSlice
			shorterSlice = tmp
		}

		for x := 0; x < len(longerSlice)-len(shorterSlice); x++ {
			row = append(row, longerSlice[x])
		}

		for x := len(longerSlice) - len(shorterSlice); x < len(longerSlice); x++ {
			if longerSlice[x] == "#" || shorterSlice[(len(shorterSlice)-1)-(x-(len(longerSlice)-len(shorterSlice)))] == "#" {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}

		newSheet = append(newSheet, row)
	}
	return newSheet
}

func foldVertically(sheet [][]string, foldLine int) [][]string {
	newSheet := make([][]string, 0)
	longerSlice := sheet[0:foldLine]
	shorterSlice := sheet[foldLine+1:]
	if len(longerSlice) < len(shorterSlice) {
		tmp := longerSlice
		longerSlice = shorterSlice
		shorterSlice = tmp
	}

	for y := 0; y < len(longerSlice)-len(shorterSlice); y++ {
		newSheet = append(newSheet, longerSlice[y])
	}

	for y := len(longerSlice) - len(shorterSlice); y < len(longerSlice); y++ {
		row := make([]string, 0)
		for x := 0; x < len(longerSlice[y]); x++ {
			if longerSlice[y][x] == "#" || shorterSlice[(len(shorterSlice)-1)-(y-(len(longerSlice)-len(shorterSlice)))][x] == "#" {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}
		newSheet = append(newSheet, row)
	}
	return newSheet
}

func Code(filename string) {
	sheet, foldInstructions := readData(filename)

	for _, instruction := range foldInstructions {
		if instruction.axis == "x" {
			sheet = foldHorizontally(sheet, instruction.line)
		} else {
			sheet = foldVertically(sheet, instruction.line)
		}
	}

	for _, row := range sheet {
		for _, dot := range row {
			fmt.Print(dot)
		}
		fmt.Println()
	}
}
