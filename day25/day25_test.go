package day25

import (
	"testing"
)

func TestPlayTestGame(t *testing.T) {
	state := WhichStep("day25.txt")
	if state != 471 {
		t.Fatalf("state = %v", state)
	}
}
