package day25

import (
	"bufio"
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
	goOn := true
	step := 0
	for ; goOn; step++ {
		goOn = performStep(state)
	}
	return step
}
