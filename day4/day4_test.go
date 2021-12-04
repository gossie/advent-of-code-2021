package day4

import (
	"testing"
)

func TestBingoFirstWin(t *testing.T) {
	points := BingoFirstWin("day4.txt")
	if points != 6592 {
		t.Fatalf("points = %v", points)
	}
}

func TestBingoLastWin(t *testing.T) {
	points := BingoLastWin("day4.txt")
	if points != 31755 {
		t.Fatalf("lifeSupport = %v", points)
	}
}
