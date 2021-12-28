package day11

import (
	"bufio"
	"image"
	"image/color"
	"image/gif"
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

	var images []*image.Paletted
	var delays []int

	sum := 0
	for i := 0; ; i++ {
		for y, _ := range octopuses {
			for x, _ := range octopuses[y] {
				octopuses[y][x].increaseEnergyLevel()
			}
		}

		images = append(images, renderOctopus(octopuses))
		delays = append(delays, 0)

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

		images = append(images, renderOctopus(octopuses))
		delays = append(delays, 0)

		if allFlashed(octopuses) {
			images = append(images, renderOctopus(octopuses))
			delays = append(delays, 5)
			createGif("day11/octopus.gif", images, delays)
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

func createGif(name string, images []*image.Paletted, delays []int) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic("gif")
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func renderOctopus(octopusses [][]octopus) *image.Paletted {
	palette := []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0xff, 0xff},
	}
	step := 0xff / 10
	for i := 0; i < 10; i++ {
		palette = append(palette, color.RGBA{uint8(step * i), uint8(step * i), uint8(step * i), 0xff})
	}

	img := image.NewPaletted(image.Rect(0, 0, 100, 100), palette)

	for y, row := range octopusses {
		for x, o := range row {
			step := 0xff / 10
			for i := 0; i < 10; i++ {
				for j := 0; j < 10; j++ {
					img.Set(x*10+i, y*10+j, color.RGBA{uint8(step * o.energyLevel), uint8(step * o.energyLevel), uint8(step * o.energyLevel), 0xff})
				}
			}
		}
	}

	return img
}
