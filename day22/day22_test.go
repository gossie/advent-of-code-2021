package day22

import (
	"testing"
)

func TestNumberOfEnabledCubes(t *testing.T) {
	count := NumberOfEnabledCubes("day22.txt")
	if count != 609563 {
		t.Fatalf("count = %v", count)
	}
}
