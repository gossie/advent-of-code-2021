package day10

import "testing"

func TestScoreForBrokenLines(t *testing.T) {
	score := ScoreForBrokenLines("day10.txt")
	if score != 265527 {
		t.Fatalf("score = %v", score)
	}
}

func TestScoreForIncompleteLines(t *testing.T) {
	score := ScoreForIncompleteLines("day10.txt")
	if score != 3969823589 {
		t.Fatalf("score = %v", score)
	}
}
