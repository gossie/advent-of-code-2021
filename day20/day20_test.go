package day20

import (
	"testing"
)

func TestNumberOfLitPixels(t *testing.T) {
	count := NumberOfLitPixels("day20.txt", 2)
	if count != 5203 {
		t.Fatalf("count = %v", count)
	}
}

func TestNumberOfLitPixelsLarge(t *testing.T) {
	count := NumberOfLitPixels("day20.txt", 50)
	if count != 18806 {
		t.Fatalf("count = %v", count)
	}
}
