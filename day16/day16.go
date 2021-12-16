package day16

import (
	"bufio"
	"math"
	"os"

	"github.com/gossie/bitset"
)

func readData(filename string) *bitset.BitSet {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	masks := []int{8, 4, 2, 1}
	mapping := map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15}
	bits := bitset.BitSet{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	for i, hex := range scanner.Text() {
		dec := mapping[hex]
		for j, mask := range masks {
			if mask&dec != 0 {
				bits.Set(uint(i*4 + j))
			}
		}
	}

	return &bits
}

func version(bits *bitset.BitSet, startIndex int) int {
	version := 0
	for i := uint(startIndex); i < uint(startIndex+3); i++ {
		if bits.IsSet(i) {
			version |= int(math.Pow(2.0, float64(2-(i-uint(startIndex)))))
		}
	}
	return version
}

func typeId(bits *bitset.BitSet, startIndex int) int {
	version := 0
	for i := uint(startIndex); i < uint(startIndex+3); i++ {
		if bits.IsSet(i) {
			version |= int(math.Pow(2.0, float64(2-(i-uint(startIndex)))))
		}
	}
	return version
}

func Headers(filename string) int {
	bits := readData(filename)
	version := version(bits, 0)
	typeId := typeId(bits, 3)
	if typeId == 4 {

	}

	return version
}
