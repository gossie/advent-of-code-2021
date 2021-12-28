package day25

import (
	"bufio"
	"image"
	"image/color"
	"image/gif"
	"os"
)

type point struct {
	x, y int
}

type cucumber struct {
	view        string
	position    point
	markForMove bool
}

func (c *cucumber) move(direction string, bottom [][]*cucumber) {
	bottom[c.position.y][c.position.x] = nil
	if direction == "east" {
		c.position.x++
		if c.position.x == len(bottom[c.position.y]) {
			c.position.x = 0
		}
		bottom[c.position.y][c.position.x] = c
	} else {
		c.position.y++
		if c.position.y == len(bottom) {
			c.position.y = 0
		}
		bottom[c.position.y][c.position.x] = c
	}

	c.markForMove = false
}

type state struct {
	easternCucumbers  []*cucumber
	southernCucumbers []*cucumber
	bottom            [][]*cucumber
}

func readData(filename string) *state {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	easternCucumbers := make([]*cucumber, 0)
	southernCucumbers := make([]*cucumber, 0)
	bottom := make([][]*cucumber, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := make([]*cucumber, 0)
		for _, r := range scanner.Text() {
			if r == '.' {
				row = append(row, nil)
			} else if r == '>' {
				c := cucumber{position: point{x: len(row), y: len(bottom)}, view: ">"}
				row = append(row, &c)
				easternCucumbers = append(easternCucumbers, &c)
			} else if r == 'v' {
				c := cucumber{position: point{x: len(row), y: len(bottom)}, view: "v"}
				row = append(row, &c)
				southernCucumbers = append(southernCucumbers, &c)
			}
		}
		bottom = append(bottom, row)
	}
	return &state{bottom: bottom, easternCucumbers: easternCucumbers, southernCucumbers: southernCucumbers}
}

func eastShiftPossible(c *cucumber, bottom [][]*cucumber) bool {
	if c.position.x == len(bottom[c.position.y])-1 {
		return bottom[c.position.y][0] == nil
	}
	return bottom[c.position.y][c.position.x+1] == nil
}

func southShiftPossible(c *cucumber, bottom [][]*cucumber) bool {
	if c.position.y == len(bottom)-1 {
		return bottom[0][c.position.x] == nil
	}
	return bottom[c.position.y+1][c.position.x] == nil
}

func performStep(state *state) bool {
	change := false
	for _, c := range state.easternCucumbers {
		if eastShiftPossible(c, state.bottom) {
			c.markForMove = true
			change = true
		}
	}

	for _, c := range state.easternCucumbers {
		if c.markForMove {
			c.move("east", state.bottom)
		}
	}

	for _, c := range state.southernCucumbers {
		if southShiftPossible(c, state.bottom) {
			c.markForMove = true
			change = true
		}
	}

	for _, c := range state.southernCucumbers {
		if c.markForMove {
			c.move("south", state.bottom)
		}
	}

	return change
}

func WhichStep(filename string) int {
	state := readData(filename)

	var images []*image.Paletted
	var delays []int

	goOn := true
	step := 0
	for ; goOn; step++ {
		goOn = performStep(state)
		images = append(images, renderCucumbers(state))
		delays = append(delays, 0)
	}

	createGif("day25/cucumbers.gif", images, delays)

	return step
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

func renderCucumbers(state *state) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{0xff, 0xff, 0xff, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0xff, 0xff},
	}
	img := image.NewPaletted(image.Rect(0, 0, 136, 139), palette)

	for _, c := range state.easternCucumbers {
		img.Set(c.position.x, c.position.y, color.RGBA{0x00, 0x00, 0xff, 0xff})
	}

	for _, c := range state.southernCucumbers {
		img.Set(c.position.x, c.position.y, color.RGBA{0xff, 0x00, 0xff, 0xff})
	}

	return img
}
