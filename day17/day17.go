package day17

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

type area struct {
	startPoint point
	endPoint   point
}

func readData(filename string) area {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.Split(scanner.Text()[13:], ", ")
	x := strings.Split(line[0], "..")
	y := strings.Split(line[1], "..")

	x1, _ := strconv.Atoi(x[0][2:])
	x2, _ := strconv.Atoi(x[1])
	y1, _ := strconv.Atoi(y[0][2:])
	y2, _ := strconv.Atoi(y[1])

	startPoint := point{x: int(math.Min(float64(x1), float64(x2))), y: int(math.Max(float64(y1), float64(y2)))}
	endPoint := point{x: int(math.Max(float64(x1), float64(x2))), y: int(math.Min(float64(y1), float64(y2)))}

	return area{startPoint: startPoint, endPoint: endPoint}
}

func findXVelocity(targetX int) (int, int) {
	for x := 1; ; x++ {
		reachedXPos := 0
		steps := 0
		for ; x-steps > 0; steps++ {
			reachedXPos += (x - steps)
		}
		if reachedXPos >= targetX {
			return x, steps
		}
	}
}

func isInTargetRegion(xPos, yPos int, target area) bool {
	return xPos >= target.startPoint.x && xPos <= target.endPoint.x && yPos <= target.startPoint.y && yPos >= target.endPoint.y
}

func heighestY(xVelocity, yVelocity int, target area) int {
	hitTarget := false
	heighestY := 0
	xPos, yPos := 0, 0
	for i := 0; xPos <= target.endPoint.x && yPos >= target.endPoint.y; i++ {
		xPos += int(math.Max(float64(xVelocity-i), 0.0))
		yPos += (yVelocity - i)
		heighestY = int(math.Max(float64(heighestY), float64(yPos)))
		hitTarget = hitTarget || isInTargetRegion(xPos, yPos, target)
	}
	if hitTarget {
		return heighestY
	}
	return math.MinInt
}

func hitsTarget(xVelocity, yVelocity int, target area) bool {
	heighestY := 0
	xPos, yPos := 0, 0
	for i := 0; xPos <= target.endPoint.x && yPos >= target.endPoint.y; i++ {
		xPos += int(math.Max(float64(xVelocity-i), 0.0))
		yPos += (yVelocity - i)
		heighestY = int(math.Max(float64(heighestY), float64(yPos)))
		if isInTargetRegion(xPos, yPos, target) {
			return true
		}
	}
	return false
}

func Heighest(filename string) int {
	targetArea := readData(filename)
	return int(math.Abs(float64(targetArea.endPoint.y))) * (int(math.Abs(float64(targetArea.endPoint.y))) - 1) / 2
}

func HowMany(filename string) int {
	targetArea := readData(filename)

	count := 0
	for x := 1; x < targetArea.endPoint.x+1; x++ {
		for y := -100; y < 100; y++ {
			if hitsTarget(x, y, targetArea) {
				count++
			}
		}
	}
	return count
}
