package day16

import (
	"bufio"
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

func Headers(filename string) int {
	readData(filename)
	return 0
}
