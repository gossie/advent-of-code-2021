package day18

import (
	"testing"
)

func TestMagnitude(t *testing.T) {
	magnitude := Magnitude("day18.txt")
	if magnitude != 4202 {
		t.Fatalf("magnitude = %v", magnitude)
	}
}

func TestLargestMagnitude(t *testing.T) {
	magnitude := LargestMagnitude("day18.txt")
	if magnitude != 4779 {
		t.Fatalf("magnitude = %v", magnitude)
	}
}
