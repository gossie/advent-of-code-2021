package day15

import (
	"bufio"
	"container/heap"
	"os"
)

type point struct {
	x    int
	y    int
	risk int
}

type path struct {
	currentPosition point
	totalRisk       int
}

type priorityQueue []*path

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*path)
	*pq = append(*pq, item)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].totalRisk < pq[j].totalRisk
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Len() int { return len(*pq) }

func (pq *priorityQueue) contains(p *path) bool {
	for i := range *pq {
		if (*pq)[i].currentPosition == p.currentPosition && (*pq)[i].totalRisk == p.totalRisk {
			return true
		}
	}
	return false
}

func readData(filename string, factor int) [][]point {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	field := make([][]point, 0, 100*factor)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := make([]point, 0, 100*factor)
		line := scanner.Text()
		for x, r := range line {
			row = append(row, point{x: x, y: len(field), risk: int(r - '0')})
		}
		row = extendRowByFactor(row, factor, field)
		field = append(field, row)
	}

	for fy := 1; fy < factor; fy++ {
		for y := 0; y < 100; y++ {
			row := createAdditionalRow(field, 100*factor, y, fy)
			row = extendRowByFactor(row, factor, field)
			field = append(field, row)
		}
	}

	return field
}

func createAdditionalRow(field [][]point, rowCapacity int, baseY, yFactor int) []point {
	row := make([]point, 0, rowCapacity)
	for x := 0; x < 100; x++ {
		risk := ((field[baseY][x].risk + yFactor - 1) % 9) + 1
		row = append(row, point{x: x, y: len(field), risk: risk})
	}
	return row
}

func extendRowByFactor(row []point, factor int, field [][]point) []point {
	for fx := 1; fx < factor; fx++ {
		for x := 0; x < 100; x++ {
			risk := ((row[x].risk + fx - 1) % 9) + 1
			row = append(row, point{x: fx*100 + x, y: len(field), risk: risk})
		}
	}
	return row
}

func appendPathIfNecessary(newPaths []*path, currentPath *path, p point, visitedRiskMapping map[point]int) []*path {
	if value, present := visitedRiskMapping[p]; !present || value > currentPath.totalRisk+p.risk {
		newPaths = append(newPaths, &path{currentPosition: p, totalRisk: currentPath.totalRisk + p.risk})
	}
	return newPaths
}

func pathsToNeighbors(field [][]point, currentPath *path, visitedRiskMapping map[point]int) []*path {
	x, y := currentPath.currentPosition.x, currentPath.currentPosition.y
	newPaths := make([]*path, 0)
	if y-1 >= 0 {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y-1][x], visitedRiskMapping)
	}
	if x+1 <= len(field[0])-1 {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y][x+1], visitedRiskMapping)
	}
	if y+1 <= len(field)-1 {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y+1][x], visitedRiskMapping)
	}
	if x-1 >= 0 {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y][x-1], visitedRiskMapping)
	}
	return newPaths
}

func MinimalRisk(filename string, factor int) int {
	field := readData(filename, factor)

	priorityQueue := make(priorityQueue, 0)
	heap.Init(&priorityQueue)

	start := field[0][0]
	visitedRiskMapping := map[point]int{start: start.risk}
	bestPath := path{currentPosition: start, totalRisk: start.risk}
	for bestPath.currentPosition.x != 100*factor-1 || bestPath.currentPosition.y != 100*factor-1 {
		for _, p := range pathsToNeighbors(field, &bestPath, visitedRiskMapping) {
			if !priorityQueue.contains(p) {
				heap.Push(&priorityQueue, p)
			}
		}

		bestPath = *heap.Pop(&priorityQueue).(*path)
		visitedRiskMapping[bestPath.currentPosition] = bestPath.totalRisk
	}
	return bestPath.totalRisk - field[0][0].risk
}
