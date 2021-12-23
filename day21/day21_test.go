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

func TestMultipleUniverses(t *testing.T) {
	count := MultipleUniverses()
	if count != 486638407378784 {
		t.Fatalf("count = %v", count)
	}
}
