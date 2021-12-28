package day19

import "testing"

func TestDistinctBeacons(t *testing.T) {
	count := DistinctBeacons("day19.txt")
	if count != 1014 {
		t.Fatalf("count = %v", count)
	}
}
