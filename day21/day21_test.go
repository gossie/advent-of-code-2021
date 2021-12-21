package day21

import (
	"testing"
)

func TestPlayTestGame(t *testing.T) {
	count := PlayTestGame()
	if count != 925605 {
		t.Fatalf("count = %v", count)
	}
}
