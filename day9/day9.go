package day9

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/gossie/adventofcode2021/util"
)

type point struct {
	x int
	y int
}

func readData(filename string) [][]int {
	heights := make([][]int, 0)

	scanner := bufio.NewScanner(util.LoadFile(filename))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := make([]int, 0)
		text := scanner.Text()
		for i, _ := range text {
			l, _ := strconv.Atoi(text[i : i+1])
			row = append(row, l)
		}
		heights = append(heights, row)
	}
	return heights
}

func isLowPoint(data [][]int, y int, x int) bool {
	height := data[y][x]
	if y > 0 && y < len(data)-1 && x > 0 && x < len(data[y])-1 {
		return height < data[y-1][x] && height < data[y+1][x] && height < data[y][x-1] && height < data[y][x+1]
	} else if y == 0 && x > 0 && x < len(data[y])-1 {
		return height < data[y+1][x] && height < data[y][x-1] && height < data[y][x+1]
	} else if y == len(data)-1 && x > 0 && x < len(data[y])-1 {
		return height < data[y-1][x] && height < data[y][x-1] && height < data[y][x+1]
	} else if y > 0 && y < len(data)-1 && x == 0 {
		return height < data[y-1][x] && height < data[y+1][x] && height < data[y][x+1]
	} else if y > 0 && y < len(data)-1 && x == len(data[y])-1 {
		return height < data[y-1][x] && height < data[y+1][x] && height < data[y][x-1]
	} else if y == 0 && x == 0 {
		return height < data[y+1][x] && height < data[y][x+1]
	} else if y == len(data)-1 && x == len(data[y])-1 {
		return height < data[y-1][x] && height < data[y][x-1]
	} else if y == 0 && x == len(data[y])-1 {
		return height < data[y+1][x] && height < data[y][x-1]
	} else if y == len(data)-1 && x == 0 {
		return height < data[y-1][x] && height < data[y][x+1]
	}
	return false
}

func RiskLevel(filename string) int {
	sum := 0
	data := readData(filename)
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if isLowPoint(data, y, x) {
				sum += data[y][x] + 1
			}
		}
	}
	return sum
}

func reachedBorderOfBasin(data [][]int, y, x, currentHeight int) bool {
	return y < 0 || y >= len(data) || x < 0 || x >= len(data[y]) || data[y][x] <= currentHeight || data[y][x] == 9
}

func basinsSize(data [][]int, y, x, currentHeight int, visitedPoints map[point]bool) int {
	if _, visited := visitedPoints[point{x: x, y: y}]; visited || reachedBorderOfBasin(data, y, x, currentHeight) {
		return 0
	}
	visitedPoints[point{x: x, y: y}] = true
	return 1 + basinsSize(data, y-1, x, data[y][x], visitedPoints) + basinsSize(data, y+1, x, data[y][x], visitedPoints) + basinsSize(data, y, x-1, data[y][x], visitedPoints) + basinsSize(data, y, x+1, data[y][x], visitedPoints)
}

func Basin(filename string) int {
	basins := make([]int, 0)
	data := readData(filename)
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if isLowPoint(data, y, x) {
				visitedPoints := map[point]bool{{x: x, y: y}: true}
				size := 1 + basinsSize(data, y-1, x, data[y][x], visitedPoints) + basinsSize(data, y+1, x, data[y][x], visitedPoints) + basinsSize(data, y, x-1, data[y][x], visitedPoints) + basinsSize(data, y, x+1, data[y][x], visitedPoints)
				basins = append(basins, size)
			}
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]
}
