package day5

import (
	"testing"
)

func TestSaveAreaExcludingDiagonal(t *testing.T) {
	result := AvoidDangerousArea("day5.txt", false)
	if result != 6548 {
		t.Fatalf("result = %v", result)
	}
}

func TestSaveAreaIncludingDiagonal(t *testing.T) {
	result := AvoidDangerousArea("day5.txt", true)
	if result != 19663 {
		t.Fatalf("result = %v", result)
	}
}
