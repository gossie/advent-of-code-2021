package day19

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gossie/adventofcode2021/day19/beacons"
)

func readData(filename string) []beacons.BeaconScanner {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	scanners := make([]beacons.BeaconScanner, 0)
	currentBeacons := make([]beacons.Beacon, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if strings.HasPrefix(line, "---") {
				scanners = append(scanners, beacons.CreateBeaconScanner(currentBeacons))
				currentBeacons = make([]beacons.Beacon, 0)
			} else {
				var coordinates []int
				for _, sCoordinate := range strings.Split(line, ",") {
					coordinate, _ := strconv.Atoi(sCoordinate)
					coordinates = append(coordinates, coordinate)
				}

				currentBeacons = append(currentBeacons, beacons.CreateBeacon(beacons.CreatePoint(coordinates[0], coordinates[1], coordinates[2])))
			}
		}
	}
	scanners = append(scanners, beacons.CreateBeaconScanner(currentBeacons))
	return scanners[1:]
}

// func assumeBeaconEquality(b1 *Beacon, b2 *Beacon) (int, int, int) {
// 	return b1.position.x - b2.position.x, b1.position.y - b2.position.y, b1.position.z - b2.position.z
// }

// func checkScannerEquality(s1 *beaconScanner, s2 *beaconScanner) bool {
// 	return false
// }

func overlapping(angles []int, first *beacons.BeaconScanner, other beacons.BeaconScanner) (bool, *beacons.Point) {
	for _, angleX := range angles {
		for _, angleY := range angles {
			for _, angleZ := range angles {
				//println("rotation x: ", angleX, " y: ", angleY, " z: ", angleZ)
				rotated := other.Rotate(angleX, angleY, angleZ)
				for _, b1 := range first.Beacons {
					for _, b2 := range rotated.Beacons {
						delta := b2.AssumeToBe(b1)
						transposed := rotated.Transpose(delta)
						if first.Overlapping(transposed) {
							println("overlap: adding missing beacons to first")
							first.AddMissingBeacons(transposed.Beacons)
							println("added missing beacons to first")
							return true, &delta
						}
					}
				}
			}
		}
	}
	return false, nil
}

func contains(s []int, value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func DistinctBeacons(filename string) int {
	scanners := readData(filename)
	angles := []int{0, 90, 180, 270}

	merged := make([]int, 0)
	first := scanners[0]
	changed := true

	for changed {
		changed = false
		for index := 1; index < len(scanners); index++ {
			if !contains(merged, index) {
				println("check index ", index)
				other := scanners[index]
				overlapped, scannerPosition := overlapping(angles, &first, other)
				if overlapped {
					other.Position = *scannerPosition
					merged = append(merged, index)
					changed = true
				}
			}
		}
	}

	return len(first.Beacons)
}

func ManhattenDistance(filename string) int {
	scanners := readData(filename)
	angles := []int{0, 90, 180, 270}

	merged := make([]int, 0)
	first := scanners[0]
	changed := true

	first.Position = beacons.CreatePoint(0, 0, 0)

	for changed {
		changed = false
		for index := 1; index < len(scanners); index++ {
			if !contains(merged, index) {
				println("check index ", index)
				overlapped, scannerPosition := overlapping(angles, &first, scanners[index])
				if overlapped {
					scanners[index].Position = *scannerPosition
					merged = append(merged, index)
					changed = true
				}
			}
		}
	}

	maxDistance := 0
	for i := 0; i < len(scanners); i++ {
		for j := i + 1; j < len(scanners); j++ {
			distance := calcManhattenDistance(&scanners[i], &scanners[j])
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}

func calcManhattenDistance(s1, s2 *beacons.BeaconScanner) int {
	return int(math.Abs(float64(s1.Position.X-s2.Position.X)) + math.Abs(float64(s1.Position.Y-s2.Position.Y)) + math.Abs(float64(s1.Position.Z-s2.Position.Z)))
}

/*
var scatter3DColor = []string{"#313695"}

func genScatter3dData(scanners []beaconScanner) []opts.Chart3DData {
	data := make([]opts.Chart3DData, 0)

	for _, s := range scanners {
		data = append(data, opts.Chart3DData{Value: []interface{}{s.position.x, s.position.y, s.position.z}, ItemStyle: &opts.ItemStyle{Opacity: 1.0}})
		for _, b := range s.beacons {
			data = append(data, opts.Chart3DData{Value: []interface{}{b.position.x, b.position.y, b.position.z}, ItemStyle: &opts.ItemStyle{Opacity: 0.5}})
		}
	}
	return data
}

func renderScanners(scanners []beaconScanner) {
	page := components.NewPage()
	page.AddCharts(scatter3DBase(scanners))

	f, err := os.Create("scanner.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

func scatter3DBase(scanners []beaconScanner) *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic Scatter3D example"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        100,
			InRange:    &opts.VisualMapInRange{Color: scatter3DColor},
		}),
	)

	scatter3d.AddSeries("scatter3d", genScatter3dData(scanners))
	return scatter3d
}
*/
