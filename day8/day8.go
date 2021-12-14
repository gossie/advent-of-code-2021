package day8

import (
	"bufio"
	"math"
	"strings"

	"github.com/gossie/adventofcode2021/util"
)

type line struct {
	input  []string
	output []string
}

func readData(filename string) []line {
	lines := make([]line, 0)

	scanner := bufio.NewScanner(util.LoadFile(filename))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lineComponents := strings.Split(scanner.Text(), " | ")
		lines = append(lines, line{input: strings.Split(lineComponents[0], " "), output: strings.Split(lineComponents[1], " ")})
	}
	return lines
}

func NumberOf1And4And7And8(filename string) int {
	sum := 0
	for _, l := range readData(filename) {
		for _, out := range l.output {
			length := len(out)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				sum++
			}
		}
	}
	return sum
}

func initialCodeNumberMappings(input []string) map[int]string {
	numberCodeMapping := make(map[int]string)
	for _, code := range input {
		switch len(code) {
		case 2:
			numberCodeMapping[1] = code
		case 3:
			numberCodeMapping[7] = code
		case 4:
			numberCodeMapping[4] = code
		case 7:
			numberCodeMapping[8] = code
		}
	}
	return numberCodeMapping
}

func containsAll(str string, chars string) bool {
	for _, letter := range chars {
		if !strings.ContainsRune(str, letter) {
			return false
		}
	}
	return true
}

func OutputSum(filename string) int {
	sum := 0
	for _, l := range readData(filename) {
		numberCodeMapping := initialCodeNumberMappings(l.input)
		for len(numberCodeMapping) < 10 {
			for _, code := range l.input {
				if len(code) == 6 {
					// it is the 0, the 6 or the 9
					if containsAll(code, numberCodeMapping[1]) {
						if containsAll(code, numberCodeMapping[4]) {
							numberCodeMapping[9] = code
						} else {
							numberCodeMapping[0] = code
						}
					} else {
						numberCodeMapping[6] = code
					}
				}

				if len(code) == 5 {
					if containsAll(code, numberCodeMapping[1]) {
						numberCodeMapping[3] = code
					} else {
						if six, present := numberCodeMapping[6]; present {
							var sect rune
							one := numberCodeMapping[1]
							for _, letter := range one {
								if strings.ContainsRune(six, letter) {
									sect = letter
									break
								}
							}
							if strings.ContainsRune(code, sect) {
								numberCodeMapping[5] = code
							} else {
								numberCodeMapping[2] = code
							}
						}
					}
				}
			}
		}

		for index, out := range l.output {
			for number, code := range numberCodeMapping {
				if len(code) == len(out) && containsAll(out, code) {
					sum += number * int(math.Pow10(3-index))
					break
				}
			}
		}

	}
	return sum
}
