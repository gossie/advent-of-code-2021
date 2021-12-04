package day2

import (
	"testing"
)

func TestPosition1(t *testing.T) {
	position := SimplePosition("day2.txt")
	if position != 1654760 {
		t.Fatalf("position = %v", position)
	}
}

func TestPosition2(t *testing.T) {
	position := AimedPosition("day2.txt")
	if position != 1956047400 {
		t.Fatalf("position = %v", position)
	}
}
