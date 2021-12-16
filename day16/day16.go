package day16

import (
	"bufio"
	"math"
	"os"

	"github.com/gossie/bitset"
)

type packet struct {
	version    int
	subPackets []*packet
	operator   int
	literal    int
}

func (p *packet) summedVersions() int {
	sum := p.version
	for _, sp := range p.subPackets {
		sum += sp.summedVersions()
	}
	return sum
}

func (p *packet) calculate() int {
	if p.literal >= 0 {
		return p.literal
	}

	switch p.operator {
	case 0:
		sum := 0
		for _, sp := range p.subPackets {
			sum += sp.calculate()
		}
		return sum
	case 1:
		product := 1
		for _, sp := range p.subPackets {
			product *= sp.calculate()
		}
		return product
	case 2:
		min := math.MaxInt
		for _, sp := range p.subPackets {
			min = int(math.Min(float64(min), float64(sp.calculate())))
		}
		return min
	case 3:
		max := math.MinInt
		for _, sp := range p.subPackets {
			max = int(math.Max(float64(max), float64(sp.calculate())))
		}
		return max
	case 5:
		if p.subPackets[0].calculate() > p.subPackets[1].calculate() {
			return 1
		}
		return 0
	case 6:
		if p.subPackets[0].calculate() < p.subPackets[1].calculate() {
			return 1
		}
		return 0
	case 7:
		if p.subPackets[0].calculate() == p.subPackets[1].calculate() {
			return 1
		}
		return 0
	default:
		panic("unknown operator")
	}

}

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

func version(bits *bitset.BitSet, startIndex uint) int {
	version := 0
	for i := startIndex; i < startIndex+3; i++ {
		if bits.IsSet(i) {
			version |= int(math.Pow(2.0, float64(2-(i-startIndex))))
		}
	}
	return version
}

func typeId(bits *bitset.BitSet, startIndex uint) int {
	version := 0
	for i := startIndex; i < startIndex+3; i++ {
		if bits.IsSet(i) {
			version |= int(math.Pow(2.0, float64(2-(i-startIndex))))
		}
	}
	return version
}

func parseLiteral(bits *bitset.BitSet, startIndex uint, masks []int) (int, uint) {
	literal := 0
	index := startIndex
	goOn := true
	for goOn {
		goOn = bits.IsSet(index)
		literal <<= 4
		index++
		for i := index; i < index+4; i++ {
			if bits.IsSet(i) {
				literal |= masks[i-index]
			}
		}
		index += 4
	}
	return literal, index
}

func lengthTypeId(bits *bitset.BitSet, startIndex uint) int {
	if bits.IsSet(startIndex) {
		return 1
	}
	return 0
}

func readBits(bits *bitset.BitSet, startIndex, numberOfBits uint) int {
	result := 0
	for i := startIndex; i < startIndex+numberOfBits; i++ {
		if bits.IsSet(i) {
			exponent := (numberOfBits - 1) - (i - startIndex)
			mask := int(math.Pow(2.0, float64(exponent)))
			result |= mask
		}
	}
	return result
}

func parsePacket(bits *bitset.BitSet, index uint, masks []int) (*packet, uint) {
	version := version(bits, index)
	index += 3
	// fmt.Println("version", version)
	typeId := typeId(bits, index)
	index += 3
	// fmt.Println("typeId", typeId)
	subPackets := make([]*packet, 0)
	literal := -1
	operator := -1
	if typeId == 4 {
		literal, index = parseLiteral(bits, index, masks)
		// fmt.Println("literal", literal)
	} else {
		operator = typeId
		lengthTypeId := lengthTypeId(bits, index)
		index++
		if lengthTypeId == 0 {
			subPacketLength := readBits(bits, index, 15)
			// fmt.Println("subPacketLength", subPacketLength)
			index += 15
			for subPacketLength > 0 {
				subPacket, newIndex := parsePacket(bits, index, masks)
				subPackets = append(subPackets, subPacket)
				subPacketLength -= int(newIndex) - int(index)
				index = newIndex
			}
		} else {
			numberOfSubPackets := readBits(bits, index, 11)
			// fmt.Println("numberOfSubPackets", numberOfSubPackets)
			index += 11
			for i := 0; i < numberOfSubPackets; i++ {
				subPacket, newIndex := parsePacket(bits, index, masks)
				subPackets = append(subPackets, subPacket)
				index = newIndex
			}
		}
	}

	return &packet{version: version, subPackets: subPackets, literal: literal, operator: operator}, index
}

func Versions(filename string) int {
	bits := readData(filename)

	masks := []int{8, 4, 2, 1}
	index := uint(0)

	rootPacket, _ := parsePacket(bits, index, masks)

	return rootPacket.summedVersions()
}

func Calculate(filename string) int {
	bits := readData(filename)

	masks := []int{8, 4, 2, 1}
	index := uint(0)

	rootPacket, _ := parsePacket(bits, index, masks)

	return rootPacket.calculate()
}
