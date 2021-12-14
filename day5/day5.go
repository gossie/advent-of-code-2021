package day5

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func parseCoordinates(coordinates []string) point {
	x, err := strconv.Atoi(coordinates[0])
	if err != nil {
		panic("x coordinate cannot be parsed")
	}

	y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		panic("y coordinate cannot be parsed")
	}

	return point{x: x, y: y}
}

func readData(filename string, includeDiagonals bool) ([]line, [][]int) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []line
	maxX := 0
	maxY := 0

	for i := 0; scanner.Scan(); i++ {
		textLine := scanner.Text()
		startAndEnd := strings.Split(textLine, " -> ")
		start := parseCoordinates(strings.Split(startAndEnd[0], ","))
		end := parseCoordinates(strings.Split(startAndEnd[1], ","))
		if includeDiagonals || start.x == end.x || start.y == end.y {
			lines = append(lines, line{start: start, end: end})
			maxX = int(math.Max(float64(maxX), math.Max(float64(start.x), float64(end.x))))
			maxY = int(math.Max(float64(maxY), math.Max(float64(start.y), float64(end.y))))
		}
	}

	resultMap := make([][]int, 0, maxY)
	for i := 0; i < maxY+1; i++ {
		resultMap = append(resultMap, make([]int, maxX+1, maxX+1))
	}

	return lines, resultMap
}

func delta(line line) (int, int) {
	xDelta := 0
	if line.start.x < line.end.x {
		xDelta = 1
	} else if line.start.x > line.end.x {
		xDelta = -1
	}

	yDelta := 0
	if line.start.y < line.end.y {
		yDelta = 1
	} else if line.start.y > line.end.y {
		yDelta = -1
	}

	return xDelta, yDelta
}

func scanLines(lines []line, resultMap [][]int) {
	for _, line := range lines {
		x := line.start.x
		y := line.start.y
		xDelta, yDelta := delta(line)
		for x != line.end.x || y != line.end.y {
			resultMap[y][x]++
			x += xDelta
			y += yDelta
		}
		resultMap[y][x]++
	}
}

func AvoidDangerousArea(filename string, includeDiagonals bool) int {
	lines, resultMap := readData(filename, includeDiagonals)

	scanLines(lines, resultMap)

	result := 0
	for _, row := range resultMap {
		for _, field := range row {
			if field > 1 {
				result++
			}
		}
	}

	return result
}
