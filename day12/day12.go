package day12

import (
	"bufio"
	"os"
	"strings"
)

type cave struct {
	name      string
	neighbors []*cave
}

func (c *cave) addNeighbor(neighbor *cave) {
	c.neighbors = append(c.neighbors, neighbor)
}

func (c *cave) small() bool {
	return c.name == strings.ToLower(c.name)
}

func readData(filename string) *cave {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	allCaves := make(map[string]*cave)
	var start *cave = nil

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tunnel := strings.Split(scanner.Text(), "-")
		if _, present := allCaves[tunnel[0]]; !present {
			allCaves[tunnel[0]] = &cave{name: tunnel[0]}
		}

		if _, present := allCaves[tunnel[1]]; !present {
			allCaves[tunnel[1]] = &cave{name: tunnel[1]}
		}

		allCaves[tunnel[0]].addNeighbor(allCaves[tunnel[1]])
		allCaves[tunnel[1]].addNeighbor(allCaves[tunnel[0]])

		if allCaves[tunnel[0]].name == "start" {
			start = allCaves[tunnel[0]]
		}

		if allCaves[tunnel[1]].name == "start" {
			start = allCaves[tunnel[1]]
		}
	}

	return start
}

func copyMap(src map[string]bool) map[string]bool {
	target := make(map[string]bool)
	for key, value := range src {
		target[key] = value
	}
	return target
}

func NumberOfPaths1(filename string) int {
	start := readData(filename)
	visitedSmallCaves := map[string]bool{"start": true}

	return len(visit1(start, []*cave{}, visitedSmallCaves))
}

func visit1(c *cave, currentPath []*cave, visitedSmallCaves map[string]bool) [][]*cave {
	currentPath = append(currentPath, c)
	if c.name == "end" {
		return [][]*cave{currentPath}
	}

	if c.small() {
		visitedSmallCaves[c.name] = true
	}

	paths := make([][]*cave, 0)
	for _, n := range c.neighbors {
		if !visitedSmallCaves[n.name] {
			newPath := make([]*cave, len(currentPath))
			copy(newPath, currentPath)
			for _, p := range visit1(n, newPath, copyMap(visitedSmallCaves)) {
				if len(p) > 0 {
					paths = append(paths, p)
				}
			}
		}
	}
	return paths
}

func NumberOfPaths2(filename string) int {
	start := readData(filename)
	visitedSmallCaves := map[string]bool{"start": true}

	return len(visit2(start, []*cave{}, visitedSmallCaves, false))
}

func visit2(c *cave, currentPath []*cave, visitedSmallCaves map[string]bool, visitedSmallCaveTwice bool) [][]*cave {
	currentPath = append(currentPath, c)
	if c.name == "end" {
		return [][]*cave{currentPath}
	}

	if c.small() {
		visitedSmallCaves[c.name] = true
	}

	paths := make([][]*cave, 0)
	for _, n := range c.neighbors {
		if n.name != "start" && (!visitedSmallCaves[n.name] || !visitedSmallCaveTwice) {
			newPath := make([]*cave, len(currentPath))
			copy(newPath, currentPath)
			for _, p := range visit2(n, newPath, copyMap(visitedSmallCaves), visitedSmallCaves[n.name] || visitedSmallCaveTwice) {
				if len(p) > 0 {
					paths = append(paths, p)
				}
			}
		}
	}
	return paths
}
