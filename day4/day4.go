package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type field struct {
	number int
	marked bool
}

func (f *field) mark() {
	f.marked = true
}

type bingoSheet struct {
	fields [][]field
}

func (sheet *bingoSheet) selectNumber(number int) {
	for i, row := range sheet.fields {
		for j, field := range row {
			if field.number == number {
				sheet.fields[i][j].mark()
			}
		}
	}
}

func (sheet *bingoSheet) won() bool {
	for _, row := range sheet.fields {
		for fieldIndex, field := range row {
			if !field.marked {
				break
			}
			if fieldIndex == len(row)-1 {
				return true
			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			field := sheet.fields[j][i]
			if !field.marked {
				break
			}
			if j == 4 {
				return true
			}
		}
	}

	return false
}

func (sheet *bingoSheet) points(number int) int {
	sum := 0
	for _, row := range sheet.fields {
		for _, field := range row {
			if !field.marked {
				sum += field.number
			}
		}
	}
	return number * sum
}

func readBingoData(filename string) ([]int, []bingoSheet) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var sheets []bingoSheet
	var polledNumbers []int
	currentSheet := bingoSheet{}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if i == 0 {
			for _, number := range strings.Split(line, ",") {
				if parsed, err := strconv.Atoi(number); err == nil {
					polledNumbers = append(polledNumbers, parsed)
				}
			}
		} else if line == "" {
			if i > 1 {
				sheets = append(sheets, currentSheet)
				currentSheet = bingoSheet{}
			}
		} else {
			var fields []field
			for _, number := range strings.Split(line, " ") {
				if parsed, err := strconv.Atoi(number); err == nil {
					fields = append(fields, field{number: parsed, marked: false})
				}
			}
			currentSheet.fields = append(currentSheet.fields, fields)
		}
	}
	sheets = append(sheets, currentSheet)
	return polledNumbers, sheets
}

func BingoFirstWin(file string) int {
	numbers, sheets := readBingoData(file)

	for _, number := range numbers {
		for _, sheet := range sheets {
			sheet.selectNumber(number)
			if sheet.won() {
				return sheet.points(number)
			}
		}
	}
	panic("no winner")
}

func BingoLastWin(file string) int {
	numbers, sheets := readBingoData(file)

	var lastWinner bingoSheet
	var lastNumber int

	for _, number := range numbers {
		for _, sheet := range sheets {
			if !sheet.won() {
				sheet.selectNumber(number)
				if sheet.won() {
					lastWinner = sheet
					lastNumber = number
				}
			}
		}
	}
	return lastWinner.points(lastNumber)
}
