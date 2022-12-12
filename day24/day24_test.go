package day24

import "testing"

func TestMaxModelNumber(t *testing.T) {
	modelNumber := MaxModelNumber("day24.txt")
	if modelNumber != "99911993949684" {
		t.Fatalf("model number = %v", modelNumber)
	}
}

func TestMinModelNumber(t *testing.T) {
	modelNumber := MinModelNumber("day24.txt")
	if modelNumber != "62911941716111" {
		t.Fatalf("model number = %v and not %v", modelNumber)
	}
}
