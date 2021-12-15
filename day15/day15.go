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

type PriorityQueue []*path

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*path)
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on expiration number as the priority
	// The lower the expiry, the higher the priority
	return pq[i].totalRisk < pq[j].totalRisk
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Len() int { return len(*pq) }

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
		for fx := 1; fx < factor; fx++ {
			for x := 0; x < 100; x++ {
				risk := row[x].risk + fx
				if risk > 9 {
					risk -= 9
				}
				row = append(row, point{x: 100 + fx*x, y: len(field), risk: risk})
			}
		}
		field = append(field, row)
	}

	for fy := 1; fy < factor; fy++ {
		for y := 0; y < 100; y++ {
			row := make([]point, 0, 100*factor)
			for x := 0; x < 100; x++ {
				risk := field[y][x].risk + fy
				if risk > 9 {
					risk -= 9
				}
				row = append(row, point{x: x, y: len(field), risk: risk})
			}

			for fx := 1; fx < factor; fx++ {
				for x := 0; x < 100; x++ {
					risk := row[x].risk + fx
					if risk > 9 {
						risk -= 9
					}
					row = append(row, point{x: 100 + fx*x, y: len(field), risk: risk})
				}
			}

			field = append(field, row)
		}
	}

	// for _, row := range field {
	// 	for _, p := range row {
	// 		fmt.Print(p.risk)
	// 	}
	// 	fmt.Println()
	// }

	return field
}

