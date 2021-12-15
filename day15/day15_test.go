package day15

import "testing"

func TestShortestPath(t *testing.T) {
	totalRisk := MinimalRisk("day15.txt", 1)
	if totalRisk != 595 {
		t.Fatalf("totalRisk = %v", totalRisk)
	}
}

func TestShortestPathLarge(t *testing.T) {
	totalRisk := MinimalRisk("day15.txt", 5)
	if totalRisk != 2914 {
		t.Fatalf("totalRisk = %v", totalRisk)
	}
}
