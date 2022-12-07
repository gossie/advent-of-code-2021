package day22

import (
	"testing"
)

func TestNumberOfEnabledCubes(t *testing.T) {
	count := NumberOfEnabledCubes("day22.txt", true)
	if count != 609563 {
		t.Fatalf("count = %v", count)
	}
}

func TestNumberOfEnabledCubesWithoutLimit(t *testing.T) {
	count := NumberOfEnabledCubes("day22.txt", false)
	if count != 2758514936282235 {
		t.Fatalf("count = %v", count)
	}
}
