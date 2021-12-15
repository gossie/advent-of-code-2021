package day15

import "testing"

func TestShortestPath(t *testing.T) {
	totalRisk := MinimalRisk("day15.txt", 1)
	if totalRisk != 595 {
		t.Fatalf("totalRisk = %v", totalRisk)
	}
}
