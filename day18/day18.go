package day18

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

type node struct {
	parent *node
	left   *node
	right  *node
	value  int
}

func (n *node) leaf() bool {
	return n.left == nil && n.right == nil
}

func (n *node) add(other *node) *node {
	newNode := node{left: n, right: other}
	newNode.left.parent = &newNode
	newNode.right.parent = &newNode
	return &newNode
}

func (n *node) explode() {
	if (n.parent.left) == n {
		// I am the left node
		current := n.parent.right
		for current.left != nil {
			current = current.left
		}
		current.value += n.right.value

		current = n
		parentNode := n.parent
		for parentNode != nil && parentNode.left == current {
			current = parentNode
			parentNode = parentNode.parent
		}
		if parentNode != nil {
			current = parentNode.left
			for current.right != nil {
				current = current.right
			}
			current.value += n.left.value
		}
	} else {
		// I am the right node
		current := n.parent.left
		for current.right != nil {
			current = current.right
		}
		current.value += n.left.value

		current = n
		parentNode := n.parent
		for parentNode != nil && parentNode.right == current {
			current = parentNode
			parentNode = parentNode.parent
		}
		if parentNode != nil {
			current = parentNode.right
			for current.left != nil {
				current = current.left
			}
			current.value += n.right.value
		}
	}
}

func (n *node) reduceCompletey() *node {
	changed := true
	current := n
	for changed {
		changed, current = current.reduce(1, true)
		for changed {
			changed, current = current.reduce(1, true)
		}

		changed, current = current.reduce(1, false)
	}

	return current
}

func (n *node) reduce(depth int, explode bool) (bool, *node) {
	if explode {
		if !n.leaf() && depth >= 5 {
			n.explode()
			return true, &node{value: 0, parent: n.parent}
		} else {
			if n.left != nil {
				leftChanged, newLeft := n.left.reduce(depth+1, explode)
				if leftChanged {
					n.left = newLeft
					return true, n
				}
			}

			if n.right != nil {
				rightChanged, newRight := n.right.reduce(depth+1, explode)
				if rightChanged {
					n.right = newRight
					return true, n
				}
			}

			return false, n
		}
	} else {
		if n.leaf() {
			if n.value >= 10 {
				left := &node{value: int(math.Floor(float64(n.value) / 2.0))}
				right := &node{value: int(math.Ceil(float64(n.value) / 2.0))}
				newNode := node{left: left, right: right, parent: n.parent}
				left.parent = &newNode
				right.parent = &newNode
				return true, &newNode
			}
			return false, n
		} else {
			if n.left != nil {
				leftChanged, newLeft := n.left.reduce(depth+1, explode)
				if leftChanged {
					n.left = newLeft
					return true, n
				}
			}

			if n.right != nil {
				rightChanged, newRight := n.right.reduce(depth+1, explode)
				if rightChanged {
					n.right = newRight
					return true, n
				}
			}

			return false, n
		}
	}
}

func (n *node) magnitude() int {
	if n.leaf() {
		return n.value
	}
	return 3*(n.left).magnitude() + 2*(n.right).magnitude()
}

func (n *node) asString() string {
	if n.leaf() {
		return strconv.Itoa(n.value)
	}
	return "[" + n.left.asString() + "," + n.right.asString() + "]"
}

func indexOfSeperatingComma(number string) int {
	numberOfOpensBeforeClose := 0
	for i, c := range number {
		if c == ']' {
			numberOfOpensBeforeClose--
			if numberOfOpensBeforeClose == 0 {
				return i + 1
			}
		}
		if c == '[' {
			numberOfOpensBeforeClose++
		}
	}
	panic("no comma")
}

func parsePair(number string) *node {
	index := 1
	if number[0] == '[' {
		index = indexOfSeperatingComma(number)
	}
	left := parseNumber(number[0:index])
	var right *node
	if number[index+1] == '[' {
		right = parseNumber(number[index+1:])
	} else {
		right = parseNumber(number[index+1:])
	}
	newNode := node{left: left, right: right}
	newNode.left.parent = &newNode
	newNode.right.parent = &newNode
	return &newNode
}

func parseNumber(line string) *node {
	if line[0] == '[' {
		return parsePair(line[1 : len(line)-1])
	} else {
		n, _ := strconv.Atoi(line)
		return &node{value: n}
	}
}

func readData(filename string) []*node {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	numbers := []*node{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		numbers = append(numbers, parseNumber(scanner.Text()))
	}

	return numbers
}

func Magnitude(filename string) int {
	numbers := readData(filename)
	current := numbers[0]

	for i := 1; i < len(numbers); i++ {
		current = current.add(numbers[i])
		current = current.reduceCompletey()
	}
	return current.magnitude()
}

func LargestMagnitude(filename string) int {
	numbers := readData(filename)

	max := math.MinInt
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				numbers = readData(filename)
				magnitude := numbers[i].add(numbers[j]).reduceCompletey().magnitude()
				max = int(math.Max(float64(magnitude), float64(max)))

				numbers = readData(filename)
				magnitude = numbers[j].add(numbers[i]).reduceCompletey().magnitude()
				max = int(math.Max(float64(magnitude), float64(max)))
			}
		}
	}
	return max
}
