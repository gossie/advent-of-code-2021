package day20

import (
	"bufio"
	"os"
)

func readData(filename string) ([][]rune, string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	inputImage := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	enhancement := scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			row := make([]rune, 0)
			for _, r := range line {
				row = append(row, r)
			}
			inputImage = append(inputImage, row)
		}
	}

	return inputImage, enhancement
}

func determineEnhancementIndex(inputImage [][]rune, y, x int) int {
	index := 0
	if y-1 >= 0 && y-1 < len(inputImage) && x-1 >= 0 && x-1 < len(inputImage[0]) && inputImage[y-1][x-1] == '#' {
		index |= 0b100
	}
	if y-1 > 0 && y-1 < len(inputImage) && x > 0 && x < len(inputImage[0]) && inputImage[y-1][x] == '#' {
		index |= 0b010
	}
	if y-1 > 0 && y-1 < len(inputImage) && x+1 > 0 && x+1 < len(inputImage[0]) && inputImage[y-1][x+1] == '#' {
		index |= 0b001
	}
	index <<= 3

	if y > 0 && y < len(inputImage) && x-1 > 0 && x-1 < len(inputImage[0]) && inputImage[y][x-1] == '#' {
		index |= 0b100
	}
	if y > 0 && y < len(inputImage) && x > 0 && x < len(inputImage[0]) && inputImage[y][x] == '#' {
		index |= 0b010
	}
	if y > 0 && y < len(inputImage) && x+1 > 0 && x+1 < len(inputImage[0]) && inputImage[y][x+1] == '#' {
		index |= 0b001
	}
	index <<= 3

	if y+1 > 0 && y+1 < len(inputImage) && x-1 > 0 && x-1 < len(inputImage[0]) && inputImage[y+1][x-1] == '#' {
		index |= 0b100
	}
	if y+1 > 0 && y+1 < len(inputImage) && x > 0 && x < len(inputImage[0]) && inputImage[y+1][x] == '#' {
		index |= 0b010
	}
	if y+1 > 0 && y+1 < len(inputImage) && x+1 > 0 && x+1 < len(inputImage[0]) && inputImage[y+1][x+1] == '#' {
		index |= 0b001
	}

	return index
}

func extendInputImage(inputImage [][]rune, toAdd rune) [][]rune {
	for i := 0; i < len(inputImage); i++ {
		inputImage[i] = append(inputImage[i], []rune{toAdd, toAdd, toAdd, toAdd, toAdd}...)
		inputImage[i] = append([]rune{toAdd, toAdd, toAdd, toAdd, toAdd}, inputImage[i]...)
	}
	firstRow := make([]rune, 0, len(inputImage[0]))
	lastRow := make([]rune, 0, len(inputImage[0]))
	for i := 0; i < len(inputImage[0]); i++ {
		firstRow = append(firstRow, toAdd)
		lastRow = append(lastRow, toAdd)
	}

	inputImage = append(inputImage, [][]rune{lastRow, lastRow, lastRow, lastRow, lastRow}...)
	inputImage = append([][]rune{firstRow, firstRow, firstRow, firstRow, firstRow}, inputImage...)

	return inputImage
}

func cutOutputImage(outputImage [][]rune) [][]rune {
	outputImage = outputImage[4 : len(outputImage)-4]
	for i := 0; i < len(outputImage); i++ {
		outputImage[i] = outputImage[i][4 : len(outputImage[i])-4]
	}
	return outputImage
}

func NumberOfLitPixels(filename string, iterations int) int {
	inputImage, enhancement := readData(filename)
	var output [][]rune

	for i := 1; i <= iterations; i++ {
		toAdd := '.'
		if i%2 == 0 {
			toAdd = '#'
		}
		inputImage = extendInputImage(inputImage, toAdd)
		output = make([][]rune, 0, len(inputImage))

		for y := 0; y < len(inputImage); y++ {
			outputRow := make([]rune, 0, len(inputImage[0]))
			for x := 0; x < len(inputImage[0]); x++ {
				enhancementIndex := determineEnhancementIndex(inputImage, y, x)
				outputRow = append(outputRow, rune(enhancement[enhancementIndex]))
			}
			output = append(output, outputRow)
		}

		output = cutOutputImage(output)

		inputImage = output
	}

	litPixels := 0
	for _, row := range output {
		for _, pixel := range row {
			if pixel == '#' {
				litPixels++
			}
		}
	}

	return litPixels
}
