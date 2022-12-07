package day19

import "testing"

func TestDistinctBeacons(t *testing.T) {
	count := DistinctBeacons("day19_test.txt")
	if count != 79 {
		t.Fatalf("count = %v", count)
	}
}

func TestManhattenDistance(t *testing.T) {
	distance := ManhattenDistance("day19_test.txt")
	if distance != 3621 {
		t.Fatalf("distance = %v", distance)
	}
}