func paths(field [][]point, currentPath *path, visited map[point]int) []*path {
	x, y := currentPath.currentPosition.x, currentPath.currentPosition.y
	if x > 0 && x < len(field[0])-1 && y > 0 && y < len(field)-1 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y-1][x]]; !present || value > currentPath.totalRisk+field[y-1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y-1][x], totalRisk: currentPath.totalRisk + field[y-1][x].risk})
		}
		if value, present := visited[field[y][x+1]]; !present || value > currentPath.totalRisk+field[y][x+1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x+1], totalRisk: currentPath.totalRisk + field[y][x+1].risk})
		}
		if value, present := visited[field[y+1][x]]; !present || value > currentPath.totalRisk+field[y+1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y+1][x], totalRisk: currentPath.totalRisk + field[y+1][x].risk})
		}
		if value, present := visited[field[y][x-1]]; !present || value > currentPath.totalRisk+field[y][x-1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x-1], totalRisk: currentPath.totalRisk + field[y][x-1].risk})
		}
		return newPaths
	} else if x == 0 && y > 0 && y < len(field)-1 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y-1][x]]; !present || value > currentPath.totalRisk+field[y-1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y-1][x], totalRisk: currentPath.totalRisk + field[y-1][x].risk})
		}
		if value, present := visited[field[y][x+1]]; !present || value > currentPath.totalRisk+field[y][x+1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x+1], totalRisk: currentPath.totalRisk + field[y][x+1].risk})
		}
		if value, present := visited[field[y+1][x]]; !present || value > currentPath.totalRisk+field[y+1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y+1][x], totalRisk: currentPath.totalRisk + field[y+1][x].risk})
		}
		return newPaths
	} else if x == len(field[0])-1 && y > 0 && y < len(field)-1 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y-1][x]]; !present || value > currentPath.totalRisk+field[y-1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y-1][x], totalRisk: currentPath.totalRisk + field[y-1][x].risk})
		}
		if value, present := visited[field[y+1][x]]; !present || value > currentPath.totalRisk+field[y+1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y+1][x], totalRisk: currentPath.totalRisk + field[y+1][x].risk})
		}
		if value, present := visited[field[y][x-1]]; !present || value > currentPath.totalRisk+field[y][x-1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x-1], totalRisk: currentPath.totalRisk + field[y][x-1].risk})
		}
		return newPaths
	} else if x > 0 && x < len(field[0])-1 && y == 0 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y][x+1]]; !present || value > currentPath.totalRisk+field[y][x+1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x+1], totalRisk: currentPath.totalRisk + field[y][x+1].risk})
		}
		if value, present := visited[field[y+1][x]]; !present || value > currentPath.totalRisk+field[y+1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y+1][x], totalRisk: currentPath.totalRisk + field[y+1][x].risk})
		}
		if value, present := visited[field[y][x-1]]; !present || value > currentPath.totalRisk+field[y][x-1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x-1], totalRisk: currentPath.totalRisk + field[y][x-1].risk})
		}
		return newPaths
	} else if x > 0 && x < len(field[0])-1 && y == len(field)-1 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y-1][x]]; !present || value > currentPath.totalRisk+field[y-1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y-1][x], totalRisk: currentPath.totalRisk + field[y-1][x].risk})
		}
		if value, present := visited[field[y][x+1]]; !present || value > currentPath.totalRisk+field[y][x+1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x+1], totalRisk: currentPath.totalRisk + field[y][x+1].risk})
		}
		if value, present := visited[field[y][x-1]]; !present || value > currentPath.totalRisk+field[y][x-1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x-1], totalRisk: currentPath.totalRisk + field[y][x-1].risk})
		}
		return newPaths
	} else if x == 0 && y == 0 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y][x+1]]; !present || value > currentPath.totalRisk+field[y][x+1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x+1], totalRisk: currentPath.totalRisk + field[y][x+1].risk})
		}
		if value, present := visited[field[y+1][x]]; !present || value > currentPath.totalRisk+field[y+1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y+1][x], totalRisk: currentPath.totalRisk + field[y+1][x].risk})
		}
		return newPaths
	} else if x == len(field[0])-1 && y == len(field)-1 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y-1][x]]; !present || value > currentPath.totalRisk+field[y-1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y-1][x], totalRisk: currentPath.totalRisk + field[y-1][x].risk})
		}
		if value, present := visited[field[y][x-1]]; !present || value > currentPath.totalRisk+field[y][x-1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x-1], totalRisk: currentPath.totalRisk + field[y][x-1].risk})
		}
		return newPaths
	} else if x == 0 && y == len(field)-1 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y-1][x]]; !present || value > currentPath.totalRisk+field[y-1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y-1][x], totalRisk: currentPath.totalRisk + field[y-1][x].risk})
		}
		if value, present := visited[field[y][x+1]]; !present || value > currentPath.totalRisk+field[y][x+1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x+1], totalRisk: currentPath.totalRisk + field[y][x+1].risk})
		}
		return newPaths
	} else if x == len(field[0])-1 && y == 0 {
		newPaths := make([]*path, 0)
		if value, present := visited[field[y+1][x]]; !present || value > currentPath.totalRisk+field[y+1][x].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y+1][x], totalRisk: currentPath.totalRisk + field[y+1][x].risk})
		}
		if value, present := visited[field[y][x-1]]; !present || value > currentPath.totalRisk+field[y][x-1].risk {
			newPaths = append(newPaths, &path{currentPosition: field[y][x-1], totalRisk: currentPath.totalRisk + field[y][x-1].risk})
		}
		return newPaths
	}
	panic("what?")
}

func MinimalRisk(filename string, factor int) int {
	field := readData(filename, factor)

	priorityQueue := make(PriorityQueue, 0)

	heap.Init(&priorityQueue)

	start := field[0][0]
	visited := map[point]int{start: start.risk}

	currentPath := path{currentPosition: start, totalRisk: start.risk}
	heap.Push(&priorityQueue, &currentPath)
	var bestPath path

	for len(priorityQueue) > 0 {
		for _, p := range paths(field, &currentPath, visited) {
			heap.Push(&priorityQueue, p)
		}
		currentPath = *heap.Pop(&priorityQueue).(*path)
		if currentPath.currentPosition.x == 99 && currentPath.currentPosition.y == 99 {
			if bestPath.totalRisk == 0 || bestPath.totalRisk > currentPath.totalRisk {
				bestPath = currentPath
			}
		}

		visited[currentPath.currentPosition] = currentPath.totalRisk
	}
	return bestPath.totalRisk - field[0][0].risk
}
