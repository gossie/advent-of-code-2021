package day19

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type point struct {
	x, y, z int
}

type beacon struct {
	position point
}

type beaconScanner struct {
	beacons  []beacon
	position point
}

func readData(filename string) []beaconScanner {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	scanners := make([]beaconScanner, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if strings.HasPrefix(line, "---") {
				scanners = append(scanners, beaconScanner{beacons: make([]beacon, 0)})
			} else {
				var coordinates []int
				for _, sCoordinate := range strings.Split(line, ",") {
					coordinate, _ := strconv.Atoi(sCoordinate)
					coordinates = append(coordinates, coordinate)
				}
				beacon := beacon{position: point{x: coordinates[0], y: coordinates[1], z: coordinates[2]}}
				scanners[len(scanners)-1].beacons = append(scanners[len(scanners)-1].beacons, beacon)
			}
		}
	}
	return scanners
}

func assumeBeaconEquality(b1 *beacon, b2 *beacon) (int, int, int) {
	return b1.position.x - b2.position.x, b1.position.y - b2.position.y, b1.position.z - b2.position.z
}

func checkScannerEquality(s1 *beaconScanner, s2 *beaconScanner) bool {
	return false
}

func DistinctBeacons(filename string) int {
	scanners := readData(filename)
	numberOfBeacons := 0
	scanners[0].position = point{0, 0, 0}

	// for _, s:= range scanners[1:] {
	// 	for orientationIndex := 0; orientationIndex < 8; orientationIndex++ {
	// 		matches := 0

	// 	}
	// }

	// for orientationIndex := 0; orientationIndex < 8; orientationIndex++ {
	// 	matches := 0
	// 	for srcIndex, src := range scanners {
	// 		for targetIndex, target := range scanners {
	// 			if srcIndex != targetIndex {
	// 				for _, srcBeacon := range src.beacons {
	// 					for _, targetBeacon := range target.beacons {
	// 						offsetX, offsetY, offsetZ := assumeBeaconEquality(&srcBeacon, &targetBeacon)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}

	// 	if matches >= 12 {

	// 		break
	// 	}
	// }
	renderScanners(scanners[0:1])
	return numberOfBeacons
}

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
